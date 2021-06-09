package testutil

import (
	"fmt"
	s "strings"
	"testing"
)

func TestStrings_Modify(t *testing.T) {
	stringSlice := []string{"Hello", "World", "TeSt"}
	expected := []string{"hello", "world", "test"}

	input := Use.Mock.Strings.Modify(stringSlice, func(index int, value string) string {
		return s.ToLower(value)
	})

	Use.Assert.Equal(t, expected, input)
}

func TestStringsHelper_GenerateRandom(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Run(fmt.Sprintf("Length=%d", i), func(t *testing.T) {
			Use.Assert.Equal(t, i, len(Use.Mock.Strings.GenerateRandom(i, 1)[0]))
		})
	}

	for i := 0; i < 20; i++ {
		t.Run(fmt.Sprintf("Count=%d", i), func(t *testing.T) {
			Use.Assert.Equal(t, i, len(Use.Mock.Strings.GenerateRandom(5, i)))
		})
	}
}
