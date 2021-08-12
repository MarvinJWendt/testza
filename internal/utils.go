package internal

import (
	"path/filepath"
	"runtime"
	"testing"
)

type testingT interface {
	Error(args ...interface{})
}

type helper interface {
	Helper()
}

// GetTest converts to *testing.T.
func GetTest(t testingT) *testing.T {
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
