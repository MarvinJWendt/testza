package testza

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/MarvinJWendt/testza/internal"
)

// MockInputsIntsHelper contains integer test sets.
type MockInputsIntsHelper struct{}

// Full returns a combination of every integer testset and some random integers (positive and negative).
func (h MockInputsIntsHelper) Full() (ints []int) {
	for i := 0; i < 50; i++ {
		ints = append(ints,
			h.GenerateRandomPositive(1, i*1000)[0],
			h.GenerateRandomNegative(1, i*1000*-1)[0],
		)
	}
	return
}

// GenerateRandomRange generates random integers with a range of min to max.
func (h MockInputsIntsHelper) GenerateRandomRange(count, min, max int) (ints []int) {
	for i := 0; i < count; i++ {
		ints = append(ints, rand.Intn(max-min)+min)
	}

	return
}

// GenerateRandomPositive generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func (h MockInputsIntsHelper) GenerateRandomPositive(count, max int) (ints []int) {
	if max <= 0 {
		max = math.MaxInt64
	}

	ints = append(ints, h.GenerateRandomRange(count, 1, max)...)

	return
}

// GenerateRandomNegative generates random negative integers with a minimum of min.
// If the minimum is 0, or above, the maximum will be set to math.MinInt64.
func (h MockInputsIntsHelper) GenerateRandomNegative(count, min int) (ints []int) {
	if min >= 0 {
		min = math.MinInt64
	}

	min = int(math.Abs(float64(min)))

	randomPositives := h.GenerateRandomPositive(count, min)

	for _, p := range randomPositives {
		ints = append(ints, p*-1)
	}

	return
}

// RunTests runs a test for every value in a testset.
// You can use the value as input parameter for your functions, to sanity test against many different cases.
// This ensures that your functions have a correct error handling and enables you to test against hunderts of cases easily.
func (s MockInputsIntsHelper) RunTests(t testRunner, testSet []int, testFunc func(t *testing.T, index int, i int)) {
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
func (h MockInputsIntsHelper) Modify(inputSlice []int, f func(index int, value int) int) (ints []int) {
	for i, input := range inputSlice {
		ints = append(ints, f(i, input))
	}

	return
}
