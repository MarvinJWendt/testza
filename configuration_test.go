package testza

import (
	"testing"
	"time"

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

func TestSetRandomSeed(t *testing.T) {
	SetRandomSeed(1337)
	AssertEqual(t, int64(1337), randomSeed)
	AssertEqual(t, "4U390O49B9", FuzzInputStringGenerateRandom(1, 10)[0])
	un := time.Now().UnixNano()
	SetRandomSeed(un)
	AssertEqual(t, un, randomSeed)
}

func TestSetShowStartupMessage(t *testing.T) {
	t.Run("Default is true", func(t *testing.T) {
		AssertTrue(t, showStartupMessage)

	})

	t.Run("Set to false", func(t *testing.T) {
		SetShowStartupMessage(false)
		AssertFalse(t, showStartupMessage)
	})

	t.Run("Set to true", func(t *testing.T) {
		SetShowStartupMessage(true)
		AssertTrue(t, showStartupMessage)
	})
}
