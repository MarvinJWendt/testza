package testza

import (
	"github.com/MarvinJWendt/testza/internal"
	"github.com/pterm/pterm"
)

// SetColorsEnabled controls if testza should print colored output.
// You should use this in the init() method of the package, which contains your tests.
//
// Example:
//  init() {
//    testza.SetColorsEnabled(false) // Disable colored output
//    testza.SetColorsEnabled(true)  // Enable colored output
//  }
func SetColorsEnabled(enabled bool) {
	if enabled {
		pterm.EnableColor()
	} else {
		pterm.DisableColor()
	}
}

// SetLineNumbersEnabled controls if line numbers should be printed in failing tests.
// You should use this in the init() method of the package, which contains your tests.
//
// Example:
//  init() {
//    testza.SetLineNumbersEnabled(false) // Disable line numbers
//    testza.SetLineNumbersEnabled(true)  // Enable line numbers
//  }
func SetLineNumbersEnabled(enabled bool) {
	internal.LineNumbersEnabled = enabled
}
