package testza_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/MarvinJWendt/testza"
)

func TestMockInputStringModify(t *testing.T) {
	stringSlice := []string{"Hello", "World", "TeSt"}
	expected := []string{"hello", "world", "test"}

	input := MockInputStringModify(stringSlice, func(index int, value string) string {
		return strings.ToLower(value)
	})

	AssertEqual(t, expected, input)
}

func TestMockInputStringLimit(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Limit=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(MockInputStringLimit(MockInputStringFull(), i)))
		})
	}
}

func TestMockInputStringGenerateRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Length=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(MockInputStringGenerateRandom(1, i)[0]))
		})
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Count=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(MockInputStringGenerateRandom(i, 5)))
		})
	}
}

func TestMockInputStringNumeric(t *testing.T) {
	for _, v := range MockInputStringNumeric() {
		t.Run(v, func(t *testing.T) {
			f, err := strconv.ParseFloat(v, 64)
			AssertNumeric(t, f)
			AssertNoError(t, err)
		})
	}
}

func TestMockInputStringFull(t *testing.T) {
	AssertGreater(t, len(MockInputStringFull()), 0)
}

func TestMockInputBoolFull(t *testing.T) {
	AssertEqual(t, []bool{true, false}, MockInputBoolFull())
}

func TestMockInputBoolRunTests(t *testing.T) {
	MockInputBoolRunTests(t, MockInputBoolFull(), func(t *testing.T, index int, f bool) {
		AssertNotNil(t, f)
	})
}

func TestMockInputIntGenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("GenerateRandomPositive generates positive numbers only", func(t *testing.T) {
			AssertGreater(t, MockInputIntGenerateRandomPositive(1, 100)[0], 0)
		})

		t.Run("GenerateRandomNegative generates negative numbers only", func(t *testing.T) {
			AssertLess(t, MockInputIntGenerateRandomNegative(1, 100)[0], 0)
		})
	}
}

func TestMockInputIntFull(t *testing.T) {
	AssertGreater(t, len(MockInputIntFull()), 0)
}

func TestMockInputIntRunTests(t *testing.T) {
	MockInputIntRunTests(t, MockInputIntFull(), func(t *testing.T, index int, f int) {
		AssertNotNil(t, f)
	})
}

func TestMockInputIntModify(t *testing.T) {
	testSet := MockInputIntFull()
	s := MockInputIntModify(testSet, func(index int, value int) int {
		return value * -1
	})

	for i, f := range testSet {
		t.Run("Number should be inverted", func(t *testing.T) {
			AssertEqual(t, f, s[i]*-1)
		})
	}
}

func TestMockInputFloat64Full(t *testing.T) {
	AssertGreater(t, len(MockInputFloat64Full()), 0)
}

func TestMockInputFloat64RunTests(t *testing.T) {
	MockInputFloat64RunTests(t, MockInputFloat64Full(), func(t *testing.T, index int, f float64) {
		AssertNotNil(t, f)
	})
}

func TestMockInputFloat64GenerateRandomNegative(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := MockInputFloat64GenerateRandomNegative(1, 0)[0]
		t.Run(fmt.Sprintf("%v must be negative", n), func(t *testing.T) {
			AssertLess(t, n, 0)
		})
	}
}

func TestMockInputFloat64GenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := MockInputFloat64GenerateRandomPositive(1, 0)[0]
		t.Run(fmt.Sprintf("%v must be positive", n), func(t *testing.T) {
			AssertGreater(t, n, 0)
		})
	}
}

func TestMockInputFloat64Modify(t *testing.T) {
	testSet := MockInputFloat64Full()
	s := MockInputFloat64Modify(testSet, func(index int, value float64) float64 {
		return value * -1
	})

	for i, f := range testSet {
		t.Run("Number should be inverted", func(t *testing.T) {
			AssertEqual(t, f, s[i]*-1)
		})
	}
}

func TestMockInputBoolModify(t *testing.T) {
	s := MockInputBoolModify(MockInputBoolFull(), func(index int, value bool) bool {
		return !value
	})
	AssertEqual(t, []bool{false, true}, s)
}
