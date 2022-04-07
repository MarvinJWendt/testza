package testza

import (
	"math"
	"math/rand"
)

// FuzzIntFull returns a combination of every integer testset and some random integers (positive and negative).
func FuzzIntFull() (ints []int) {
	for i := 0; i < 50; i++ {
		ints = append(ints,
			FuzzIntGenerateRandomPositive(1, i*1000)[0],
			FuzzIntGenerateRandomNegative(1, i*1000*-1)[0],
		)
	}
	return
}

// FuzzIntGenerateRandomRange generates random integers with a range of min to max.
func FuzzIntGenerateRandomRange(count, min, max int) (ints []int) {
	for i := 0; i < count; i++ {
		ints = append(ints, rand.Intn(max-min)+min)
	}

	return
}

// FuzzIntGenerateRandomPositive generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func FuzzIntGenerateRandomPositive(count, max int) (ints []int) {
	if max <= 0 {
		max = math.MaxInt64
	}

	ints = append(ints, FuzzIntGenerateRandomRange(count, 1, max)...)

	return
}

// FuzzIntGenerateRandomNegative generates random negative integers with a minimum of min.
// If the minimum is 0, or above, the maximum will be set to math.MinInt64.
func FuzzIntGenerateRandomNegative(count, min int) (ints []int) {
	if min >= 0 {
		min = math.MinInt64
	}

	min = int(math.Abs(float64(min)))

	randomPositives := FuzzIntGenerateRandomPositive(count, min)

	for _, p := range randomPositives {
		ints = append(ints, p*-1)
	}

	return
}
