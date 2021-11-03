package testza

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/pterm/pterm"
)

type testRunner interface {
	Error(args ...interface{})
}

// TestingPackageWithFailFunctions contains every function that fails a test in testing.T.
type TestingPackageWithFailFunctions interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type helper interface {
	Helper()
}

var green = pterm.NewStyle(pterm.Bold, pterm.FgLightGreen).Sprint
var red = pterm.NewStyle(pterm.Bold, pterm.FgLightRed).Sprint
var highlight = red

func generateMsg(msg []interface{}, addon ...interface{}) (out string) {
	for _, s := range addon {
		out += fmt.Sprint(s)
	}
	for _, s := range msg {
		out += fmt.Sprint(s)
	}

	return
}

func getCurrentScriptDirectory() string {
	_, scriptPath, _, _ := runtime.Caller(2)
	return filepath.Join(scriptPath, "..")
}
