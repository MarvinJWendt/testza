package testza

import (
	"testing"

	"github.com/MarvinJWendt/testza/internal"
	"github.com/pterm/pterm"
)

func TestSetColorsEnabled(t *testing.T) {
	t.Run("Disable", func(t *testing.T) {
		SetColorsEnabled(false)
		AssertFalse(t, pterm.PrintColor)
	})

	t.Run("Enable", func(t *testing.T) {
		SetColorsEnabled(true)
		AssertTrue(t, pterm.PrintColor)
	})
}

func TestSetLineNumbersEnabled(t *testing.T) {
	t.Run("Disable", func(t *testing.T) {
		SetLineNumbersEnabled(false)
		AssertFalse(t, internal.LineNumbersEnabled)
	})

	t.Run("Enable", func(t *testing.T) {
		SetLineNumbersEnabled(true)
		AssertTrue(t, internal.LineNumbersEnabled)
	})
}
