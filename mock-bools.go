package testza

import (
	"fmt"
	"testing"

	"github.com/MarvinJWendt/testza/internal"
)

type MockInputsBoolsHelper struct{}

// Full returns true and false in a boolean slice.
func (MockInputsBoolsHelper) Full() []bool {
	return []bool{true, false}
}

// RunTests runs a test for every value in a testset.
// You can use the value as input parameter for your functions, to sanity test against many different cases.
// This ensures that your functions have a correct error handling and enables you to test against hunderts of cases easily.
func (s MockInputsBoolsHelper) RunTests(t testRunner, testSet []bool, testFunc func(t *testing.T, index int, f bool)) {
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

// Modify returns a modified version of a test set.
func (h MockInputsBoolsHelper) Modify(inputSlice []bool, f func(index int, value bool) bool) (floats []bool) {
	for i, input := range inputSlice {
		floats = append(floats, f(i, input))
	}

	return
}
