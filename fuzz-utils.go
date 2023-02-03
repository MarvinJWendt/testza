package testza

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/MarvinJWendt/testza/internal"
)

// FuzzUtilMergeSets merges multiple test sets into one.
// All test sets must have the same type.
//
// Example:
//
//	mergedSet := testza.FuzzUtilMergeSets(testza.FuzzIntGenerateRandomNegative(3, 0), testza.FuzzIntGenerateRandomPositive(2, 0))
func FuzzUtilMergeSets[setType any](sets ...[]setType) (merged []setType) {
	for _, set := range sets {
		merged = append(merged, set...)
	}

	return merged
}

// FuzzUtilRunTests runs a test for every value in a test set.
// You can use the value as input parameter for your functions, to sanity test against many different cases.
// This ensures that your functions have a correct error handling and enables you to test against hundreds of cases easily.
//
// Example:
//
//	testza.FuzzUtilRunTests(t, testza.FuzzStringEmailAddresses(), func(t *testing.T, index int, emailAddress string) {
//		// Test logic
//		// err := YourFunction(emailAddress)
//		// testza.AssertNoError(t, err)
//		// ...
//	})
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

// FuzzUtilModifySet returns a modified version of a test set.
//
// Example:
//
//	 modifiedSet := testza.FuzzUtilModifySet(testza.FuzzIntFull(), func(i int, value int) int {
//			return i * 2 // double every value in the test set
//		})
func FuzzUtilModifySet[setType any](inputSet []setType, modifier func(index int, value setType) setType) (floats []setType) {
	for i, input := range inputSet {
		floats = append(floats, modifier(i, input))
	}

	return
}

// FuzzUtilLimitSet returns a random sample of a test set with a maximal size.
//
// Example:
//
//	limitedSet := testza.FuzzUtilLimitSet(testza.FuzzStringFull(), 10)
func FuzzUtilLimitSet[setType any](testSet []setType, max int) []setType {
	if len(testSet) <= max {
		return testSet
	}

	if max <= 0 {
		return []setType{}
	}

	rand.Shuffle(len(testSet), func(i, j int) { testSet[i], testSet[j] = testSet[j], testSet[i] })

	return testSet[:max]
}

// FuzzUtilDistinctSet returns a set with removed duplicates.
//
// Example:
//
//	uniqueSet := testza.FuzzUtilDistinctSet([]string{"A", "C", "A", "B", "A", "B", "C"})
//	// uniqueSet => []string{"A", "C", "B"}
func FuzzUtilDistinctSet[setType comparable](testSet []setType) []setType {
	seen := map[setType]bool{}
	var result []setType

	for _, v := range testSet {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}
