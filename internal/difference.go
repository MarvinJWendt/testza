package internal

import (
	"fmt"
	"math"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var DiffContextLines = 2

func NewDiffObject(expected, actual any, raw ...bool) Object {
	return Object{
		Name:      "Difference",
		NameStyle: pterm.NewStyle(pterm.FgYellow),
		Data:      GetDifference(expected, actual, raw...),
		Raw:       true,
	}
}

// GetDifference returns the diff for two projects.
func GetDifference(a, b any, raw ...bool) string {
	dmp := diffmatchpatch.New()

	var aString, bString string
	var diffs []diffmatchpatch.Diff
	if len(raw) == 0 || !raw[0] {
		var aOk, bOk bool
		aString, aOk = a.(string)
		bString, bOk = b.(string)
		if aOk && bOk {
			diffs = dmp.DiffMain(aString, bString, false)
		} else {
			aString = strings.TrimSpace(spew.Sdump(a))
			bString = strings.TrimSpace(spew.Sdump(b))
			diffs = dmp.DiffMain(aString, bString, false)
		}
	} else {
		aString = fmt.Sprint(a)
		bString = fmt.Sprint(b)
		diffs = dmp.DiffMain(aString, bString, false)
	}

	diffs = dmp.DiffCleanupEfficiency(diffs)
	diffs = dmp.DiffCleanupSemanticLossless(diffs)

	maxNewlines := math.Max(float64(strings.Count(aString, "\n")), float64(strings.Count(bString, "\n"))) + 1

	d := &diffPrinter{
		ExpectedI:       1,
		ActualI:         1,
		UnchangedPrefix: "#",
		ExpectedPrefix:  pterm.FgRed.Sprint("-"),
		ActualPrefix:    pterm.FgGreen.Sprint("+"),
		CounterWidth:    int(math.Log10(maxNewlines)) + 1,
	}

	return d.ProcessDiffs(diffs)
}

type textLine struct {
	Text      string
	Operation diffmatchpatch.Operation
}

type diffPrinter struct {
	Text []textLine

	ExpectedI int
	ActualI   int

	ExpectedLine strings.Builder
	ActualLine   strings.Builder

	DiffBuffer strings.Builder

	ExpectedGroupBuffer []textLine
	ActualGroupBuffer   []textLine

	ExpectedFlushable bool
	ActualFlushable   bool

	UnchangedPrefix string
	ExpectedPrefix  string
	ActualPrefix    string

	CounterWidth int
}

func (d *diffPrinter) ProcessDiffs(diffs []diffmatchpatch.Diff) string {
	for _, diff := range diffs {
		for _, char := range diff.Text {
			if char == '\n' {
				d.FlushDiff(diff.Type, true)
				d.FlushLine(diff.Type)
			} else {
				d.DiffBuffer.WriteRune(char)
			}
		}
		d.FlushDiff(diff.Type, false)
	}

	lastOp := diffmatchpatch.DiffEqual
	if len(diffs) > 0 {
		lastOp = diffs[len(diffs)-1].Type
	}

	d.FlushLine(lastOp)
	return d.Finalize(lastOp)
}

func (d *diffPrinter) FlushDiff(operation diffmatchpatch.Operation, newLine bool) {
	if d.DiffBuffer.Len() > 0 {
		if operation == diffmatchpatch.DiffDelete {
			d.ExpectedLine.WriteString(pterm.BgDarkGray.Sprint(pterm.Bold.Sprint(d.DiffBuffer.String())))
		} else if operation == diffmatchpatch.DiffInsert {
			d.ActualLine.WriteString(pterm.BgDarkGray.Sprint(pterm.Bold.Sprint(d.DiffBuffer.String())))
		} else {
			d.ExpectedLine.WriteString(d.DiffBuffer.String())
			d.ActualLine.WriteString(d.DiffBuffer.String())
		}
	}

	if operation == diffmatchpatch.DiffDelete {
		d.ExpectedFlushable = d.ExpectedFlushable || newLine
	} else if operation == diffmatchpatch.DiffInsert {
		d.ActualFlushable = d.ActualFlushable || newLine
	} else {
		d.ExpectedFlushable = d.ExpectedFlushable || newLine
		d.ActualFlushable = d.ActualFlushable || newLine
	}

	d.DiffBuffer.Reset()
}

func (d *diffPrinter) FlushLine(operation diffmatchpatch.Operation) {
	if d.ExpectedLine.String() == d.ActualLine.String() {
		if operation == diffmatchpatch.DiffEqual {
			d.Text = append(d.Text, d.ExpectedGroupBuffer...)
			d.Text = append(d.Text, d.ActualGroupBuffer...)

			d.ExpectedGroupBuffer = make([]textLine, 0)
			d.ActualGroupBuffer = make([]textLine, 0)

			d.Text = append(d.Text, textLine{
				Text:      pterm.FgGray.Sprintf("(%*d. %s) ", d.CounterWidth, d.ActualI, d.UnchangedPrefix) + d.ExpectedLine.String() + "\n",
				Operation: operation,
			})

			d.ActualI++
			d.ExpectedI = d.ActualI

			d.ExpectedFlushable = true
			d.ActualFlushable = true
		} else if operation == diffmatchpatch.DiffDelete {
			d.ExpectedGroupBuffer = append(d.ExpectedGroupBuffer, textLine{
				Text:      pterm.FgGray.Sprintfln("(%*d. %s) %s", d.CounterWidth, d.ExpectedI, d.ExpectedPrefix, pterm.FgRed.Sprint(d.ExpectedLine.String())),
				Operation: diffmatchpatch.DiffDelete,
			})
			d.ExpectedI++
			d.ExpectedFlushable = true
		} else {
			d.ActualGroupBuffer = append(d.ActualGroupBuffer, textLine{
				Text:      pterm.FgGray.Sprintfln("(%*d. %s) %s", d.CounterWidth, d.ActualI, d.ActualPrefix, pterm.FgGreen.Sprint(d.ActualLine.String())),
				Operation: diffmatchpatch.DiffInsert,
			})
			d.ActualI++
			d.ActualFlushable = true
		}
	} else {
		if d.ExpectedFlushable {
			d.ExpectedGroupBuffer = append(d.ExpectedGroupBuffer, textLine{
				Text:      pterm.FgGray.Sprintfln("(%*d. %s) %s", d.CounterWidth, d.ExpectedI, d.ExpectedPrefix, pterm.FgRed.Sprint(d.ExpectedLine.String())),
				Operation: diffmatchpatch.DiffDelete,
			})
			d.ExpectedI++
		}

		if d.ActualFlushable {
			d.ActualGroupBuffer = append(d.ActualGroupBuffer, textLine{
				Text:      pterm.FgGray.Sprintfln("(%*d. %s) %s", d.CounterWidth, d.ActualI, d.ActualPrefix, pterm.FgGreen.Sprint(d.ActualLine.String())),
				Operation: diffmatchpatch.DiffInsert,
			})
			d.ActualI++
		}
	}

	if d.ExpectedFlushable {
		d.ExpectedFlushable = false
		d.ExpectedLine.Reset()
	}

	if d.ActualFlushable {
		d.ActualFlushable = false
		d.ActualLine.Reset()
	}

	d.DiffBuffer.Reset()
}

func (d *diffPrinter) Finalize(operation diffmatchpatch.Operation) string {
	d.ExpectedFlushable = true
	d.ActualFlushable = true

	if d.ExpectedLine.Len() > 0 || d.ActualLine.Len() > 0 {
		d.FlushLine(operation)
	}

	d.Text = append(d.Text, d.ExpectedGroupBuffer...)
	d.Text = append(d.Text, d.ActualGroupBuffer...)

	var resultBuffer strings.Builder
	if DiffContextLines >= 0 {
		requiredLines := make(map[int]bool)

		for i, line := range d.Text {
			if line.Operation != diffmatchpatch.DiffEqual {
				for j := int(math.Max(0, float64(i-DiffContextLines))); j < int(math.Min(float64(len(d.Text)), float64(i+DiffContextLines+1))); j++ {
					requiredLines[j] = true
				}
			}
		}

		hasSnip := false
		for i, line := range d.Text {
			if _, ok := requiredLines[i]; ok {
				resultBuffer.WriteString(line.Text)
				hasSnip = false
			} else if !hasSnip {
				resultBuffer.WriteString(pterm.FgMagenta.Sprint("[...]\n"))
				hasSnip = true
			}
		}
	} else {
		for _, line := range d.Text {
			resultBuffer.WriteString(line.Text)
		}
	}

	return resultBuffer.String()
}
