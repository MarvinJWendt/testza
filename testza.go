package testza

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/klauspost/cpuid/v2"
	"github.com/pterm/pterm"
)

var randomSeed int64
var infoPrinter = pterm.DefaultSection.WithStyle(pterm.NewStyle(pterm.FgMagenta)).WithLevel(2).WithBottomPadding(0).WithTopPadding(0)
var secondary = pterm.LightCyan

func init() {
	randomSeed = time.Now().UnixNano()
	rand.Seed(randomSeed)

	if showStartupMessage {
		infoPrinter.WithLevel(1).Println("Running tests with " + secondary("Testza"))
		infoPrinter.Printfln(`Using seed "%s" for random operations`, secondary(randomSeed))
		infoPrinter.Printfln(`System info: OS=%s | arch=%s | cpu=%s | go=%s`, secondary(runtime.GOOS), secondary(runtime.GOARCH), secondary(cpuid.CPU.BrandName), secondary(runtime.Version()))
		fmt.Println()
	}
}
