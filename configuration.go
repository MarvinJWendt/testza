package testza

import (
	"flag"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/MarvinJWendt/testza/internal"
	"github.com/pterm/pterm"
)

var showStartupMessage = true

func init() {
	// Defining flags to show up in the help message
	flag.Bool("testza.disable-color", false, "disables colored output")
	flag.Bool("testza.disable-line-numbers", false, "disables line numbers in output")
	flag.Bool("testza.disable-startup-message", false, "disable the startup message")
	flag.Int64("testza.seed", 0, "seed used for random operations")
	flag.Int("testza.diff-context-lines", 2, "sets the context line count in difference output")

	for i, arg := range os.Args {
		// Check if the argument is a flag
		if !strings.HasPrefix(arg, "--") {
			continue
		}

		// Figure out if the flag has a value
		var value string
		if i != len(os.Args)-1 {
			value = os.Args[i+1]
			if strings.HasPrefix(value, "-") {
				value = ""
			}
		}

		// Check for set flags and run the appropriate function
		switch strings.TrimPrefix(arg, "--testza.") {
		case "disable-color":
			SetColorsEnabled(false)
		case "disable-line-numbers":
			SetLineNumbersEnabled(false)
		case "disable-startup-message":
			SetShowStartupMessage(false)
		case "seed":
			seed, err := strconv.Atoi(value)
			pterm.Fatal.PrintOnError(err)
			SetRandomSeed(int64(seed))
		case "diff-context-lines":
			v, err := strconv.Atoi(value)
			pterm.Fatal.PrintOnError(err)
			SetDiffContextLines(v)
		}
	}
}

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

// SetDiffContextLines controls how many lines are shown around a changed diff line.
// If set to -1 it will show full diff.
// You should use this in the init() method of the package, which contains your tests.
//
// Example:
//  init() {
//    testza.SetDiffContextLines(-1) // Show all diff lines
//    testza.SetDiffContextLines(3)  // Show 3 lines around every changed line
//  }
func SetDiffContextLines(lines int) {
	internal.DiffContextLines = lines
}
