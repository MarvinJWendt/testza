package internal

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/pterm/pterm"
)

// GetDifference returns the diff for two projects.
func GetDifference(a, b interface{}) string {
	diff := difflib.UnifiedDiff{
		A: difflib.SplitLines(spew.Sdump(a)),
		B: difflib.SplitLines(spew.Sdump(b)),
	}

	text, _ := difflib.GetUnifiedDiffString(diff)

	var newText string
	for _, v := range strings.Split(text, "\n") {
		if strings.HasPrefix(v, "- ") {
			v = strings.TrimPrefix(v, "- ")
			newText += pterm.FgGreen.Sprint(v) + "\n"
		} else if strings.HasPrefix(v, "+ ") {
			v = strings.TrimPrefix(v, "+ ")
			newText += pterm.FgRed.Sprint(v) + "\n"
		}
	}
	text = newText

	return text
}
