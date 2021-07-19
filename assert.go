package testza

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/MarvinJWendt/testza/internal"
	"github.com/pterm/pterm"
)

// ** Getter Methods **

func isKind(expectedKind reflect.Kind, value interface{}) bool {
	return reflect.TypeOf(value).Kind() == expectedKind
}

func isNumber(value interface{}) bool {
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
		if isKind(k, value) {
			return true
		}
	}

	return false
}

// completesIn returns if a function completes in a specific time.
func completesIn(duration time.Duration, f func()) bool {
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

func isZero(value interface{}) bool {
	return value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func isEqual(expected interface{}, actual interface{}) bool {
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

func hasEqualValues(expected interface{}, actual interface{}) bool {
	if isEqual(expected, actual) {
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

func doesImplement(interfaceObject, object interface{}) bool {
	interfaceType := reflect.TypeOf(interfaceObject).Elem()

	if object == nil {
		return false
	}
	if !reflect.TypeOf(object).Implements(interfaceType) {
		return false
	}

	return true
}

func doesContain(object, element interface{}) bool {
	objectValue := reflect.ValueOf(object)
	objectKind := reflect.TypeOf(object).Kind()

	switch objectKind {
	case reflect.String:
		return strings.Contains(objectValue.String(), reflect.ValueOf(element).String())
	case reflect.Map:
	default:
		for i := 0; i < objectValue.Len(); i++ {
			if isEqual(objectValue.Index(i).Interface(), element) {
				return true
			}
		}
	}

	return false
}

// ** Helper Methods **

// AssertKindOf asserts that the object is a type of kind exptectedKind.
func AssertKindOf(t testRunner, expectedKind reflect.Kind, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !isKind(expectedKind, object) {
		internal.Fail(t,
			fmt.Sprintf("A value that !!should be a type of kind %s!! is a type of kind %s.", expectedKind.String(), reflect.TypeOf(object).Kind().String()),
			internal.NewObjectsExpectedActual(expectedKind, object),
			msg...,
		)
	}
}

// AssertNotKindOf asserts that the object is not a type of kind `kind`.
func AssertNotKindOf(t testRunner, kind reflect.Kind, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if isKind(kind, object) {
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
func AssertNumeric(t testRunner, object interface{}, msg ...interface{}) {
	if !isNumber(object) {
		internal.Fail(t, "An object that !!should be a number!! is not of a numeric type.", internal.NewObjectsSingleNamed("object", object))
	}
}

// AssertNotNumeric checks if the object is not a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
func AssertNotNumeric(t testRunner, object interface{}, msg ...interface{}) {
	if isNumber(object) {
		internal.Fail(t, "An object that !!should not be a number!! is of a numeric type.", internal.NewObjectsSingleNamed("object", object))
	}
}

// AssertZero asserts that the value is the zero value for it's type.
//     testzZero(t, 0)
//     testzZero(t, false)
//     testzZero(t, "")
func AssertZero(t testRunner, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !isZero(value) {
		internal.Fail(t, "An object that !!should have it's zero value!!, does not have it's zero value.", internal.NewObjectsSingleNamed("object", value))
	}
}

// AssertNotZero asserts that the value is not the zero value for it's type.
//     testzNotZero(t, 1337)
//     testzNotZero(t, true)
//     testzNotZero(t, "Hello, World")
func AssertNotZero(t testRunner, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if isZero(value) {
		internal.Fail(t, "An object that !!should not have it's zero value!!, does have it's zero value.", internal.NewObjectsSingleNamed("object", value))
	}
}

// AssertEqual asserts that two objects are equal.
func AssertEqual(t testRunner, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !isEqual(expected, actual) {
		internal.Fail(t, "Two objects that !!should be equal!!, are not equal.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// AssertNotEqual asserts that two objects are not equal.
func AssertNotEqual(t testRunner, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if isEqual(expected, actual) {
		internal.Fail(t, "Two objects that !!should not be equal!!, are equal.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// AssertEqualValues asserts that two objects have equal values.
func AssertEqualValues(t testRunner, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !hasEqualValues(expected, actual) {
		internal.Fail(t, "Two objects that !!should have equal values!!, do not have equal values.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// AssertNotEqualValues asserts that two objects do not have equal values.
func AssertNotEqualValues(t testRunner, expected interface{}, actual interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if hasEqualValues(expected, actual) {
		internal.Fail(t, "Two objects that !!should not have equal values!!, do have equal values.", internal.NewObjectsExpectedActual(expected, actual), msg...)
	}
}

// AssertTrue asserts that an expression or object resolves to true.
func AssertTrue(t testRunner, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value != true {
		internal.Fail(t, "Value !!should be true!! but is not.", internal.NewObjectsExpectedActual(true, value))
	}
}

// AssertFalse asserts that an expression or object resolves to false.
func AssertFalse(t testRunner, value interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value == true {
		internal.Fail(t, "Value !!should be false!! but is not.", internal.NewObjectsExpectedActual(false, value))
	}
}

// AssertImplements checks if an objects implements an interface.
//
//	testza.AssertImplements(t, (*YourInterface)(nil), new(YourObject))
//	testza.AssertImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass
func AssertImplements(t testRunner, interfaceObject, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !doesImplement(interfaceObject, object) {
		internal.Fail(t, fmt.Sprintf("An object that !!should implement %s!! does not implement it.", reflect.TypeOf(interfaceObject).String()), internal.Objects{}, msg...)
	}
}

// AssertNotImplements checks if an object does not implement an interface.
//
//	testza.AssertNotImplements(t, (*YourInterface)(nil), new(YourObject))
//	testza.AssertNotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.
func AssertNotImplements(t testRunner, interfaceObject, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if doesImplement(interfaceObject, object) {
		internal.Fail(t, fmt.Sprintf("An object that !!should not implement %s!! does implement it.", reflect.TypeOf(interfaceObject).String()), internal.Objects{}, msg...)
	}
}

func AssertContains(t testRunner, object, element interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !doesContain(object, element) {
		internal.Fail(t, "An object !!does not contain!! the object it should contain.", internal.Objects{{Name: "object", Data: object}, {Name: "element that is missing in object", Data: element}})
	}
}

func AssertNotContains(t testRunner, object, element interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if doesContain(object, element) {
		internal.Fail(t, "An object !!does contain!! an object it should not contain.", internal.Objects{{Name: "object", Data: object}, {Name: "element that should not be in object", Data: element}})
	}
}

// AssertPanic asserts that a function panics.
func AssertPanic(t testRunner, f func(), msg ...interface{}) {
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

// AssertNotPanic asserts that a function does not panic.
func AssertNotPanic(t testRunner, f func(), msg ...interface{}) {
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
func AssertNil(t testRunner, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if object != nil {
		internal.Fail(t, "An object that !!should be nil!! is not nil.", internal.NewObjectsExpectedActual(nil, object))
	}
}

// AssertNotNil asserts that an object is not nil.
func AssertNotNil(t testRunner, object interface{}, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if object == nil {
		internal.Fail(t, "An object that !!should not be nil!! is nil.", internal.NewObjectsSingleNamed("object", object))
	}
}

// AssertCompletesIn asserts that a function completes in a given time.
// Use this function to test that functions do not take too long to complete.
//
// NOTE: Every system takes a different amount of time to complete a function.
// Do not set the duration too low, if you want consistent results.
func AssertCompletesIn(t testRunner, duration time.Duration, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !completesIn(duration, f) {
		internal.Fail(t, fmt.Sprintf("The function !!should complete in %s!!, but it did not.", duration), internal.Objects{}, msg...)
	}
}

// AssertNotCompletesIn asserts that a function does not complete in a given time.
// Use this function to test that functions do not complete to quickly.
// For example if your database connection completes in under a millisecond, there might be something wrong.
//
// NOTE: Every system takes a different amount of time to complete a function.
// Do not set the duration too high, if you want consistent results.
func AssertNotCompletesIn(t testRunner, duration time.Duration, f func(), msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if completesIn(duration, f) {
		internal.Fail(t, fmt.Sprintf("The function !!should not complete in %s!!, but it did.", duration), internal.Objects{}, msg...)
	}
}

// AssertNoError asserts that an error is nil.
func AssertNoError(t testRunner, err interface{}, msg ...interface{}) {
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

// AssertGreater asserts that the first object is greater than the second.
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
