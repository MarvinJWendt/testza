package testza_test

import (
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

type snapshotComplexObjectType struct {
	Name     string
	Username string
	Birthday time.Time
	Nested   struct {
		ChildName    string
		PointerField *string
		Set          []int
	}
}

var snapshotObject = snapshotObjectType{
	Name:     "Marvin Wendt",
	Username: "MarvinJWendt",
	Birthday: time.Date(2001, time.January, 24, 0, 0, 0, 0, time.UTC),
}
var snapshotName = "TestSnapshotCreate_file_content"

var snapshotString = `hello world`
var snapshotNameString = "TestSnapshotCreate_file_content_string"

var snapshotComplexObject = snapshotComplexObjectType{
	Name:     "Marvin Wendt",
	Username: "MarvinJWendt",
	Birthday: time.Date(2001, time.January, 24, 0, 0, 0, 0, time.UTC),
	Nested: struct {
		ChildName    string
		PointerField *string
		Set          []int
	}{
		ChildName:    "as yet untitled",
		PointerField: &snapshotString,
		Set:          []int{1, 2, 3, 4, 3, 2, 1},
	},
}

func TestSnapshotCreate_file_content(t *testing.T) {
	err := testza.SnapshotCreate(snapshotName, snapshotObject)
	testza.AssertNoError(t, err)

	snapshotContent, err := os.ReadFile(internal.GetCurrentScriptDirectory() + "/testdata/snapshots/" + t.Name() + ".testza")
	testza.AssertNoError(t, err)

	testza.AssertEqual(t, spew.Sdump(snapshotObject), string(snapshotContent))
}

func TestSnapshotCreate_file_content_string(t *testing.T) {
	err := testza.SnapshotCreate(snapshotNameString, snapshotString)
	testza.AssertNoError(t, err)

	snapshotContent, err := os.ReadFile(internal.GetCurrentScriptDirectory() + "/testdata/snapshots/" + t.Name() + ".testza")
	testza.AssertNoError(t, err)

	testza.AssertEqual(t, spew.Sdump(snapshotString), string(snapshotContent))
}

func TestSnapshotValidate(t *testing.T) {
	err := testza.SnapshotValidate(t, snapshotName, snapshotObject)
	testza.AssertNoError(t, err)

	err = testza.SnapshotValidate(t, snapshotNameString, snapshotString)
	testza.AssertNoError(t, err)
}

func TestSnapshotValidate_fails(t *testing.T) {
	modifiedSnapshotObject := snapshotObject
	modifiedSnapshotObject.Username = "NotMarvinJWendt"
	testza.AssertTestFails(t, func(t testza.TestingPackageWithFailFunctions) {
		err := testza.SnapshotValidate(t, snapshotName, modifiedSnapshotObject)
		testza.AssertNoError(t, err)
	})

	testza.AssertTestFails(t, func(t testza.TestingPackageWithFailFunctions) {
		err := testza.SnapshotValidate(t, snapshotNameString, `foo bar`)
		testza.AssertNoError(t, err)
	})
}

func TestSnapshotCreateOrValidate(t *testing.T) {
	err := testza.SnapshotCreateOrValidate(t, t.Name(), snapshotObject)
	testza.AssertNoError(t, err)
}

func TestSnapshotCreateOrValidate_complex_object(t *testing.T) {
	err := testza.SnapshotCreateOrValidate(t, t.Name(), snapshotComplexObject)
	testza.AssertNoError(t, err)
}

func TestSnapshotCreateOrValidate_create_random(t *testing.T) {
	name := t.Name() + testza.FuzzStringGenerateRandom(1, rand.Intn(10))[0]
	err := testza.SnapshotCreateOrValidate(t, name, snapshotObject)
	testza.AssertNoError(t, err)

	err = os.Remove(internal.GetCurrentScriptDirectory() + "/testdata/snapshots/" + name + ".testza")
	testza.AssertNoError(t, err)
}

func TestSnapshotCreateOrValidate_invalid_name(t *testing.T) {
	err := testza.SnapshotCreateOrValidate(t, string(rune(0))+"><", snapshotObject)
	testza.AssertNotNil(t, err)
}

func TestSnapshotCreateOrValidate_nested_test_name(t *testing.T) {
	snapshotNestedDirName := t.Name()
	snapshotName := snapshotNestedDirName + "/nested_name"

	snapshotFullPath := internal.GetCurrentScriptDirectory() + "/testdata/snapshots/" + snapshotNestedDirName

	// ensure snapshot does not exist - remove snapshot and empty parent directory
	if _, err := os.Stat(snapshotFullPath); err == nil {
		_ = os.RemoveAll(snapshotFullPath)
	}

	err := testza.SnapshotCreateOrValidate(t, snapshotName, "snapshot-data")

	// try to clean up before the assert so we leave the working copy clean
	if _, err := os.Stat(snapshotFullPath); err == nil {
		_ = os.RemoveAll(snapshotFullPath)
	}

	testza.AssertNoError(t, err)
}

func TestSnapshotCreateOrValidate_key_order_stability_with_map(t *testing.T) {
	// unsorted stable test data
	data := []string{
		"257ad566-2f5e-431e-b50a-d30c64849fb1",
		"67b2fc0f-ab4a-4305-b240-deff819bb080",
		"5778dc84-8a74-486d-ac08-ec7ff2f53348",
		"b27e2155-af0a-46a5-ac8c-73f6d497c95c",
		"b822f4da-57bc-4057-8c7f-7f63e71752c5",
		"8227c74f-e621-44b8-be21-5a53f3504659",
		"9021c878-425d-4b26-91b8-fa4c670ad111",
		"dfc0583a-f36c-4b75-9277-98796cf5421a",
		"5137c065-a216-45b7-a0c9-58b9425c1416",
		"0dea5e9e-361c-4c17-87cc-fee3e41dc92d",
		"619bbea9-bd2e-4a19-b085-33a2677ced7e",
		"45e032f2-8eaf-4d17-9a8d-55688da19104",
		"7591cb03-b561-4ed9-8d3f-76ee0f578d34",
		"12787615-054e-4b09-b7ea-b81a6d30a0a1",
		"a3563ca3-daaf-4591-afc0-f393c8b66993",
		"e83adea8-0cd0-4fb8-843a-8f438f9f6054",
		"78f2ab90-084b-4afc-801d-5d06d00f0761",
		"a240ecc3-f77a-4a8b-922a-b9cde7ac7416",
		"fc54c8c7-9a30-4b9f-a9c3-bd97df9e899b",
		"9f58c6b7-9555-458c-9eb9-8ff400d53819",
	}

	// Create a map from unsorted data, where the value is the insertion order.
	// Ordering stability of maps is not guaranteed.
	snapshotMap := make(map[string]int, len(data))
	for i, d := range data {
		snapshotMap[d] = i
	}

	err := testza.SnapshotCreateOrValidate(t, t.Name(), snapshotMap)
	testza.AssertNoError(t, err)
}
