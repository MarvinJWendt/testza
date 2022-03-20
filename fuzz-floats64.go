package testza

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/MarvinJWendt/testza/internal"
)

// FuzzInputFloat64Full returns a combination of every float64 testset and some random float64s (positive and negative).
func FuzzInputFloat64Full() (floats []float64) {
	for i := 0; i < 50; i++ {
		floats = append(floats,
			FuzzInputFloat64GenerateRandomPositive(1, float64(i*1000))[0],
			FuzzInputFloat64GenerateRandomNegative(1, float64(i*1000*-1))[0],
		)
	}
	return
}

// FuzzInputFloat64GenerateRandomRange generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func FuzzInputFloat64GenerateRandomRange(count int, min, max float64) (floats []float64) {
	for i := 0; i < count; i++ {
		floats = append(floats, min+rand.Float64()*(max-min))
	}

	return
}

// FuzzInputFloat64GenerateRandomPositive generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func FuzzInputFloat64GenerateRandomPositive(count int, max float64) (floats []float64) {
	if max <= 0 {
		max = math.MaxFloat64
	}

	floats = append(floats, FuzzInputFloat64GenerateRandomRange(count, 0, max)...)

	return
}

// FuzzInputFloat64GenerateRandomNegative generates random negative integers with a minimum of min.
// If the minimum is positive, it will be converted to a negative number.
// If it is set to 0, there is no limit.
func FuzzInputFloat64GenerateRandomNegative(count int, min float64) (floats []float64) {
	if min > 0 {
		min *= -1
	} else if min == 0 {
		min = math.MaxFloat64 * -1
	}

	floats = append(floats, FuzzInputFloat64GenerateRandomRange(count, min, 0)...)

	return
}

// FuzzInputFloat64RunTests runs a test for every value in a testset.
// You can use the value as input parameter for your functions, to sanity test against many different cases.
// This ensures that your functions have a correct error handling and enables you to test against hunderts of cases easily.
//
// Example:
//  testza.FuzzInputFloat64RunTests(t, testza.FuzzInputFloat64Full(), func(t *testing.T, index int, f float64) {
//  	// Test logic
//  	// err := YourFunction(f)
//  	// testza.AssertNoError(t, err)
//  	// ...
//  })
func FuzzInputFloat64RunTests(t testRunner, testSet []float64, testFunc func(t *testing.T, index int, f float64)) {
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

// FuzzInputFloat64Modify returns a modified version of a test set.
//
// Example:
//  testset := testza.FuzzInputFloat64Modify(testza.FuzzInputFloat64Full(), func(index int, value float64) float64 {
//  	return value * 2
//  })
func FuzzInputFloat64Modify(inputSlice []float64, modifier func(index int, value float64) float64) (floats []float64) {
	for i, input := range inputSlice {
		floats = append(floats, modifier(i, input))
	}

	return
}
