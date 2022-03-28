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
			diffs = dmp.DiffMain(spew.Sdump(a), spew.Sdump(b), false)
		}
	} else {
		diffs = dmp.DiffMain(fmt.Sprint(a), fmt.Sprint(b), false)
	}

	diffs = dmp.DiffCleanupSemanticLossless(diffs)

	var expectedLine strings.Builder
	var actualLine strings.Builder

	text := ""

	var expectedGroupBuffer strings.Builder
	var actualGroupBuffer strings.Builder

	unchangedPrefix := "#"
	expectedPrefix := pterm.FgRed.Sprint("-")
	actualPrefix := pterm.FgGreen.Sprint("+")

	if pterm.RawOutput {
		unchangedPrefix = ""
		expectedPrefix = ""
		actualPrefix = ""
	}

	expectedI := 1
	actualI := 1

	for _, diff := range diffs {
		var expectedBuffer strings.Builder
		var actualBuffer strings.Builder

		for _, char := range diff.Text {
			if char == '\n' {
				if expectedBuffer.Len() > 0 || actualBuffer.Len() > 0 {
					if diff.Type == diffmatchpatch.DiffDelete {
						expectedLine.WriteString(pterm.BgDarkGray.Sprint(pterm.Bold.Sprint(expectedBuffer.String())))
					} else if diff.Type == diffmatchpatch.DiffInsert {
						actualLine.WriteString(pterm.BgDarkGray.Sprint(pterm.Bold.Sprint(actualBuffer.String())))
					} else {
						expectedLine.WriteString(expectedBuffer.String())
						actualLine.WriteString(actualBuffer.String())
					}
				}

				if expectedLine.String() == actualLine.String() {
					if diff.Type == diffmatchpatch.DiffEqual {
						text += expectedGroupBuffer.String()
						text += actualGroupBuffer.String()

						expectedGroupBuffer.Reset()
						actualGroupBuffer.Reset()

						text += pterm.FgGray.Sprintf("(%d. %s) ", actualI, unchangedPrefix) + expectedLine.String() + "\n"
						actualI++
						expectedI = actualI
					} else if diff.Type == diffmatchpatch.DiffDelete {
						expectedGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", expectedI, expectedPrefix, pterm.FgRed.Sprint(expectedLine.String())))
						expectedI++
					} else {
						actualGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", actualI, actualPrefix, pterm.FgGreen.Sprint(actualLine.String())))
						actualI++
					}
				} else {
					expectedGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", expectedI, expectedPrefix, pterm.FgRed.Sprint(expectedLine.String())))
					actualGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", actualI, actualPrefix, pterm.FgGreen.Sprint(actualLine.String())))
					expectedI++
					actualI++
				}

				expectedLine.Reset()
				actualLine.Reset()

				expectedBuffer.Reset()
				actualBuffer.Reset()
			} else {
				if diff.Type == diffmatchpatch.DiffEqual {
					expectedBuffer.WriteRune(char)
					actualBuffer.WriteRune(char)
				} else if diff.Type == diffmatchpatch.DiffInsert {
					actualBuffer.WriteRune(char)
				} else {
					expectedBuffer.WriteRune(char)
				}
			}
		}

		if expectedBuffer.Len() > 0 || actualBuffer.Len() > 0 {
			if diff.Type == diffmatchpatch.DiffDelete {
				expectedLine.WriteString(pterm.BgDarkGray.Sprint(pterm.Bold.Sprint(expectedBuffer.String())))
			} else if diff.Type == diffmatchpatch.DiffInsert {
				actualLine.WriteString(pterm.BgDarkGray.Sprint(pterm.Bold.Sprint(actualBuffer.String())))
			} else {
				expectedLine.WriteString(expectedBuffer.String())
				actualLine.WriteString(actualBuffer.String())
			}
		}
	}

	lastOp := diffmatchpatch.DiffEqual
	if len(diffs) > 0 {
		lastOp = diffs[len(diffs)-1].Type
	}

	if expectedLine.String() == actualLine.String() {
		if lastOp == diffmatchpatch.DiffEqual {
			text += expectedGroupBuffer.String()
			text += actualGroupBuffer.String()

			expectedGroupBuffer.Reset()
			actualGroupBuffer.Reset()

			text += pterm.FgGray.Sprintf("(%d. %s) ", actualI, unchangedPrefix) + expectedLine.String() + "\n"
		} else if lastOp == diffmatchpatch.DiffDelete {
			expectedGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", expectedI, expectedPrefix, pterm.FgRed.Sprint(expectedLine.String())))
		} else {
			actualGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", actualI, actualPrefix, pterm.FgGreen.Sprint(actualLine.String())))
		}
	} else {
		expectedGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", expectedI, expectedPrefix, pterm.FgRed.Sprint(expectedLine.String())))
		actualGroupBuffer.WriteString(pterm.FgGray.Sprintfln("(%d. %s) %s", actualI, actualPrefix, pterm.FgGreen.Sprint(actualLine.String())))
	}

	text += expectedGroupBuffer.String()
	text += actualGroupBuffer.String()

	pterm.Println(FailS(text, Objects{}))

	return text
}
