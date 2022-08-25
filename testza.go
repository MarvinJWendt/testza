/*
Package testza is a full-featured testing framework for Go. It integrates with the default test runner, so you can use it with the standard `go test` tool. Testza contains easy to use methods, like assertions, output capturing, mocking, and much more.
*/
package testza

import (
	"sync"

	"github.com/pterm/pterm"
)

var infoPrinter = pterm.DefaultSection.WithStyle(pterm.NewStyle(pterm.FgMagenta)).WithLevel(2).WithBottomPadding(0).WithTopPadding(0)
var secondary = pterm.LightCyan

var initSync sync.Mutex
