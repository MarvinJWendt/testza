package testza

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/pterm/pterm"

	"github.com/MarvinJWendt/testza/internal"
)

type testMock struct {
	ErrorCalled  bool
	ErrorMessage string
}

func (m *testMock) fail(msg ...interface{}) {
	m.ErrorCalled = true
	m.ErrorMessage = fmt.Sprint(msg...)
}

func (m *testMock) Error(args ...interface{}) {
	m.fail(args...)
}

// Errorf is a mock of testing.T.
func (m *testMock) Errorf(format string, args ...interface{}) {
	m.fail(fmt.Sprintf(format, args...))
}

// Fail is a mock of testing.T.
func (m *testMock) Fail() {
	m.fail()
}

// FailNow is a mock of testing.T.
func (m *testMock) FailNow() {
	m.fail()
}

// Fatal is a mock of testing.T.
func (m *testMock) Fatal(args ...interface{}) {
	m.fail(args...)
}

// Fatalf is a mock of testing.T.
func (m *testMock) Fatalf(format string, args ...interface{}) {
	m.fail(fmt.Sprintf(format, args...))
}

// ** Helper Methods **

// AssertKindOf asserts that the object is a type of kind exptectedKind.
//
// Example:
//  testza.AssertKindOf(t, reflect.Slice, []int{1,2,3})
//  testza.AssertKindOf(t, reflect.Slice, []string{"Hello", "World"})
//  testza.AssertKindOf(t, reflect.Int, 1337)
//  testza.AssertKindOf(t, reflect.Bool, true)
//  testza.AssertKindOf(t, reflect.Map, map[string]bool{})
func AssertKindOf(t testRunner, expectedKind reflect.Kind, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.IsKind(expectedKind, object) {
		internal.Fail(t,
			fmt.Sprintf("A value that !!should be a type of kind %s!! is a type of kind %s.", expectedKind.String(), reflect.TypeOf(object).Kind().String()),
			internal.NewObjectsExpectedActual(expectedKind, object),
			msg...,
		)
	}
}

// AssertNotKindOf asserts that the object is not a type of kind `kind`.
//
// Example:
//  testza.AssertNotKindOf(t, reflect.Slice, "Hello, World")
//  testza.AssertNotKindOf(t, reflect.Slice, true)
//  testza.AssertNotKindOf(t, reflect.Int, 13.37)
//  testza.AssertNotKindOf(t, reflect.Bool, map[string]bool{})
//  testza.AssertNotKindOf(t, reflect.Map, false)
func AssertNotKindOf(t testRunner, kind reflect.Kind, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.IsKind(kind, object) {
		internal.Fail(t,
			fmt.Sprintf("A value that !!should not be a type of kind %s!! is a type of kind %s.", kind.String(), reflect.TypeOf(object).Kind().String()),
			internal.NewObjectsExpectedActual(kind, object),
			msg...,
		)
	}
}

// AssertNumeric asserts that the object is a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
//
// Example:
//  testza.AssertNumeric(t, 123)
//  testza.AssertNumeric(t, 1.23)
//  testza.AssertNumeric(t, uint(123))
func AssertNumeric(t testRunner, object interface{}, msg ...interface{}) {
	if !internal.IsNumber(object) {
		internal.Fail(t, "An object that !!should be a number!! is not of a numeric type.", internal.NewObjectsSingleNamed("object", object))
	}
}

// AssertNotNumeric checks if the object is not a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
//
// Example:
//  testza.AssertNotNumeric(t, true)
//  testza.AssertNotNumeric(t, "123")
func AssertNotNumeric(t testRunner, object interface{}, msg ...interface{}) {
	if internal.IsNumber(object) {
		internal.Fail(t, "An object that !!should not be a number!! is of a numeric type.", internal.NewObjectsSingleNamed("object", object))
	}
}

// AssertZero asserts that the value is the zero value for it's type.
//
// Example:
//  testza.AssertZero(t, 0)
//  testza.AssertZero(t, false)
//  testza.AssertZero(t, "")
func AssertZero(t testRunner, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.IsZero(value) {
		internal.Fail(t, "An object that !!should have it's zero value!!, does not have it's zero value.", internal.NewObjectsSingleNamed("object", value))
	}
}

