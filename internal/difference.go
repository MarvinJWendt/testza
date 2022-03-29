package internal

import (
	"fmt"
	"math"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// GetDifference returns the diff for two projects.
func GetDifference(a, b interface{}, raw ...bool) string {
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

	diffLevenshtein := dmp.DiffLevenshtein(diffs)
	maxInput := math.Max(float64(len(aString)), float64(len(bString)))

	fmt.Println("Diff:", diffLevenshtein, "Max:", maxInput, "Diff/Max:", float64(diffLevenshtein)/maxInput)

	diffs = dmp.DiffCleanupEfficiency(diffs)
	diffs = dmp.DiffCleanupSemanticLossless(diffs)

	maxNewlines := math.Max(float64(strings.Count(aString, "\n")), float64(strings.Count(bString, "\n"))) + 1

	d := &diffPrinter{
		ExpectedI:       1,
		ActualI:         1,
		UnchangedPrefix: "#",
		ExpectedPrefix:  pterm.FgRed.Sprint("-"),
		ActualPrefix:    pterm.FgGreen.Sprint("+"),
		LineCount:       int(maxNewlines),
	}

	d.ProcessDiffs(diffs)

	return d.Text
}

type diffPrinter struct {
	Text string

	ExpectedI int
	ActualI   int

	ExpectedLine strings.Builder
	ActualLine   strings.Builder

	DiffBuffer strings.Builder

	ExpectedGroupBuffer strings.Builder
	ActualGroupBuffer   strings.Builder

	ExpectedFlushable bool
	ActualFlushable   bool

	UnchangedPrefix string
	ExpectedPrefix  string
	ActualPrefix    string

	LineCount int
}

func (d *diffPrinter) ProcessDiffs(diffs []diffmatchpatch.Diff) {
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
	d.Finalize(lastOp)
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
			d.Text += d.ExpectedGroupBuffer.String()
			d.Text += d.ActualGroupBuffer.String()

			d.ExpectedGroupBuffer.Reset()
			d.ActualGroupBuffer.Reset()

			d.Text += pterm.FgGray.Sprintf("(%d. %s) ", d.ActualI, d.UnchangedPrefix) + d.ExpectedLine.String() + "\n"
			d.ActualI++
			d.ExpectedI = d.ActualI

			d.ExpectedFlushable = true
			d.ActualFlushable = true
		} else if operation == diffmatchpatch.DiffDelete {
			d.ExpectedGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", d.ExpectedI, d.ExpectedPrefix, pterm.FgRed.Sprint(d.ExpectedLine.String())))
			d.ExpectedI++
			d.ExpectedFlushable = true
		} else {
			d.ActualGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", d.ActualI, d.ActualPrefix, pterm.FgGreen.Sprint(d.ActualLine.String())))
			d.ActualI++
			d.ActualFlushable = true
		}
	} else {
		if d.ExpectedFlushable {
			d.ExpectedGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", d.ExpectedI, d.ExpectedPrefix, pterm.FgRed.Sprint(d.ExpectedLine.String())))
			d.ExpectedI++
		}

		if d.ActualFlushable {
			d.ActualGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", d.ActualI, d.ActualPrefix, pterm.FgGreen.Sprint(d.ActualLine.String())))
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

func (d *diffPrinter) Finalize(operation diffmatchpatch.Operation) {
	d.ExpectedFlushable = true
	d.ActualFlushable = true

	if d.ExpectedLine.Len() > 0 || d.ActualLine.Len() > 0 {
		d.FlushLine(operation)
	}

	d.Text += d.ExpectedGroupBuffer.String()
	d.Text += d.ActualGroupBuffer.String()
}
