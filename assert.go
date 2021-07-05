package testza

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/atomicgo/testza/internal"
	"github.com/pterm/pterm"
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

// completesIn returns if a function completes in a specific time.
func (a AssertHelper) completesIn(duration time.Duration, f func()) bool {
	done := make(chan bool)
	go func() {
		f()
		done <- true
	}()

	select {
	case <-time.After(duration):
		return false
	case <-done:
		return true
	}
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
		internal.Fail(t,
			fmt.Sprintf("A value that !!should be a type of kind %s!! is a type of kind %s.", expectedKind.String(), reflect.TypeOf(object).Kind().String()),
			internal.NewObjectsExpectedActual(expectedKind, object),
			msg...,
		)
	}
}

// NotKindOf asserts that the object is not a type of kind `kind`.
func (a AssertHelper) NotKindOf(t testingT, kind reflect.Kind, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if a.isKind(kind, object) {
		internal.Fail(t,
			fmt.Sprintf("A value that !!should not be a type of kind %s!! is a type of kind %s.", kind.String(), reflect.TypeOf(object).Kind().String()),
			internal.NewObjectsExpectedActual(kind, object),
			msg...,
		)
	}
}

// Numeric asserts that the object is a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
func (a AssertHelper) Numeric(t testingT, object interface{}, msg ...interface{}) {
	if !a.isNumber(object) {
		internal.Fail(t, "An object that !!should be a number!! is not of a numeric type.", internal.NewObjectsSingleNamed("object", object))
	}
}

// Number checks if the object is not a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
func (a AssertHelper) NotNumeric(t testingT, object interface{}, msg ...interface{}) {
	if a.isNumber(object) {
		internal.Fail(t, "An object that !!should not be a number!! is of a numeric type.", internal.NewObjectsSingleNamed("object", object))
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
		internal.Fail(t, "An object that !!should have it's zero value!!, does not have it's zero value.", internal.NewObjectsSingleNamed("object", value))
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
		internal.Fail(t, "An object that !!should not have it's zero value!!, does have it's zero value.", internal.NewObjectsSingleNamed("object", value))
	}
}

// Equal asserts that two objects are equal.
func (a AssertHelper) Equal(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.isEqual(expected, actual) {
		internal.Fail(t, "Two objects that !!should be equal!!, are not equal.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// NotEqual asserts that two objects are not equal.
func (a AssertHelper) NotEqual(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if Use.Assert.isEqual(expected, actual) {
		internal.Fail(t, "Two objects that !!should not be equal!!, are equal.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// EqualValues asserts that two objects have equal values.
func (a AssertHelper) EqualValues(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.hasEqualValues(expected, actual) {
		internal.Fail(t, "Two objects that !!should have equal values!!, do not have equal values.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// NotEqualValues asserts that two objects do not have equal values.
func (a AssertHelper) NotEqualValues(t testingT, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if Use.Assert.hasEqualValues(expected, actual) {
		internal.Fail(t, "Two objects that !!should not have equal values!!, do have equal values.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// True asserts that an expression or object resolves to true.
func (a AssertHelper) True(t testingT, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value != true {
		internal.Fail(t, "Value !!should be true!! but is not.", internal.NewObjectsExpectedActual(true, value))
	}
}

// False asserts that an expression or object resolves to false.
func (a AssertHelper) False(t testingT, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value == true {
		internal.Fail(t, "Value !!should be false!! but is not.", internal.NewObjectsExpectedActual(false, value))
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
		internal.Fail(t, fmt.Sprintf("An object that !!should implement %s!! does not implement it.", reflect.TypeOf(interfaceObject).String()), internal.Objects{}, msg...)
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
		internal.Fail(t, fmt.Sprintf("An object that !!should not implement %s!! does implement it.", reflect.TypeOf(interfaceObject).String()), internal.Objects{}, msg...)
	}
}

func (a AssertHelper) Contains(t testingT, object, element interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !Use.Assert.doesContain(object, element) {
		internal.Fail(t, "An object !!does not contain!! the object it should contain.", internal.Objects{{Name: "object", Data: object}, {Name: "element that is missing in object", Data: element}})
	}
}

func (a AssertHelper) NotContains(t testingT, object, element interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if Use.Assert.doesContain(object, element) {
		internal.Fail(t, "An object !!does contain!! an object it should not contain.", internal.Objects{{Name: "object", Data: object}, {Name: "element that should not be in object", Data: element}})
	}
}

// Panic asserts that a function panics.
func (a AssertHelper) Panic(t testingT, f func(), msg ...interface{}) {
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

// NotPanic asserts that a function does not panic.
func (a AssertHelper) NotPanic(t testingT, f func(), msg ...interface{}) {
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

// Nil asserts that an object is nil.
func (a AssertHelper) Nil(t testingT, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if object != nil {
		internal.Fail(t, "An object that !!should be nil!! is not nil.", internal.NewObjectsExpectedActual(nil, object))
	}
}

// NotNil assertsthat an object is not nil.
func (a AssertHelper) NotNil(t testingT, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if object == nil {
		internal.Fail(t, "An object that !!should not be nil!! is nil.", internal.NewObjectsSingleNamed("object", object))
	}
}

// CompletesIn asserts that a function completes in a given time.
// Use this function to test that functions do not take too long to complete.
//
// NOTE: Every system takes a different amount of time to complete a function.
// Do not set the duration too low, if you want consistent results.
func (a AssertHelper) CompletesIn(t testingT, duration time.Duration, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !a.completesIn(duration, f) {
		internal.Fail(t, fmt.Sprintf("The function !!should complete in %s!!, but it did not.", duration), internal.Objects{}, msg...)
	}
}

// NotCompletesIn asserts that a function does not complete in a given time.
// Use this function to test that functions do not complete to quickly.
// For example if your database connection completes in under a millisecond, there might be something wrong.
//
// NOTE: Every system takes a different amount of time to complete a function.
// Do not set the duration too high, if you want consistent results.
func (a AssertHelper) NotCompletesIn(t testingT, duration time.Duration, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if a.completesIn(duration, f) {
		internal.Fail(t, fmt.Sprintf("The function !!should not complete in %s!!, but it did.", duration), internal.Objects{}, msg...)
	}
}

// NoError asserts that an error is nil.
func (a AssertHelper) NoError(t testingT, err interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if err != nil {
		internal.Fail(t, "An error that !!should be nil!! is not nil.", internal.Objects{{
			Name:      "Error",
			NameStyle: pterm.NewStyle(pterm.FgRed),
			Data:      err,
			DataStyle: pterm.NewStyle(pterm.FgRed),
		}}, msg...)
	}
}
