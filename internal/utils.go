package internal

import (
	"github.com/pterm/pterm"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

type testRunner interface {
	Error(args ...interface{})
}

type helper interface {
	Helper()
}

// GetTest converts to *testing.T.
func GetTest(t testRunner) *testing.T {
	if test, ok := t.(*testing.T); ok {
		return test
	}

	return nil
}

// GetCurrentScriptDirectory returns the directory of the current Go file.
func GetCurrentScriptDirectory() string {
	_, scriptPath, _, _ := runtime.Caller(1)
	return filepath.Join(scriptPath, "..")
}

// CompareTwoValuesInASlice is a helper function.
func CompareTwoValuesInASlice(a reflect.Value, compareFunc func(a, b reflect.Value) bool) (ret bool) {
	ret = true
	for i := 1; i < a.Len(); i++ {
		if !compareFunc(a.Index(i-1), a.Index(i)) {
			ret = false
		}
	}

	return
}

func AssertRegexpHelper(t testRunner, txt string, regex string, shouldMatch bool, msg ...interface{}) {
	match, _ := regexp.MatchString(regex, txt)
	if shouldMatch != match {
		failText := "!!does not match!! the string."
		if !shouldMatch {
			failText = "!!does match!! the string !!but shouldn't!!."
		}
		Fail(t, "The regex pattern "+failText, Objects{
			{
				Name:      "Regex Pattern",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      regex + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
			{
				Name:      "String",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      txt + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
		}, msg...)
	}
}
