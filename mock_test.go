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

	input := Use.Mock.Inputs.Strings.Modify(stringSlice, func(index int, value string) string {
		return strings.ToLower(value)
	})

	Use.Assert.Equal(t, expected, input)
}

func TestStringsHelper_Limit(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Limit=%d", i), func(t *testing.T) {
			Use.Assert.Equal(t, i, len(Use.Mock.Inputs.Strings.Limit(Use.Mock.Inputs.Strings.Full(), i)))
		})
	}
}

func TestStringsHelper_GenerateRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Length=%d", i), func(t *testing.T) {
			Use.Assert.Equal(t, i, len(Use.Mock.Inputs.Strings.GenerateRandom(1, i)[0]))
		})
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Count=%d", i), func(t *testing.T) {
			Use.Assert.Equal(t, i, len(Use.Mock.Inputs.Strings.GenerateRandom(i, 5)))
		})
	}
}

func TestStringsHelper_Numeric(t *testing.T) {
	for _, v := range Use.Mock.Inputs.Strings.Numeric() {
		t.Run(v, func(t *testing.T) {
			f, err := strconv.ParseFloat(v, 64)
			Use.Assert.Numeric(t, f)
			Use.Assert.NoError(t, err)
		})
	}
}

func TestStringsHelper_Full(t *testing.T) {
	Use.Assert.Greater(t, len(Use.Mock.Inputs.Strings.Full()), 0)
}

func TestBoolsHelper_Full(t *testing.T) {
	Use.Assert.Contains(t, Use.Mock.Inputs.Bools.Full(), true)
	Use.Assert.Contains(t, Use.Mock.Inputs.Bools.Full(), false)
}

func TestBoolsHelper_RunTests(t *testing.T) {
	Use.Mock.Inputs.Bools.RunTests(t, Use.Mock.Inputs.Bools.Full(), func(t *testing.T, index int, f bool) {
		Use.Assert.NotNil(t, f)
	})
}

func TestIntsHelper_GenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("GenerateRandomPositive generates positive numbers only", func(t *testing.T) {
			Use.Assert.Greater(t, Use.Mock.Inputs.Ints.GenerateRandomPositive(1, 100)[0], 0)
		})

		t.Run("GenerateRandomNegative generates negative numbers only", func(t *testing.T) {
			Use.Assert.Less(t, Use.Mock.Inputs.Ints.GenerateRandomNegative(1, 100)[0], 0)
		})
	}
}

func TestIntsHelper_Full(t *testing.T) {
	Use.Assert.Greater(t, len(Use.Mock.Inputs.Ints.Full()), 0)
}

func TestIntsHelper_RunTests(t *testing.T) {
	Use.Mock.Inputs.Ints.RunTests(t, Use.Mock.Inputs.Ints.Full(), func(t *testing.T, index int, f int) {
		Use.Assert.NotNil(t, f)
	})
}

func TestFloats64Helper_Full(t *testing.T) {
	Use.Assert.Greater(t, len(Use.Mock.Inputs.Floats64.Full()), 0)
}

func TestFloats64Helper_RunTests(t *testing.T) {
	Use.Mock.Inputs.Floats64.RunTests(t, Use.Mock.Inputs.Floats64.Full(), func(t *testing.T, index int, f float64) {
		Use.Assert.NotNil(t, f)
	})
}
