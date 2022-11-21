package testza

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
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
//
//	testza.SnapshotCreate(t.Name(), objectToBeSnapshotted)
func SnapshotCreate(name string, snapshotObject any) error {
	dir := getCurrentScriptDirectory() + "/testdata/snapshots/"
	return snapshotCreateForDir(dir, name, snapshotObject)
}

func snapshotCreateForDir(dir string, name string, snapshotObject any) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("creating snapshot failed: %w", err)
	}

	dump := strings.ReplaceAll(createSnapshotText(snapshotObject), "\r\n", "\n")
	err = os.WriteFile(path.Clean(dir+name+".testza"), []byte(dump), 0755)
	if err != nil {
		return fmt.Errorf("creating snapshot failed: %w", err)
	}

	return nil
}

// SnapshotValidate validates an already exisiting snapshot of an object.
// You most likely want to use SnapshotCreateOrValidate.
//
// NOTICE: \r\n will be replaced with \n to make the files consistent between operating systems.
//
// Example:
//
//	testza.SnapshotValidate(t, t.Name(), objectToBeValidated)
//	testza.SnapshotValidate(t, t.Name(), objectToBeValidated, "Optional message")
func SnapshotValidate(t testRunner, name string, actual any, msg ...any) error {
	dir := getCurrentScriptDirectory() + "/testdata/snapshots/"
	return snapshotValidateFromDir(dir, t, name, actual, msg...)
}

var snapshotStringMatcher = regexp.MustCompile(`(?m)^\(.+?\)\s\(len=\d+\)\s(".+")$`)

func snapshotValidateFromDir(dir string, t testRunner, name string, actual any, msg ...any) error {
	snapshotPath := path.Clean(dir + name + ".testza")
	snapshotContent, err := os.ReadFile(snapshotPath)
	if err != nil {
		return fmt.Errorf("validating snapshot failed: %w", err)
	}
	snapshot := strings.ReplaceAll(string(snapshotContent), "\r\n", "\n")

	actualSnapshot := createSnapshotText(actual)

	if actualSnapshot != snapshot {
		var diffObject *internal.Object
		if strActual, ok := actual.(string); ok {
			if match := snapshotStringMatcher.FindStringSubmatch(snapshot); len(match) > 0 {
				if unquoted, err := strconv.Unquote(match[1]); err == nil {
					object := internal.NewDiffObject(unquoted, strActual, true)
					diffObject = &object
				}
			}
		}

		if diffObject == nil {
			object := internal.NewDiffObject(snapshot, actualSnapshot, true)
			diffObject = &object
		}

		internal.Fail(t,
			generateMsg(msg,
				fmt.Sprintf("Snapshot '%s' failed to validate", name)),
			internal.Objects{
				*diffObject,
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
					Data:      actualSnapshot,
					DataStyle: pterm.NewStyle(pterm.FgRed),
					Raw:       true,
				},
			})
	}

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
//
//	testza.SnapshotCreateOrValidate(t, t.Name(), object)
//	testza.SnapshotCreateOrValidate(t, t.Name(), object, "Optional Message")
func SnapshotCreateOrValidate(t testRunner, name string, object any, msg ...any) error {
	dir := getCurrentScriptDirectory() + "/testdata/snapshots/"
	snapshotPath := path.Clean(dir + name + ".testza")
	if strings.Contains(name, "/") {
		err := os.MkdirAll(path.Dir(snapshotPath), 0755)
		if err != nil {
			return fmt.Errorf("creating snapshot directories failed: %w", err)
		}
	}

	if _, err := os.Stat(snapshotPath); err == nil {
		err = snapshotValidateFromDir(dir, t, name, object, msg...)
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		err = snapshotCreateForDir(dir, name, object)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func createSnapshotText(object any) string {
	cfg := spew.NewDefaultConfig()

	cfg.DisablePointerAddresses = true
	cfg.DisableCapacities = true
	cfg.SortKeys = true

	return cfg.Sdump(object)
}
