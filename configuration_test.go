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
		AssertFalse(t, GetColorsEnabled())
	})

	t.Run("Enable", func(t *testing.T) {
		SetColorsEnabled(true)
		AssertTrue(t, pterm.PrintColor)
		AssertTrue(t, GetColorsEnabled())
	})
}

func TestSetLineNumbersEnabled(t *testing.T) {
	t.Run("Disable", func(t *testing.T) {
		SetLineNumbersEnabled(false)
		AssertFalse(t, internal.LineNumbersEnabled)
		AssertFalse(t, GetLineNumbersEnabled())
	})

	t.Run("Enable", func(t *testing.T) {
		SetLineNumbersEnabled(true)
		AssertTrue(t, internal.LineNumbersEnabled)
		AssertTrue(t, GetLineNumbersEnabled())
	})
}

func TestSetRandomSeed(t *testing.T) {
	SetRandomSeed(1337)
	AssertEqual(t, int64(1337), randomSeed)
	AssertEqual(t, int64(1337), GetRandomSeed())
	AssertEqual(t, "4U390O49B9", FuzzStringGenerateRandom(1, 10)[0])
	un := time.Now().UnixNano()
	SetRandomSeed(un)
	AssertEqual(t, un, randomSeed)
	AssertEqual(t, un, GetRandomSeed())
}

func TestSetShowStartupMessage(t *testing.T) {
	t.Run("Default is true", func(t *testing.T) {
		AssertTrue(t, showStartupMessage)
		AssertTrue(t, GetShowStartupMessage())
	})

	t.Run("Set to false", func(t *testing.T) {
		SetShowStartupMessage(false)
		AssertFalse(t, showStartupMessage)
		AssertFalse(t, GetShowStartupMessage())
	})

	t.Run("Set to true", func(t *testing.T) {
		SetShowStartupMessage(true)
		AssertTrue(t, showStartupMessage)
		AssertTrue(t, GetShowStartupMessage())
	})
}

func TestSetEqualContextLineCount(t *testing.T) {
	t.Run("Default is 2", func(t *testing.T) {
		AssertEqual(t, 2, internal.DiffContextLines)
		AssertEqual(t, 2, GetDiffContextLines())
	})

	t.Run("Set to -1", func(t *testing.T) {
		SetDiffContextLines(-1)
		AssertEqual(t, -1, internal.DiffContextLines)
		AssertEqual(t, -1, GetDiffContextLines())
	})

	t.Run("Set to 3", func(t *testing.T) {
		SetDiffContextLines(3)
		AssertEqual(t, 3, internal.DiffContextLines)
		AssertEqual(t, 3, GetDiffContextLines())
	})

	SetDiffContextLines(2)
}
