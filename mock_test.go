package testza

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringsHelper_Modify(t *testing.T) {
	stringSlice := []string{"Hello", "World", "TeSt"}
	expected := []string{"hello", "world", "test"}

	input := MockInputStringModify(stringSlice, func(index int, value string) string {
		return strings.ToLower(value)
	})

	AssertEqual(t, expected, input)
}

func TestStringsHelper_Limit(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Limit=%d", i), func(t *testing.T) {
			AssertEqual(t, i, len(MockInputStringLimit(MockInputStringFull(), i)))
		})
	}
}

func TestStringsHelper_GenerateRandom(t *testing.T) {
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

func TestStringsHelper_Numeric(t *testing.T) {
	for _, v := range MockInputStringNumeric() {
		t.Run(v, func(t *testing.T) {
			f, err := strconv.ParseFloat(v, 64)
			AssertNumeric(t, f)
			AssertNoError(t, err)
		})
	}
}

func TestStringsHelper_Full(t *testing.T) {
	AssertGreater(t, len(MockInputStringFull()), 0)
}

func TestBoolsHelper_Full(t *testing.T) {
	AssertContains(t, MockInputBoolFull(), true)
	AssertContains(t, MockInputBoolFull(), false)
}

func TestBoolsHelper_RunTests(t *testing.T) {
	MockInputBoolRunTests(t, MockInputBoolFull(), func(t *testing.T, index int, f bool) {
		AssertNotNil(t, f)
	})
}

func TestIntsHelper_GenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("GenerateRandomPositive generates positive numbers only", func(t *testing.T) {
			AssertGreater(t, MockInputIntGenerateRandomPositive(1, 100)[0], 0)
		})

		t.Run("GenerateRandomNegative generates negative numbers only", func(t *testing.T) {
			AssertLess(t, MockInputIntGenerateRandomNegative(1, 100)[0], 0)
		})
	}
}

func TestIntsHelper_Full(t *testing.T) {
	AssertGreater(t, len(MockInputIntFull()), 0)
}

func TestIntsHelper_RunTests(t *testing.T) {
	MockInputIntRunTests(t, MockInputIntFull(), func(t *testing.T, index int, f int) {
		AssertNotNil(t, f)
	})
}

func TestFloats64Helper_Full(t *testing.T) {
	AssertGreater(t, len(MockInputFloat64Full()), 0)
}

func TestFloats64Helper_RunTests(t *testing.T) {
	MockInputFloat64RunTests(t, MockInputFloat64Full(), func(t *testing.T, index int, f float64) {
		AssertNotNil(t, f)
	})
}
