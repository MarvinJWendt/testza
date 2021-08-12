package testza_test

import (
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/MarvinJWendt/testza"
	"github.com/MarvinJWendt/testza/internal"
	"github.com/davecgh/go-spew/spew"
)

type snapshotObjectType struct {
	Name     string
	Username string
	Birthday time.Time
}

var snapshotObject = snapshotObjectType{
	Name:     "Marvin Wendt",
	Username: "MarvinJWendt",
	Birthday: time.Date(2001, time.January, 24, 0, 0, 0, 0, time.UTC),
}
var snapshotName = "TestSnapshotCreate_file_content"

func TestSnapshotCreate_file_content(t *testing.T) {
	err := testza.SnapshotCreate(snapshotName, snapshotObject)
	testza.AssertNoError(t, err)

	snapshotContent, err := ioutil.ReadFile(internal.GetCurrentScriptDirectory() + "/testdata/snapshots/" + t.Name() + ".testza")
	testza.AssertNoError(t, err)

	testza.AssertEqual(t, spew.Sdump(snapshotObject), string(snapshotContent))
}

func TestSnapshotValidate(t *testing.T) {
	err := testza.SnapshotValidate(t, snapshotName, snapshotObject)
	testza.AssertNoError(t, err)
}

func TestSnapshotValidate_fails(t *testing.T) {
	modifiedSnapshotObject := snapshotObject
	modifiedSnapshotObject.Username = "NotMarvinJWendt"
	testza.AssertTestFails(t, func(t testza.TestingPackageWithFailFunctions) {
		err := testza.SnapshotValidate(t, snapshotName, modifiedSnapshotObject)
		testza.AssertNoError(t, err)
	})
}

func TestSnapshotCreateOrValidate(t *testing.T) {
	err := testza.SnapshotCreateOrValidate(t, t.Name(), snapshotObject)
	testza.AssertNoError(t, err)
}

func TestSnapshotCreateOrValidate_create_random(t *testing.T) {
	name := t.Name() + testza.MockInputStringGenerateRandom(1, rand.Intn(10))[0]
	err := testza.SnapshotCreateOrValidate(t, name, snapshotObject)
	testza.AssertNoError(t, err)

	err = os.Remove(internal.GetCurrentScriptDirectory() + "/testdata/snapshots/" + name + ".testza")
	testza.AssertNoError(t, err)
}

func TestSnapshotCreateOrValidate_invalid_name(t *testing.T) {
	err := testza.SnapshotCreateOrValidate(t, string(rune(0))+"><", snapshotObject)
	testza.AssertNotNil(t, err)
}
