package testza

import (
	"fmt"
	"testing"

	"github.com/MarvinJWendt/testza/internal"
)

// FuzzBoolFull returns true and false in a boolean slice.
func FuzzBoolFull() []bool {
	return []bool{true, false}
}

// FuzzBoolRunTests runs a test for every value in a testset.
// You can use the value as input parameter for your functions, to sanity test against many different cases.
// This ensures that your functions have a correct error handling and enables you to test against hunderts of cases easily.
//
// Example:
//  testza.FuzzBoolRunTests(t, testza.FuzzBoolFull(), func(t *testing.T, index int, b bool) {
//  	// Test logic
//  	// err := YourFunction(b)
//  	// testza.AssertNoError(t, err)
//  	// ...
//  })
func FuzzBoolRunTests(t testRunner, testSet []bool, testFunc func(t *testing.T, index int, f bool)) {
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

// FuzzBoolModify returns a modified version of a test set.
//
// Example:
//  testset := testza.FuzzBoolModify(testza.FuzzBoolFull(), func(index int, value bool) bool {
//  	return !value
//  })
func FuzzBoolModify(inputSlice []bool, modifier func(index int, value bool) bool) (floats []bool) {
	for i, input := range inputSlice {
		floats = append(floats, modifier(i, input))
	}

	return
}
