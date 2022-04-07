package testza

import (
	"math"
	"math/rand"
)

// FuzzFloat64Full returns a combination of every float64 testset and some random float64s (positive and negative).
func FuzzFloat64Full() (floats []float64) {
	for i := 0; i < 50; i++ {
		floats = append(floats,
			FuzzFloat64GenerateRandomPositive(1, float64(i*1000))[0],
			FuzzFloat64GenerateRandomNegative(1, float64(i*1000*-1))[0],
		)
	}
	return
}

// FuzzFloat64GenerateRandomRange generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func FuzzFloat64GenerateRandomRange(count int, min, max float64) (floats []float64) {
	for i := 0; i < count; i++ {
		floats = append(floats, min+rand.Float64()*(max-min))
	}

	return
}

// FuzzFloat64GenerateRandomPositive generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func FuzzFloat64GenerateRandomPositive(count int, max float64) (floats []float64) {
	if max <= 0 {
		max = math.MaxFloat64
	}

	floats = append(floats, FuzzFloat64GenerateRandomRange(count, 0, max)...)

	return
}

// FuzzFloat64GenerateRandomNegative generates random negative integers with a minimum of min.
// If the minimum is positive, it will be converted to a negative number.
// If it is set to 0, there is no limit.
func FuzzFloat64GenerateRandomNegative(count int, min float64) (floats []float64) {
	if min > 0 {
		min *= -1
	} else if min == 0 {
		min = math.MaxFloat64 * -1
	}

	floats = append(floats, FuzzFloat64GenerateRandomRange(count, min, 0)...)

	return
}
