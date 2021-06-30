package testza

import (
	"bytes"
	"reflect"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"

	"github.com/atomicgo/testza/internal"
)

// ** Getter Methods **

func (a AssertHelper) isZero(value interface{}) bool {
	return value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

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

func (a AssertHelper) doesContain(object, element interface{}) bool {
	objectValue := reflect.ValueOf(object)
	objectKind := reflect.TypeOf(object).Kind()

	switch objectKind {
	case reflect.String:
		return strings.Contains(objectValue.String(), reflect.ValueOf(element).String())
	case reflect.Map:
	default:
		for i := 0; i < objectValue.Len(); i++ {
			if Use.Assert.isEqual(objectValue.Index(i).Interface(), element) {
				return true
			}
		}
	}

	return false
}

// ** Helper Methods **

func (a AssertHelper) Zero(t testingT, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !a.isZero(value) {
		output := generateMsg(msg,
			pterm.Sprintfln("A value that %s is not zero.", highlight("should be zero")),
			pterm.Sprintfln("Actual value:"),
			spew.Sdump(value),
		)
		internal.Fail(t, output)
	}
}

func (a AssertHelper) NotZero(t testingT, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if a.isZero(value) {
		output := generateMsg(msg,
			pterm.Sprintfln("A value that %s is zero.", highlight("should not be zero")),
			pterm.Sprintfln("Actual value:"),
			spew.Sdump(value),
		)
		internal.Fail(t, output)
	}
}

// Equal checks if two objects are equal.
func (a AssertHelper) Equal(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.isEqual(expected, actual) {
		a.failNotEqual(t, expected, actual, msg...)
	}
}

// NotEqual checks if two objects are not equal.
func (a AssertHelper) NotEqual(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if Use.Assert.isEqual(expected, actual) {
		output := generateMsg(msg,
			pterm.Sprintfln("Two values that %s are equal:", highlight("should not be equal")),
			pterm.Sprintfln("Expected and actual both have the value(s):"),
			spew.Sdump(expected),
		)
		internal.Fail(t, output)
	}
}

// EqualValues checks if two objects have equal values.
func (a AssertHelper) EqualValues(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.hasEqualValues(expected, actual) {
		a.failNotEqual(t, expected, actual, msg...)
	}
}

// NotEqualValues checks if two objects do not have equal values.
func (a AssertHelper) NotEqualValues(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if Use.Assert.hasEqualValues(expected, actual) {
		output := generateMsg(msg,
			pterm.Sprintfln("Two values that %s are equal:", highlight("should not be equal")),
			pterm.Sprintfln("Expected and actual both have the value(s):"),
			spew.Sdump(expected),
		)
		internal.Fail(t, output)
	}
}

// True checks if an expression or object resolves to true.
func (a AssertHelper) True(t testingT, value interface{}, msg ...interface{}) {
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
func (a AssertHelper) False(t testingT, value interface{}, msg ...interface{}) {
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
//	testza.Use.Assert.Implements(t, (*YourInterface)(nil), new(YourObject))
//	testza.Use.Assert.Implements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass
func (a AssertHelper) Implements(t testingT, interfaceObject, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.doesImplement(interfaceObject, object) {
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
//	testza.Use.Assert.NotImplements(t, (*YourInterface)(nil), new(YourObject))
//	testza.Use.Assert.NotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.
func (a AssertHelper) NotImplements(t testingT, interfaceObject, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if Use.Assert.doesImplement(interfaceObject, object) {
		output := generateMsg(msg, pterm.Sprintfln("The object %s %s %v:\nObject:\n%s",
			pterm.Magenta(reflect.TypeOf(object)), highlight("should not implement"),
			pterm.Magenta(reflect.TypeOf(interfaceObject)), spew.Sdump(object)),
		)

		internal.Fail(t, output)
	}
}

func (a AssertHelper) Contains(t testingT, object, element interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.doesContain(object, element) {
		output := generateMsg(msg,
			pterm.Sprintfln("Object %s:\n", highlight("should contain")),
			spew.Sdump(element),
		)
		internal.Fail(t, output)
	}
}

func (a AssertHelper) NotContains(t testingT, object, element interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if Use.Assert.doesContain(object, element) {
		output := generateMsg(msg,
			pterm.Sprintfln("Object %s:\n", highlight("should not contain")),
			spew.Sdump(element),
		)
		internal.Fail(t, output)
	}
}

// ** Helper Methods **

func (a AssertHelper) failNotEqual(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	diff := internal.GetDifference(expected, actual)
	output := generateMsg(
		msg,
		pterm.Sprintfln("Two values that %s are not equal:", highlight("should be equal")),
		pterm.Sprint(diff),
	)
	internal.Fail(t, output)
}
