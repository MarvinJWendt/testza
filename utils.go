package testza

import (
	"fmt"

	"github.com/pterm/pterm"
)

type testRunner interface {
	Error(args ...interface{})
}

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
