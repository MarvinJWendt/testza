package testza

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/MarvinJWendt/testza/internal"
	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
)

// SnapshotCreate creates a snapshot of an object, which can be validated in future test runs.
// Using this function directly will override previous snapshots with the same name.
// You most likely want to use SnapshotCreateOrValidate.
//
// NOTICE: \r\n will be replaced with \n to make the files consistent between operating systems.
//
// Example:
//  testza.SnapshotCreate(t.Name(), objectToBeSnapshotted)
func SnapshotCreate(name string, snapshotObject interface{}) error {
	dir := getCurrentScriptDirectory() + "/testdata/snapshots/"

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("creating snapshot failed: %w", err)
	}

	originalSpewConfig := spew.Config.DisablePointerAddresses
	spew.Config.DisablePointerAddresses = true
	dump := strings.ReplaceAll(spew.Sdump(snapshotObject), "\r\n", "\n")
	err = ioutil.WriteFile(path.Clean(dir+name+".testza"), []byte(dump), 0755)
	if err != nil {
		return fmt.Errorf("creating snapshot failed: %w", err)
	}
	spew.Config.DisablePointerAddresses = originalSpewConfig

	return nil
}

// SnapshotValidate validates an already exisiting snapshot of an object.
// You most likely want to use SnapshotCreateOrValidate.
//
// NOTICE: \r\n will be replaced with \n to make the files consistent between operating systems.
//
// Example:
//  testza.SnapshotValidate(t, t.Name(), objectToBeValidated)
//  testza.SnapshotValidate(t, t.Name(), objectToBeValidated, "Optional message")
func SnapshotValidate(t testRunner, name string, actual interface{}, msg ...interface{}) error {
	dir := getCurrentScriptDirectory() + "/testdata/snapshots/"
	snapshotPath := path.Clean(dir + name + ".testza")
	snapshotContent, err := ioutil.ReadFile(snapshotPath)
	snapshot := strings.ReplaceAll(string(snapshotContent), "\r\n", "\n")
	if err != nil {
		return fmt.Errorf("validating snapshot failed: %w", err)
	}
	originalSpewConfig := spew.Config.DisablePointerAddresses
	spew.Config.DisablePointerAddresses = true

	if spew.Sdump(actual) != snapshot {
		internal.Fail(t,
			generateMsg(msg,
				fmt.Sprintf("Snapshot '%s' failed to validate", name)),
			internal.Objects{
				{
					Name:      "Difference",
					NameStyle: pterm.NewStyle(pterm.FgLightYellow),
					Data:      internal.GetDifference(snapshot, spew.Sdump(actual), true),
					DataStyle: pterm.NewStyle(pterm.FgGreen),
					Raw:       true,
				},
				{
					Name:      "Expected",
					NameStyle: pterm.NewStyle(pterm.FgLightGreen),
					Data:      snapshot,
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

// SnapshotCreateOrValidate creates a snapshot of an object which can be used in future test runs.
// It is good practice to name your snapshots the same as the test they are created in.
// You can do that automatically by using t.Name() as the second parameter, if you are using the inbuilt test system of Go.
// If a snapshot already exists, the function will not create a new one, but validate the exisiting one.
// To re-create a snapshot, you can delete the according file in /testdata/snapshots/.
//
// NOTICE: \r\n will be replaced with \n to make the files consistent between operating systems.
//
// Example:
//  testza.SnapshotCreateOrValidate(t, t.Name(), object)
//  testza.SnapshotCreateOrValidate(t, t.Name(), object, "Optional Message")
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
