package internal

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// GetDifference returns the diff for two projects.
func GetDifference(a, b interface{}, raw ...bool) string {
	dmp := diffmatchpatch.New()

	var diffs []diffmatchpatch.Diff
	if len(raw) == 0 || !raw[0] {
		aString, aOk := a.(string)
		bString, bOk := b.(string)
		if aOk && bOk {
			diffs = dmp.DiffMain(aString, bString, false)
		} else {
			diffs = dmp.DiffMain(strings.TrimSpace(spew.Sdump(a)), strings.TrimSpace(spew.Sdump(b)), false)
		}
	} else {
		diffs = dmp.DiffMain(fmt.Sprint(a), fmt.Sprint(b), false)
	}

	diffs = dmp.DiffCleanupEfficiency(diffs)
	diffs = dmp.DiffCleanupSemanticLossless(diffs)

	d := &diffPrinter{
		ExpectedI:       1,
		ActualI:         1,
		UnchangedPrefix: "#",
		ExpectedPrefix:  pterm.FgRed.Sprint("-"),
		ActualPrefix:    pterm.FgGreen.Sprint("+"),
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

	ExpectedBuffer strings.Builder
	ActualBuffer   strings.Builder

	ExpectedGroupBuffer strings.Builder
	ActualGroupBuffer   strings.Builder

	ExpectedFlushable bool
	ActualFlushable   bool

	UnchangedPrefix string
	ExpectedPrefix  string
	ActualPrefix    string
}

func (d *diffPrinter) ProcessDiffs(diffs []diffmatchpatch.Diff) {
	for _, diff := range diffs {
		for _, char := range diff.Text {
			if char == '\n' {
				d.FlushDiff(diff.Type)
				d.FlushLine(diff.Type)
			} else {
				if diff.Type == diffmatchpatch.DiffEqual {
					d.ExpectedBuffer.WriteRune(char)
					d.ActualBuffer.WriteRune(char)
				} else if diff.Type == diffmatchpatch.DiffInsert {
					d.ActualBuffer.WriteRune(char)
				} else {
					d.ExpectedBuffer.WriteRune(char)
				}
			}
		}
		d.FlushDiff(diff.Type)
	}

	lastOp := diffmatchpatch.DiffEqual
	if len(diffs) > 0 {
		lastOp = diffs[len(diffs)-1].Type
	}

	d.FlushLine(lastOp)
	d.Finalize()
}

func (d *diffPrinter) FlushDiff(operation diffmatchpatch.Operation) {
	if d.ExpectedBuffer.Len() > 0 || d.ActualBuffer.Len() > 0 {
		if operation == diffmatchpatch.DiffDelete {
			d.ExpectedLine.WriteString(pterm.BgDarkGray.Sprint(pterm.Bold.Sprint(d.ExpectedBuffer.String())))
			d.ExpectedFlushable = true
		} else if operation == diffmatchpatch.DiffInsert {
			d.ActualLine.WriteString(pterm.BgDarkGray.Sprint(pterm.Bold.Sprint(d.ActualBuffer.String())))
			d.ActualFlushable = true
		} else {
			d.ExpectedLine.WriteString(d.ExpectedBuffer.String())
			d.ActualLine.WriteString(d.ActualBuffer.String())
			d.ExpectedFlushable = true
			d.ActualFlushable = true
		}
	}

	d.ExpectedBuffer.Reset()
	d.ActualBuffer.Reset()
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
		} else if operation == diffmatchpatch.DiffDelete {
			d.ExpectedGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", d.ExpectedI, d.ExpectedPrefix, pterm.FgRed.Sprint(d.ExpectedLine.String())))
			d.ExpectedI++
		} else {
			d.ActualGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", d.ActualI, d.ActualPrefix, pterm.FgGreen.Sprint(d.ActualLine.String())))
			d.ActualI++
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
		d.ExpectedBuffer.Reset()
	}

	if d.ActualFlushable {
		d.ActualFlushable = false
		d.ActualLine.Reset()
		d.ActualBuffer.Reset()
	}
}

func (d *diffPrinter) Finalize() {
	d.Text += d.ExpectedGroupBuffer.String()
	d.Text += d.ActualGroupBuffer.String()
}
