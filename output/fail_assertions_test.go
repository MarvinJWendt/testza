package output

import (
	"errors"
	"fmt"
	"github.com/MarvinJWendt/testza"
	"go/types"
	"reflect"
	"testing"
	"time"
)

// This file contains all assertions and lets them fail.
// By running this test suite, you can see the output of every failed assertion.

func TestOutput(t *testing.T) {
	t.Run("AssertCompletesIn", func(t *testing.T) {
		testza.AssertCompletesIn(t, time.Millisecond, func() {
			time.Sleep(time.Millisecond * 10)
		})
	})

	t.Run("AssertContains", func(t *testing.T) {
		testza.AssertContains(t, "Hello World", "Earth")
	})

	t.Run("AssertDecreasing", func(t *testing.T) {
		testza.AssertDecreasing(t, []int{1, 2, 3, 4, 5})
	})

	t.Run("AssertDirEmpty", func(t *testing.T) {
		testza.AssertDirEmpty(t, ".")
	})

	t.Run("AssertDirExist", func(t *testing.T) {
		testza.AssertDirExist(t, ".")
	})

	t.Run("AssertDirNotEmpty", func(t *testing.T) {
		testza.AssertDirNotEmpty(t, "./empty-dir")
	})

	t.Run("AssertEqual", func(t *testing.T) {
		testza.AssertEqual(t, "Hello, World!", "Hello World")
	})

	t.Run("AssertEqualValues", func(t *testing.T) {
		testza.AssertEqualValues(t, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 9})
	})

	t.Run("AssertErrorIs", func(t *testing.T) {
		err1 := errors.New("error 1")
		err2 := errors.New("error 2")
		testza.AssertErrorIs(t, err2, err1)
	})

	t.Run("AssertFalse", func(t *testing.T) {
		testza.AssertFalse(t, true)
	})

	t.Run("AssertFileExists", func(t *testing.T) {
		testza.AssertFileExists(t, "./non-existent-file")
	})

	t.Run("AssertGreater", func(t *testing.T) {
		testza.AssertGreater(t, 1, 2)
	})

	t.Run("AssertImplements", func(t *testing.T) {
		testza.AssertImplements(t, (*fmt.Scanner)(nil), new(types.Const))
	})

	t.Run("AssertIncreasing", func(t *testing.T) {
		testza.AssertIncreasing(t, []int{5, 4, 3, 2, 1})
	})

	t.Run("AssertKindOf", func(t *testing.T) {
		testza.AssertKindOf(t, reflect.Int, false)
	})

	t.Run("AssertLen", func(t *testing.T) {
		testza.AssertLen(t, []int{1, 2, 3, 4, 5}, 100)
	})

	t.Run("AssertLess", func(t *testing.T) {
		testza.AssertLess(t, 100, 10)
	})

	t.Run("AssertNil", func(t *testing.T) {
		testza.AssertNil(t, true)
	})

	t.Run("AssertNoDirExists", func(t *testing.T) {
		testza.AssertNoDirExists(t, ".")
	})

	t.Run("AssertNoError", func(t *testing.T) {
		testza.AssertNoError(t, errors.New("err"))
	})

	t.Run("AssertNoFileExists", func(t *testing.T) {
		testza.AssertNoFileExists(t, "./fail_assertions_test.go")
	})

	t.Run("AssertNoSubset", func(t *testing.T) {
		testza.AssertNoSubset(t, []string{"Hello", "World", "Test"}, []string{"Test"})
	})

	t.Run("AssertNotCompletesIn", func(t *testing.T) {
		testza.AssertNotCompletesIn(t, time.Second, func() {})
	})

	t.Run("AssertNotContains", func(t *testing.T) {
		testza.AssertNotContains(t, "Hello World", "World")
	})

	t.Run("AssertNotEqual", func(t *testing.T) {
		testza.AssertNotEqual(t, "Hello, World!", "Hello, World!")
	})

	t.Run("AssertNotEqualValues", func(t *testing.T) {
		testza.AssertNotEqualValues(t, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})
	})

	t.Run("AssertNotErrorIs", func(t *testing.T) {
		err1 := errors.New("error 1")
		err2 := fmt.Errorf("error 2: %w", err1)
		testza.AssertNotErrorIs(t, err2, err1)
	})

	t.Run("AssertNotImplements", func(t *testing.T) {
		testza.AssertNotImplements(t, (*fmt.Stringer)(nil), new(types.Const))
	})

	t.Run("AssertNotKindOf", func(t *testing.T) {
		testza.AssertNotKindOf(t, reflect.Int, 1337)
	})

	t.Run("AssertNotNil", func(t *testing.T) {
		testza.AssertNotNil(t, nil)
	})

	t.Run("AssertNotNumeric", func(t *testing.T) {
		testza.AssertNotNumeric(t, 1337)
	})

	t.Run("AssertNotPanics", func(t *testing.T) {
		testza.AssertNotPanics(t, func() {
			panic("panic")
		})
	})

	t.Run("AssertNotRegexp", func(t *testing.T) {
		testza.AssertNotRegexp(t, "He.*", "Hello, World!")
	})

	t.Run("AssertNotSameElements", func(t *testing.T) {
		testza.AssertNotSameElements(t, []int{1, 2}, []int{1, 2})
	})

	t.Run("AssertNotZero", func(t *testing.T) {
		testza.AssertNotZero(t, 0)
	})

	t.Run("AssertNumeric", func(t *testing.T) {
		testza.AssertNumeric(t, false)
	})

	t.Run("AssertPanics", func(t *testing.T) {
		testza.AssertPanics(t, func() {})
	})

	t.Run("AssertRegexp", func(t *testing.T) {
		testza.AssertRegexp(t, "asd.*", "Hello, World!")
	})

	t.Run("AssertSameElements", func(t *testing.T) {
		testza.AssertSameElements(t, []int{1, 2, 3}, []int{4, 5, 6})
	})

	t.Run("AssertSubset", func(t *testing.T) {
		testza.AssertSubset(t, []int{1, 2, 3}, []int{4})
	})

	t.Run("AssertTestFails", func(t *testing.T) {
		testza.AssertTestFails(t, func(t testza.TestingPackageWithFailFunctions) {
			testza.AssertTrue(t, true)
		})
	})

	t.Run("AssertTrue", func(t *testing.T) {
		testza.AssertTrue(t, false)
	})

	t.Run("AssertZero", func(t *testing.T) {
		testza.AssertZero(t, "asd")
	})
}
