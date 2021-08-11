package testza

import (
	"fmt"
	"os"
	"path"

	"github.com/MarvinJWendt/testza/internal"
	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
)

// SnapshotCreate creates a snapshot of an object, which can be validated in future test runs.
func SnapshotCreate(name string, snapshotObject interface{}) error {
	dir := getCurrentScriptDirectory() + "/testdata/snapshots/"

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("creating snapshot failed: %w", err)
	}

	originalSpewConfig := spew.Config.DisablePointerAddresses
	spew.Config.DisablePointerAddresses = true
	err = os.WriteFile(path.Clean(dir+name+".testza"), []byte(spew.Sdump(snapshotObject)), 0755)
	if err != nil {
		return fmt.Errorf("creating snapshot failed: %w", err)
	}
	spew.Config.DisablePointerAddresses = originalSpewConfig

	return nil
}

func SnapshotValidate(t testRunner, name string, actual interface{}, msg ...interface{}) error {
	dir := getCurrentScriptDirectory() + "/testdata/snapshots/"
	snapshotPath := path.Clean(dir + name + ".testza")
	snapshot, err := os.ReadFile(snapshotPath)
	if err != nil {
		return fmt.Errorf("validating snapshot failed: %w", err)
	}
	originalSpewConfig := spew.Config.DisablePointerAddresses
	spew.Config.DisablePointerAddresses = true

	if spew.Sdump(actual) != string(snapshot) {
		internal.Fail(t,
			generateMsg(msg,
				fmt.Sprintf("Snapshot '%s' failed to validate", name)),
			internal.Objects{
				{
					Name:      "Difference",
					NameStyle: pterm.NewStyle(pterm.FgLightYellow),
					Data:      internal.GetDifference(spew.Sdump(actual), string(snapshot), true),
					DataStyle: pterm.NewStyle(pterm.FgGreen),
					Raw:       true,
				},
				{
					Name:      "Expected",
					NameStyle: pterm.NewStyle(pterm.FgLightGreen),
					Data:      string(snapshot),
					DataStyle: pterm.NewStyle(pterm.FgGreen),
					Raw:       true,
				},
				{
					Name:      "Actual",
					NameStyle: pterm.NewStyle(pterm.FgLightRed),
					Data:      spew.Sdump(actual),
					DataStyle: pterm.NewStyle(pterm.FgRed),
					Raw:       true,
				},
			})
	}

	spew.Config.DisablePointerAddresses = originalSpewConfig

	return nil
}

func SnapshotCreateOrValidate(t testRunner, name string, object interface{}, msg ...interface{}) error {
	dir := getCurrentScriptDirectory() + "/testdata/snapshots/"
	snapshotPath := path.Clean(dir + name + ".testza")

	if _, err := os.Stat(snapshotPath); err == nil {
		err = SnapshotValidate(t, name, object, msg...)
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		err = SnapshotCreate(name, object)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}
