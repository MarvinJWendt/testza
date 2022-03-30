package testza

import (
	"math/rand"

	"github.com/MarvinJWendt/testza/internal"
	"github.com/pterm/pterm"
)

var showStartupMessage = true

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

// SetRandomSeed sets the seed for the random generator used in testza.
// Using the same seed will result in the same random sequences each time and guarantee a reproducible test run.
// Use this setting, if you want a 100% deterministic test.
// You should use this in the init() method of the package, which contains your tests.
//
// Example:
//  init() {
//    testza.SetRandomSeed(1337) // Set the seed to 1337
//    testza.SetRandomSeed(time.Now().UnixNano()) // Set the seed back to the current time (default | non-deterministic)
//  }
func SetRandomSeed(seed int64) {
	randomSeed = seed
	rand.Seed(seed)
}

// SetShowStartupMessage controls if the startup message should be printed.
// You should use this in the init() method of the package, which contains your tests.
//
// Example:
//  init() {
//    testza.SetShowStartupMessage(false) // Disable the startup message
//    testza.SetShowStartupMessage(true)  // Enable the startup message
//  }
func SetShowStartupMessage(show bool) {
	showStartupMessage = show
}

// SetEqualContextLineCount controls how many lines are shown around a changed diff line.
// If set to -1 it will show full diff.
// You should use this in the init() method of the package, which contains your tests.
//
// Example:
//  init() {
//    testza.SetEqualContextLineCount(-1) // Show all diff lines
//    testza.SetEqualContextLineCount(3)  // Show 3 lines around every changed line
//  }
func SetEqualContextLineCount(lines int) {
	internal.EqualContextLineCount = lines
}
