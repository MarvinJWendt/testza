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

func (a AssertHelper) isKind(expectedKind reflect.Kind, value interface{}) bool {
	return reflect.TypeOf(value).Kind() == expectedKind
}

func (a AssertHelper) isNumber(value interface{}) bool {
	numberKinds := []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Float32,
		reflect.Float64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Complex64,
		reflect.Complex128,
	}

	for _, k := range numberKinds {
		if a.isKind(k, value) {
			return true
		}
	}

	return false
}

func (a AssertHelper) isGreater(base, value interface{}) {

}

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

// KindOf asserts that the object is a type of kind exptectedKind.
func (a AssertHelper) KindOf(t testingT, expectedKind reflect.Kind, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !a.isKind(expectedKind, object) {
		internal.Fail(t, generateMsg(msg,
			pterm.Sprintfln("An object that %s is a type of kind %s", highlight(pterm.Sprintf("should be a type of kind %s", expectedKind.String())), highlight(reflect.TypeOf(object).Kind())),
			spew.Sdump(object),
		))
	}
}

// NotKindOf asserts that the object is not a type of kind `kind`.
func (a AssertHelper) NotKindOf(t testingT, kind reflect.Kind, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if a.isKind(kind, value) {
		internal.Fail(t, generateMsg(msg,
			pterm.Sprintfln("A value that %s is a type of kind %s", highlight(pterm.Sprintf("should not be a type of kind %s", kind.String())), highlight(reflect.TypeOf(value).Kind())),
			spew.Sdump(value),
		))
	}
}

// Number asserts that the object is a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
func (a AssertHelper) Number(t testingT, object interface{}, msg ...interface{}) {
	if !a.isNumber(object) {
		internal.Fail(t, generateMsg(msg,
			pterm.Sprintfln("An object that %s is not a number", highlight("should be a number")),
			spew.Sdump(object),
		))
	}
}

// Number checks if the object is not a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
func (a AssertHelper) NotNumber(t testingT, object interface{}, msg ...interface{}) {
	if a.isNumber(object) {
		internal.Fail(t, generateMsg(msg,
			pterm.Sprintfln("An object that %s is a number", highlight("should not be a number")),
			spew.Sdump(object),
		))
	}
}

// Zero asserts that the value is the zero value for it's type.
//     assert.Zero(t, 0)
//     assert.Zero(t, false)
//     assert.Zero(t, "")
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

// NotZero asserts that the value is not the zero value for it's type.
//     assert.NotZero(t, 1337)
//     assert.NotZero(t, true)
//     assert.NotZero(t, "Hello, World")
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

// Equal asserts that two objects are equal.
func (a AssertHelper) Equal(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.isEqual(expected, actual) {
		a.failNotEqual(t, expected, actual, msg...)
	}
}

// NotEqual asserts that two objects are not equal.
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

// EqualValues asserts that two objects have equal values.
func (a AssertHelper) EqualValues(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.hasEqualValues(expected, actual) {
		a.failNotEqual(t, expected, actual, msg...)
	}
}

// NotEqualValues asserts that two objects do not have equal values.
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

// True asserts that an expression or object resolves to true.
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

// False asserts that an expression or object resolves to false.
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

// Panic asserts that a function panics.
func (a AssertHelper) Panic(t testingT, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	defer func() {
		if r := recover(); r == nil {
			output := generateMsg(msg,
				pterm.Sprintfln("The function %s, but does not panic", highlight("should panic")),
			)
			internal.Fail(t, output)
		}
	}()

	f()
}

// NotPanic asserts that a function does not panic.
func (a AssertHelper) NotPanic(t testingT, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	defer func() {
		if r := recover(); r != nil {
			output := generateMsg(msg,
				pterm.Sprintfln("The function %s, but it panics:\n%s", highlight("should not panic"), r),
			)
			internal.Fail(t, output)
		}
	}()

	f()
}

// Nil asserts that an object is nil.
func (a AssertHelper) Nil(t testingT, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if object != nil {
		output := generateMsg(msg,
			pterm.Sprintfln("An object %s is not nil:", highlight("should be nil")),
			spew.Sdump(object),
		)
		internal.Fail(t, output)
	}
}

// NotNil assertsthat an object is not nil.
func (a AssertHelper) NotNil(t testingT, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if object == nil {
		output := generateMsg(msg,
			pterm.Sprintfln("An object %s is nil:", highlight("that should not be nil")),
			spew.Sdump(object),
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
