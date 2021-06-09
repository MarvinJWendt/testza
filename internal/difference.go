package internal

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
)

// GetDifference returns the diff for two projects.
func GetDifference(a, b interface{}) string {
	diff := difflib.UnifiedDiff{
		A: difflib.SplitLines(spew.Sdump(a)),
		B: difflib.SplitLines(spew.Sdump(b)),
	}

	text, _ := difflib.GetUnifiedDiffString(diff)

	return text
}
