package testza

import (
	"fmt"
	"testing"

	"github.com/MarvinJWendt/testza/internal"
)

// FuzzUtilModifySet returns a modified version of a test set.
//
// Example:
//  modifiedSet := testza.FuzzUtilModifySet(testza.FuzzIntFull(), func(i int, value int) int {
//		return i * 2 // double every value in the test set
//	})
func FuzzUtilModifySet[setType any](inputSet []setType, modifier func(index int, value setType) setType) (floats []setType) {
	for i, input := range inputSet {
		floats = append(floats, modifier(i, input))
	}

	return
}

// FuzzUtilRunTests runs a test for every value in a testset.
// You can use the value as input parameter for your functions, to sanity test against many different cases.
// This ensures that your functions have a correct error handling and enables you to test against hunderts of cases easily.
//
// Example:
//  testza.FuzzUtilRunTests(t, testza.FuzzStringEmailAddresses(), func(t *testing.T, index int, emailAddress string) {
//  	// Test logic
//  	// err := YourFunction(emailAddress)
//  	// testza.AssertNoError(t, err)
//  	// ...
//  })
func FuzzUtilRunTests[setType any](t testRunner, testSet []setType, testFunc func(t *testing.T, index int, f setType)) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	test := internal.GetTest(t)
	if test == nil {
		t.Error(internal.ErrCanNotRunIfNotBuiltinTesting)
		return
	}

	for i, v := range testSet {
		test.Run(fmt.Sprint(v), func(t *testing.T) {
			t.Helper()

			testFunc(t, i, v)
		})
	}
}
