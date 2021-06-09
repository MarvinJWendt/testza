package testutil

import (
	s "strings"
	"testing"
)

func TestStrings_Modify(t *testing.T) {
	stringSlice := []string{"Hello", "World", "TeSt"}
	expected := []string{"hello", "world", "test"}

	input := Input.Strings.Modify(stringSlice, func(index int, value string) string {
		return s.ToLower(value)
	})

	Assert.Equal(t, expected, input)
}
