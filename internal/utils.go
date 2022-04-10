package internal

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

type testRunner interface {
	Error(args ...any)
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
