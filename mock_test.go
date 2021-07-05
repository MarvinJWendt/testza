package testza

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStrings_Modify(t *testing.T) {
	stringSlice := []string{"Hello", "World", "TeSt"}
	expected := []string{"hello", "world", "test"}

	input := Use.Mock.Strings.Modify(stringSlice, func(index int, value string) string {
		return strings.ToLower(value)
	})

	Use.Assert.Equal(t, expected, input)
}

func TestStringsHelper_GenerateRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Length=%d", i), func(t *testing.T) {
			Use.Assert.Equal(t, i, len(Use.Mock.Strings.GenerateRandom(1, i)[0]))
		})
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Count=%d", i), func(t *testing.T) {
			Use.Assert.Equal(t, i, len(Use.Mock.Strings.GenerateRandom(i, 5)))
		})
	}
}

func TestStringsHelper_Numeric(t *testing.T) {
	for _, v := range Use.Mock.Strings.Numeric() {
		t.Run(v, func(t *testing.T) {
			f, err := strconv.ParseFloat(v, 64)
			Use.Assert.Numeric(t, f)
			Use.Assert.NoError(t, err)
		})
	}
}

func TestBoolsHelper_Full(t *testing.T) {
	Use.Assert.Contains(t, Use.Mock.Bools.Full(), true)
	Use.Assert.Contains(t, Use.Mock.Bools.Full(), false)
}

// func TestIntsHelper_GenerateRandomPositive(t *testing.T) {
// 	for i := 0; i < 100; i++ {
// 		t.Run(fmt.Sprintf("Max=%d", i), func(t *testing.T) {
// 			Use.Assert.Equal(t, i, len(Use.Mock.Floats64.GenerateRandom(1, i)[0]))
// 		})
// 	}
// }
