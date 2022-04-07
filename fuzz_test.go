package testza_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/MarvinJWendt/testza"
)

func TestFuzzUtilModify(t *testing.T) {
	t.Run("String Slice", func(t *testing.T) {
		slice := []string{"Hello", "World", "TeSt"}
		expected := []string{"hello", "world", "test"}
		input := FuzzUtilModifySet(slice, func(index int, value string) string {
			return strings.ToLower(value)
		})
		AssertEqual(t, expected, input)
	})

	t.Run("Int Slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := []int{2, 4, 6, 8, 10}
		input := FuzzUtilModifySet(slice, func(index int, value int) int {
			return value * 2
		})
		AssertEqual(t, expected, input)
	})

	t.Run("Bool Slice", func(t *testing.T) {
		slice := []bool{true, false}
		expected := []bool{false, true}
		input := FuzzUtilModifySet(slice, func(index int, value bool) bool {
			return !value
		})
		AssertEqual(t, expected, input)
	})
}

func TestFuzzStringLimit(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Limit=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(FuzzStringLimit(FuzzStringFull(), i)))
		})
	}
}

func TestFuzzStringGenerateRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Length=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(FuzzStringGenerateRandom(1, i)[0]))
		})
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Count=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(FuzzStringGenerateRandom(i, 5)))
		})
	}
}

func TestFuzzStringNumeric(t *testing.T) {
	for _, v := range FuzzStringNumeric() {
		t.Run(v, func(t *testing.T) {
			f, err := strconv.ParseFloat(v, 64)
			AssertNumeric(t, f)
			AssertNoError(t, err)
		})
	}
}

func TestFuzzStringFull(t *testing.T) {
	AssertGreater(t, len(FuzzStringFull()), 0)
}

func TestFuzzBoolFull(t *testing.T) {
	AssertEqual(t, []bool{true, false}, FuzzBoolFull())
}

func TestFuzzBoolRunTests(t *testing.T) {
	FuzzBoolRunTests(t, FuzzBoolFull(), func(t *testing.T, index int, f bool) {
		AssertNotNil(t, f)
	})
}

func TestFuzzIntGenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("GenerateRandomPositive generates positive numbers only", func(t *testing.T) {
			AssertGreater(t, FuzzIntGenerateRandomPositive(1, 100)[0], 0)
		})

		t.Run("GenerateRandomNegative generates negative numbers only", func(t *testing.T) {
			AssertLess(t, FuzzIntGenerateRandomNegative(1, 100)[0], 0)
		})
	}
}

func TestFuzzIntFull(t *testing.T) {
	AssertGreater(t, len(FuzzIntFull()), 0)
}

func TestFuzzIntRunTests(t *testing.T) {
	FuzzIntRunTests(t, FuzzIntFull(), func(t *testing.T, index int, f int) {
		AssertNotNil(t, f)
	})
}

func TestFuzzIntModify(t *testing.T) {
	testSet := FuzzIntFull()
	s := FuzzUtilModifySet(testSet, func(index int, value int) int {
		return value * -1
	})

	for i, f := range testSet {
		t.Run("Number should be inverted", func(t *testing.T) {
			AssertEqual(t, f, s[i]*-1)
		})
	}
}

func TestFuzzFloat64Full(t *testing.T) {
	AssertGreater(t, len(FuzzFloat64Full()), 0)
}

func TestFuzzFloat64RunTests(t *testing.T) {
	FuzzFloat64RunTests(t, FuzzFloat64Full(), func(t *testing.T, index int, f float64) {
		AssertNotNil(t, f)
	})
}

func TestFuzzFloat64GenerateRandomNegative(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := FuzzFloat64GenerateRandomNegative(1, 0)[0]
		t.Run(fmt.Sprintf("%v must be negative", n), func(t *testing.T) {
			AssertLess(t, n, 0)
		})
	}
}

func TestFuzzFloat64GenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := FuzzFloat64GenerateRandomPositive(1, 0)[0]
		t.Run(fmt.Sprintf("%v must be positive", n), func(t *testing.T) {
			AssertGreater(t, n, 0)
		})
	}
}
