package testutil

import (
	"fmt"

	"github.com/pterm/pterm"
)

type TestingT interface {
	Error(args ...interface{})
}

type helper interface {
	Helper()
}

type inputValidator interface {
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
