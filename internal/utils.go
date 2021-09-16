package internal

import (
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"testing"

	"github.com/pterm/pterm"
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

func AssertRegexpHelper(t testRunner, regex interface{}, txt interface{}, shouldMatch bool, msg ...interface{}) {
	regexString := fmt.Sprint(regex)
	txtString := fmt.Sprint(txt)
	match, _ := regexp.MatchString(regexString, txtString)
	if shouldMatch != match {
		failText := "!!does not match!! the string."
		if !shouldMatch {
			failText = "!!does match!! the string !!but should not!!."
		}
		Fail(t, "The regex pattern "+failText, Objects{
			{
				Name:      "Regex Pattern",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      regexString + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
			{
				Name:      "String",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      txtString + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
		}, msg...)
	}
}
