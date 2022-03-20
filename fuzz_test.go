package testza_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/MarvinJWendt/testza"
)

func TestFuzzInputStringModify(t *testing.T) {
	stringSlice := []string{"Hello", "World", "TeSt"}
	expected := []string{"hello", "world", "test"}

	input := FuzzInputStringModify(stringSlice, func(index int, value string) string {
		return strings.ToLower(value)
	})

	AssertEqual(t, expected, input)
}

func TestFuzzInputStringLimit(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Limit=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(FuzzInputStringLimit(FuzzInputStringFull(), i)))
		})
	}
}

func TestFuzzInputStringGenerateRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Length=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(FuzzInputStringGenerateRandom(1, i)[0]))
		})
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Count=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(FuzzInputStringGenerateRandom(i, 5)))
		})
	}
}

func TestFuzzInputStringNumeric(t *testing.T) {
	for _, v := range FuzzInputStringNumeric() {
		t.Run(v, func(t *testing.T) {
			f, err := strconv.ParseFloat(v, 64)
			AssertNumeric(t, f)
			AssertNoError(t, err)
		})
	}
}

func TestFuzzInputStringFull(t *testing.T) {
	AssertGreater(t, len(FuzzInputStringFull()), 0)
}

func TestFuzzInputBoolFull(t *testing.T) {
	AssertEqual(t, []bool{true, false}, FuzzInputBoolFull())
}

func TestFuzzInputBoolRunTests(t *testing.T) {
	FuzzInputBoolRunTests(t, FuzzInputBoolFull(), func(t *testing.T, index int, f bool) {
		AssertNotNil(t, f)
	})
}

func TestFuzzInputIntGenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("GenerateRandomPositive generates positive numbers only", func(t *testing.T) {
			AssertGreater(t, FuzzInputIntGenerateRandomPositive(1, 100)[0], 0)
		})

		t.Run("GenerateRandomNegative generates negative numbers only", func(t *testing.T) {
			AssertLess(t, FuzzInputIntGenerateRandomNegative(1, 100)[0], 0)
		})
	}
}

func TestFuzzInputIntFull(t *testing.T) {
	AssertGreater(t, len(FuzzInputIntFull()), 0)
}

func TestFuzzInputIntRunTests(t *testing.T) {
	FuzzInputIntRunTests(t, FuzzInputIntFull(), func(t *testing.T, index int, f int) {
		AssertNotNil(t, f)
	})
}

func TestFuzzInputIntModify(t *testing.T) {
	testSet := FuzzInputIntFull()
	s := FuzzInputIntModify(testSet, func(index int, value int) int {
		return value * -1
	})

	for i, f := range testSet {
		t.Run("Number should be inverted", func(t *testing.T) {
			AssertEqual(t, f, s[i]*-1)
		})
	}
}

func TestFuzzInputFloat64Full(t *testing.T) {
	AssertGreater(t, len(FuzzInputFloat64Full()), 0)
}

func TestFuzzInputFloat64RunTests(t *testing.T) {
	FuzzInputFloat64RunTests(t, FuzzInputFloat64Full(), func(t *testing.T, index int, f float64) {
		AssertNotNil(t, f)
	})
}

func TestFuzzInputFloat64GenerateRandomNegative(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := FuzzInputFloat64GenerateRandomNegative(1, 0)[0]
		t.Run(fmt.Sprintf("%v must be negative", n), func(t *testing.T) {
			AssertLess(t, n, 0)
		})
	}
}

func TestFuzzInputFloat64GenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := FuzzInputFloat64GenerateRandomPositive(1, 0)[0]
		t.Run(fmt.Sprintf("%v must be positive", n), func(t *testing.T) {
			AssertGreater(t, n, 0)
		})
	}
}

func TestFuzzInputFloat64Modify(t *testing.T) {
	testSet := FuzzInputFloat64Full()
	s := FuzzInputFloat64Modify(testSet, func(index int, value float64) float64 {
		return value * -1
	})

	for i, f := range testSet {
		t.Run("Number should be inverted", func(t *testing.T) {
			AssertEqual(t, f, s[i]*-1)
		})
	}
}

func TestFuzzInputBoolModify(t *testing.T) {
	s := FuzzInputBoolModify(FuzzInputBoolFull(), func(index int, value bool) bool {
		return !value
	})
	AssertEqual(t, []bool{false, true}, s)
}
