package testza

import (
	"math"
	"math/rand"
)

// IntsHelper contains integer test sets.
type IntsHelper struct{}

func (h IntsHelper) Full() (ints []int) {
	for i := 0; i < 50; i++ {
		ints = append(ints,
			h.GenerateRandomPositive(1, i*1000)[0],
			h.GenerateRandomNegative(1, i*1000*-1)[0],
		)
	}
	return
}

// GenerateRandomPositive generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func (h IntsHelper) GenerateRandomPositive(count, max int) (ints []int) {
	if max <= 0 {
		max = math.MaxInt64
	}

	for i := 0; i < count; i++ {
		ints = append(ints, rand.Intn(max))
	}

	return
}

// GenerateRandomNegative generates random negative integers with a minimum of min.
// If the minimum is 0, or above, the maximum will be set to math.MinInt64.
func (h IntsHelper) GenerateRandomNegative(count, min int) (ints []int) {
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

// Modify returns a modified version of a test set.
func (h IntsHelper) Modify(inputSlice []int, f func(index int, value int) int) (ints []int) {
	for i, input := range inputSlice {
		ints = append(ints, f(i, input))
	}

	return
}
