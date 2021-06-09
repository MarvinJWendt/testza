package testutil

import (
	"bytes"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"

	"github.com/atomicgo/testutil/internal"
)

var Assert AssertHelper

type AssertHelper struct{}

func (a AssertHelper) isEqual(expected interface{}, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	expectedB, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	actualB, ok := actual.([]byte)
	if !ok {
		return false
	}
	if expectedB == nil || actualB == nil {
		return expectedB == nil && actualB == nil
	}

	return bytes.Equal(expectedB, actualB)
}

func (a AssertHelper) hasEqualValues(expected interface{}, actual interface{}) bool {
	if a.isEqual(expected, actual) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}

	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}

	return false
}

func (a AssertHelper) doesImplement(interfaceObject, object interface{}) bool {
	interfaceType := reflect.TypeOf(interfaceObject).Elem()

	if object == nil {
		return false
	}
	if !reflect.TypeOf(object).Implements(interfaceType) {
		return false
	}

	return true
}

// Equal checks if two objects are equal.
func (a AssertHelper) Equal(t TestingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !a.isEqual(expected, actual) {
		a.failNotEqual(t, expected, actual, msg...)
	}
}

// NotEqual checks if two objects are not equal.
func (a AssertHelper) NotEqual(t TestingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if a.isEqual(expected, actual) {
		output := generateMsg(msg,
			pterm.Sprintfln("Two values that %s are equal:", highlight("should not be equal")),
			pterm.Sprintfln("Expected and actual both have the value(s):"),
			spew.Sdump(expected),
		)
		internal.Fail(t, output)
	}
}

// EqualValues checks if two objects have equal values.
func (a AssertHelper) EqualValues(t TestingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !a.hasEqualValues(expected, actual) {
		a.failNotEqual(t, expected, actual, msg...)
	}
}

// NotEqualValues checks if two objects do not have equal values.
func (a AssertHelper) NotEqualValues(t TestingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if a.hasEqualValues(expected, actual) {
		output := generateMsg(msg,
			pterm.Sprintfln("Two values that %s are equal:", highlight("should not be equal")),
			pterm.Sprintfln("Expected and actual both have the value(s):"),
			spew.Sdump(expected),
		)
		internal.Fail(t, output)
	}
}

// True checks if an expression or object resolves to true.
func (a AssertHelper) True(t TestingT, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value != true {
		output := generateMsg(
			msg,
			pterm.Sprintfln("Value that %s is not true:", highlight("should be true")),
			pterm.Sprintfln("Expected value: %s", green("true")),
			pterm.Sprintfln("Actual value: %s", red(pterm.Sprintf("%#v", value))),
		)
		internal.Fail(t, output)
	}
}

// False checks if an expression or object resolves to false.
func (a AssertHelper) False(t TestingT, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value == true {
		output := generateMsg(
			msg,
			pterm.Sprintfln("Value that %s is not true:", highlight("should be false")),
			pterm.Sprintfln("Expected value: %s", green("false")),
			pterm.Sprintfln("Actual value: %s", red(pterm.Sprintf("%#v", value))),
		)
		internal.Fail(t, output)
	}
}

// Implements checks if an objects implements an interface.
//
//	testutil.Assert.Implements(t, (*YourInterface)(nil), new(YourObject))
//	testutil.Assert.Implements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass
func (a AssertHelper) Implements(t TestingT, interfaceObject, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !a.doesImplement(interfaceObject, object) {
		output := generateMsg(msg, pterm.Sprintfln("The object %s %s %v:\nObject:\n%s", pterm.Magenta(reflect.TypeOf(object)),
			highlight("should implement"),
			pterm.Magenta(reflect.TypeOf(interfaceObject)),
			spew.Sdump(object)),
		)
		internal.Fail(t, output)
	}
}

// NotImplements checks if an object does not implement an interface.
//
//	testutil.Assert.NotImplements(t, (*YourInterface)(nil), new(YourObject))
//	testutil.Assert.NotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.
func (a AssertHelper) NotImplements(t TestingT, interfaceObject, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if a.doesImplement(interfaceObject, object) {
		output := generateMsg(msg, pterm.Sprintfln("The object %s %s %v:\nObject:\n%s",
			pterm.Magenta(reflect.TypeOf(object)), highlight("should not implement"),
			pterm.Magenta(reflect.TypeOf(interfaceObject)), spew.Sdump(object)),
		)

		internal.Fail(t, output)
	}
}

func (a AssertHelper) failNotEqual(t TestingT, expected interface{}, actual interface{}, msg ...interface{}) {
	diff := internal.GetDifference(expected, actual)
	output := generateMsg(
		msg,
		pterm.Sprintfln("Two values that %s are not equal:", highlight("should be equal")),
		pterm.Sprint(diff),
	)
	internal.Fail(t, output)
}