// AssertNotZero asserts that the value is not the zero value for it's type.
//
// Example:
//  testza.AssertNotZero(t, 1337)
//  testza.AssertNotZero(t, true)
//  testza.AssertNotZero(t, "Hello, World")
func AssertNotZero(t testRunner, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.IsZero(value) {
		internal.Fail(t, "An object that !!should not have it's zero value!!, does have it's zero value.", internal.NewObjectsSingleNamed("object", value))
	}
}

// AssertEqual asserts that two objects are equal.
//
// Example:
//  testza.AssertEqual(t, "Hello, World!", "Hello, World!")
//  testza.AssertEqual(t, true, true)
func AssertEqual(t testRunner, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.IsEqual(expected, actual) {
		internal.Fail(t, "Two objects that !!should be equal!!, are not equal.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// AssertNotEqual asserts that two objects are not equal.
//
// Example:
//  testza.AssertNotEqual(t, true, false)
//  testza.AssertNotEqual(t, "Hello", "World")
func AssertNotEqual(t testRunner, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.IsEqual(expected, actual) {
		internal.Fail(t, "Two objects that !!should not be equal!!, are equal.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// AssertEqualValues asserts that two objects have equal values.
// The order of the values is also validated.
//
// Example:
//  testza.AssertEqualValues(t, []string{"Hello", "World"}, []string{"Hello", "World"})
//  testza.AssertEqualValues(t, []int{1,2}, []int{1,2})
//  testza.AssertEqualValues(t, []int{1,2}, []int{2,1}) // FAILS (wrong order)
//
// Comparing struct values:
//  person1 := Person{
//    Name:   "Marvin Wendt",
//    Age:    20,
//    Gender: "male",
//  }
//
//  person2 := Person{
//    Name:   "Marvin Wendt",
//    Age:    20,
//    Gender: "male",
//  }
//
//  testza.AssertEqualValues(t, person1, person2)
func AssertEqualValues(t testRunner, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.HasEqualValues(expected, actual) {
		internal.Fail(t, "Two objects that !!should have equal values!!, do not have equal values.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// AssertNotEqualValues asserts that two objects do not have equal values.
//
// Example:
//  testza.AssertNotEqualValues(t, []int{1,2}, []int{3,4})
//
// Comparing struct values:
//  person1 := Person{
//    Name:   "Marvin Wendt",
//    Age:    20,
//    Gender: "male",
//  }
//
//  person2 := Person{
//    Name:   "Marvin Wendt",
//    Age:    20,
//    Gender: "female", // <-- CHANGED
//  }
//
//  testza.AssertNotEqualValues(t, person1, person2)
func AssertNotEqualValues(t testRunner, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.HasEqualValues(expected, actual) {
		internal.Fail(t, "Two objects that !!should not have equal values!!, do have equal values.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// AssertTrue asserts that an expression or object resolves to true.
//
// Example:
//  testza.AssertTrue(t, true)
//  testza.AssertTrue(t, 1 == 1)
//  testza.AssertTrue(t, 2 != 3)
//  testza.AssertTrue(t, 1 > 0 && 4 < 5)
func AssertTrue(t testRunner, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value != true {
		internal.Fail(t, "Value !!should be true!! but is not.", internal.NewObjectsExpectedActual(true, value))
	}
}

// AssertFalse asserts that an expression or object resolves to false.
//
// Example:
//  testza.AssertFalse(t, false)
//  testza.AssertFalse(t, 1 == 2)
//  testza.AssertFalse(t, 2 != 2)
//  testza.AssertFalse(t, 1 > 5 && 4 < 0)
func AssertFalse(t testRunner, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value == true {
		internal.Fail(t, "Value !!should be false!! but is not.", internal.NewObjectsExpectedActual(false, value))
	}
}

// AssertImplements asserts that an objects implements an interface.
//
// Example:
//  testza.AssertImplements(t, (*YourInterface)(nil), new(YourObject))
//  testza.AssertImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass
func AssertImplements(t testRunner, interfaceObject, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.DoesImplement(interfaceObject, object) {
		internal.Fail(t, fmt.Sprintf("An object that !!should implement %s!! does not implement it.", reflect.TypeOf(interfaceObject).String()), internal.Objects{}, msg...)
	}
}

// AssertNotImplements asserts that an object does not implement an interface.
//
// Example:
//  testza.AssertNotImplements(t, (*YourInterface)(nil), new(YourObject))
//  testza.AssertNotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.
func AssertNotImplements(t testRunner, interfaceObject, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.DoesImplement(interfaceObject, object) {
		internal.Fail(t, fmt.Sprintf("An object that !!should not implement %s!! does implement it.", reflect.TypeOf(interfaceObject).String()), internal.Objects{}, msg...)
	}
}

// AssertContains asserts that a string/list/array/slice/map contains the specified element.
//
// Example:
//  testza.AssertContains(t, []int{1,2,3}, 2)
//  testza.AssertContains(t, []string{"Hello", "World"}, "World")
//  testza.AssertContains(t, "Hello, World!", "World")
func AssertContains(t testRunner, object, element interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.DoesContain(object, element) {
		internal.Fail(t, "An object !!does not contain!! the object it should contain.", internal.Objects{{Name: "object", Data: object}, {Name: "element that is missing in object", Data: element}})
	}
}

// AssertNotContains asserts that a string/list/array/slice/map does not contain the specified element.
//
// Example:
//  testza.AssertNotContains(t, []string{"Hello", "World"}, "Spaceship")
//  testza.AssertNotContains(t, "Hello, World!", "Spaceship")
func AssertNotContains(t testRunner, object, element interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.DoesContain(object, element) {
		internal.Fail(t, "An object !!does contain!! an object it should not contain.", internal.Objects{{Name: "object", Data: object}, {Name: "element that should not be in object", Data: element}})
	}
}

// AssertPanics asserts that a function panics.
//
// Example:
//  testza.AssertPanics(t, func() {
//  	// ...
//  	panic("some panic")
//  }) // => PASS
func AssertPanics(t testRunner, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	defer func() {
		if r := recover(); r == nil {
			internal.Fail(t, "A function that !!should panic!! did not panic.", internal.Objects{}, msg...)
		}
	}()

	f()
}

// AssertNotPanics asserts that a function does not panic.
//
// Example:
//  testza.AssertNotPanics(t, func() {
//  	// some code that does not call a panic...
//  }) // => PASS
func AssertNotPanics(t testRunner, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	defer func() {
		if r := recover(); r != nil {
			internal.Fail(t, "A function that !!should not panic!! did panic.", internal.Objects{}, msg...)
		}
	}()

	f()
}

// AssertNil asserts that an object is nil.
//
// Example:
//  testza.AssertNil(t, nil)
func AssertNil(t testRunner, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.IsNil(object) {
		internal.Fail(t, "An object that !!should be nil!! is not nil.", internal.NewObjectsExpectedActual(nil, object))
	}
}

// AssertNotNil asserts that an object is not nil.
//
// Example:
//  testza.AssertNotNil(t, true)
//  testza.AssertNotNil(t, "Hello, World!")
//  testza.AssertNotNil(t, 0)
func AssertNotNil(t testRunner, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.IsNil(object) {
		internal.Fail(t, "An object that !!should not be nil!! is nil.", internal.NewObjectsSingleNamed("object", object))
	}
}

// AssertCompletesIn asserts that a function completes in a given time.
// Use this function to test that functions do not take too long to complete.
//
// NOTE: Every system takes a different amount of time to complete a function.
// Do not set the duration too low, if you want consistent results.
//
// Example:
//  testza.AssertCompletesIn(t, 2 * time.Second, func() {
//  	// some code that should take less than 2 seconds...
//  }) // => PASS
func AssertCompletesIn(t testRunner, duration time.Duration, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.CompletesIn(duration, f) {
		internal.Fail(t, fmt.Sprintf("The function !!should complete in %s!!, but it did not.", duration), internal.Objects{}, msg...)
	}
}

// AssertNotCompletesIn asserts that a function does not complete in a given time.
// Use this function to test that functions do not complete to quickly.
// For example if your database connection completes in under a millisecond, there might be something wrong.
//
// NOTE: Every system takes a different amount of time to complete a function.
// Do not set the duration too high, if you want consistent results.
//
// Example:
//  testza.AssertNotCompletesIn(t, 2 * time.Second, func() {
//  	// some code that should take more than 2 seconds...
//  	time.Sleep(3 * time.Second)
//  }) // => PASS
func AssertNotCompletesIn(t testRunner, duration time.Duration, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.CompletesIn(duration, f) {
		internal.Fail(t, fmt.Sprintf("The function !!should not complete in %s!!, but it did.", duration), internal.Objects{}, msg...)
	}
}

// AssertNoError asserts that an error is nil.
//
// Example:
//  err := nil
//  testza.AssertNoError(t, err)
func AssertNoError(t testRunner, err error, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if err != nil {
		internal.Fail(t, "An error that !!should be nil!! is not nil.", internal.Objects{
			{
				Name:      "Error message",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      fmt.Sprintf("%q\n", err.Error()),
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
			{
				Name:      "Error object",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      err,
				DataStyle: pterm.NewStyle(pterm.FgRed),
			}}, msg...)
	}
}

// AssertGreater asserts that the first object is greater than the second.
//
// Example:
//  testza.AssertGreater(t, 5, 1)
//  testza.AssertGreater(t, 10, -10)
func AssertGreater(t testRunner, object1, object2 interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	v1, err := strconv.ParseFloat(fmt.Sprint(object1), 64)
	v2, err2 := strconv.ParseFloat(fmt.Sprint(object2), 64)

	if err != nil || err2 != nil {
		internal.Fail(t, "An error occurred while parsing the objects as numbers.", internal.NewObjectsUnknown(object1, object2))
	}

	if !(v1 > v2) {
		internal.Fail(t, "An object that !!should be greater!! than the second object is not.", internal.Objects{{Name: "Object 1", Data: object1}, {Name: "Should be greater than object 2", Data: object2}}, msg...)
	}
}

// AssertLess asserts that the first object is less than the second.
//
// Example:
//  testza.AssertLess(t, 1, 5)
//  testza.AssertLess(t, -10, 10)
func AssertLess(t testRunner, object1, object2 interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	v1, err := strconv.ParseFloat(fmt.Sprint(object1), 64)
	v2, err2 := strconv.ParseFloat(fmt.Sprint(object2), 64)

	if err != nil || err2 != nil {
		internal.Fail(t, "An error occurred while parsing the objects as numbers.", internal.NewObjectsUnknown(object1, object2))
	}

	if !(v1 < v2) {
		internal.Fail(t, "An object that !!should be less!! than the second object is not.", internal.Objects{{Name: "Object 1", Data: object1}, {Name: "Should be less than object 2", Data: object2}}, msg...)
	}
}

// AssertTestFails asserts that a unit test fails.
// A unit test fails if one of the following methods is called in the test function: Error, Errorf, Fail, FailNow, Fatal, Fatalf
//
// Example:
//  testza.AssertTestFails(t, func(t testza.TestingPackageWithFailFunctions) {
//  	testza.AssertTrue(t, false)
//  }) // => Pass
//
//  testza.AssertTestFails(t, func(t testza.TestingPackageWithFailFunctions) {
//  	// ...
//  	t.Fail() // Or any other failing method.
//  }) // => Pass
func AssertTestFails(t testRunner, test func(t TestingPackageWithFailFunctions), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	var mock testMock
	test(&mock)

	if !mock.ErrorCalled {
		internal.Fail(t, "A test that !!should fail!! did not fail.", []internal.Object{}, msg...)
	}
}

// AssertErrorIs asserts that target is inside the error chain of err.
//
// Example:
//  var testErr = errors.New("hello world")
//  var testErrWrapped = fmt.Errorf("test err: %w", testErr)
//  testza.AssertErrorIs(t, testErrWrapped ,testErr)
func AssertErrorIs(t testRunner, err, target error, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !errors.Is(err, target) {
		internal.Fail(t, "Target error !!should be in the error chain!! of err.", internal.NewObjectsExpectedActual(target.Error(), err.Error()), msg...)
	}
}

// AssertNotErrorIs
//
// Example:
//  var testErr = errors.New("hello world")
//  var test2Err = errors.New("hello world 2")
//  var testErrWrapped = fmt.Errorf("test err: %w", testErr)
//  testza.AssertNotErrorIs(t, testErrWrapped, test2Err)
func AssertNotErrorIs(t testRunner, err, target error, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if errors.Is(err, target) {
		internal.Fail(t, "Target error !!should not be in the error chain!! of err.", internal.NewObjectsExpectedActual(target.Error(), err.Error()), msg...)
	}
}

// AssertLen asserts that the length of an object is equal to the given length.
//
// Example:
//  testza.AssertLen(t, "abc", 3)
//  testza.AssertLen(t, "Assert", 6)
//  testza.AssertLen(t, []int{1, 2, 1337, 25}, 4)
//  testza.AssertLen(t, map[string]int{"asd": 1, "test": 1337}, 2)
func AssertLen(t testRunner, object interface{}, length int, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	v := reflect.ValueOf(object)
	defer func() {
		if e := recover(); e != nil {
			internal.Fail(t, "The 'object' !!does not!! have a length.", internal.NewObjectsSingleUnknown(object), msg...)
		}
	}()

	if v.Len() != length {
		internal.Fail(t, "The length of 'object' !!is not!! the expected length.", internal.Objects{
			{
				Name:      "Expected length",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      fmt.Sprint(length) + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
			{
				Name:      "Actual length",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      fmt.Sprint(v.Len()) + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
			internal.NewObjectsSingleUnknown(object)[0],
		}, msg...)
	}
}

// AssertIncreasing asserts that the values in a slice are increasing.
// the test fails if the values are not in a slice or if the values are not comparable.
//
// Valid input kinds are: []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float32, []float64.
//
// Example:
//  testza.AssertIncreasing(t, []int{1, 2, 137, 1000})
//  testza.AssertIncreasing(t, []float32{-10.3, 0.1, 7, 13.5})
func AssertIncreasing(t testRunner, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.AssertCompareHelper(t, object, 1, msg...)
}

// AssertDecreasing asserts that the values in a slice are decreasing.
// the test fails if the values are not in a slice or if the values are not comparable.
//
// Valid input kinds are: []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float32, []float64.
//
// Example:
//  testza.AssertDecreasing(t, []int{1000, 137, 2, 1})
//  testza.AssertDecreasing(t, []float32{13.5, 7, 0.1, -10.3})
func AssertDecreasing(t testRunner, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.AssertCompareHelper(t, object, -1, msg...)
}

// AssertRegexp asserts that a string matches a given regexp.
//
// Example:
//  testza.AssertRegexp(t, "^a.*c$", "abc")
func AssertRegexp(t testRunner, regex interface{}, txt interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.AssertRegexpHelper(t, regex, txt, true, msg...)
}

// AssertNotRegexp asserts that a string does not match a given regexp.
//
// Example:
//  testza.AssertNotRegexp(t, "ab.*", "Hello, World!")
func AssertNotRegexp(t testRunner, regex interface{}, txt interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.AssertRegexpHelper(t, regex, txt, false, msg...)
}

// AssertDirEmpty asserts that a directory is empty
//
// Example:
//  testza.AssertDirEmpty(t, "FolderName")
func AssertDirEmpty(t testRunner, dir string, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.AssertDirEmptyHelper(t, dir) {
		internal.Fail(t, "The directory !!is not!! empty.", internal.NewObjectsSingleNamed("Directory", dir))
	}
}

// AssertDirNotEmpty asserts that a directory is not empty
//
// Example:
//  testza.AssertDirNotEmpty(t, "FolderName")
func AssertDirNotEmpty(t testRunner, dir string, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.AssertDirEmptyHelper(t, dir) {
		internal.Fail(t, "The directory !!is!! empty.", internal.NewObjectsSingleNamed("Directory", dir))
	}
}
