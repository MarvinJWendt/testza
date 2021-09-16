package internal

import (
	"bytes"
	"reflect"
	"strings"
	"time"
)

func IsKind(expectedKind reflect.Kind, value interface{}) bool {
	return reflect.TypeOf(value).Kind() == expectedKind
}

func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}

	switch reflect.ValueOf(object).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(object).IsNil()
	}

	return false
}

func IsNumber(value interface{}) bool {
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
		if IsKind(k, value) {
			return true
		}
	}

	return false
}

// CompletesIn returns if a function completes in a specific time.
func CompletesIn(duration time.Duration, f func()) bool {
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

func IsZero(value interface{}) bool {
	return value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func IsEqual(expected interface{}, actual interface{}) bool {
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

func HasEqualValues(expected interface{}, actual interface{}) bool {
	if IsEqual(expected, actual) {
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

func DoesImplement(interfaceObject, object interface{}) bool {
	interfaceType := reflect.TypeOf(interfaceObject).Elem()

	if object == nil {
		return false
	}
	if !reflect.TypeOf(object).Implements(interfaceType) {
		return false
	}

	return true
}

func DoesContain(object, element interface{}) bool {
	objectValue := reflect.ValueOf(object)
	objectKind := reflect.TypeOf(object).Kind()

	switch objectKind {
	case reflect.String:
		return strings.Contains(objectValue.String(), reflect.ValueOf(element).String())
	case reflect.Map:
	default:
		for i := 0; i < objectValue.Len(); i++ {
			if IsEqual(objectValue.Index(i).Interface(), element) {
				return true
			}
		}
	}

	return false
}
