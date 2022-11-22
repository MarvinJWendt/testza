package internal

import (
	"atomicgo.dev/assert"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

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

// IsEqual checks if two objects are equal.
func IsEqual(expected any, actual any) bool {
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

// HasEqualValues checks if two objects have equal values.
func HasEqualValues(expected any, actual any) bool {
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

// DoesContain checks that ab objects contains an element.
func DoesContain(object, element any) bool {
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

// AssertCompareHelper option: 1 = increasing, 0 = equal, -1 = decreasing
func AssertCompareHelper(t testRunner, object any, option int, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	defer func() {
		if e := recover(); e != nil {
			Fail(t, "The 'object' !!must be a numeric slice!!.", NewObjectsSingleUnknown(object), msg...)
		}
	}()

	v := reflect.ValueOf(object)

	objKind := v.Kind()
	if objKind != reflect.Slice && objKind != reflect.Array {
		Fail(t, "The 'object' !!is neither a slice nor an array!!.", NewObjectsSingleUnknown(object), msg...)
		return
	}

	if v.Len() < 2 {
		Fail(t, "The 'object' !!is not long enough!!.", NewObjectsSingleUnknown(object), msg...)
		return
	}

	firstValue := v.Index(0).Interface()

	var ok bool

	switch firstValue.(type) {
	case int, int8, int16, int32, int64:
		ok = CompareTwoValuesInASlice(v, func(a, b reflect.Value) bool {
			if option == 1 {
				return a.Int() < b.Int()
			} else if option == 0 {
				return a.Int() == b.Int()
			} else if option == -1 {
				return a.Int() > b.Int()
			}
			return false
		})
	case uint, uint8, uint16, uint32, uint64:
		ok = CompareTwoValuesInASlice(v, func(a, b reflect.Value) bool {
			if option == 1 {
				return a.Uint() < b.Uint()
			} else if option == 0 {
				return a.Uint() == b.Uint()
			} else if option == -1 {
				return a.Uint() > b.Uint()
			}
			return false
		})
	case float32, float64:
		ok = CompareTwoValuesInASlice(v, func(a, b reflect.Value) bool {
			if option == 1 {
				return a.Float() < b.Float()
			} else if option == 0 {
				return a.Float() == b.Float()
			} else if option == -1 {
				return a.Float() > b.Float()
			}
			return false
		})
	default:
		Fail(t, "The 'object' !!must be a numeric slice!!.", NewObjectsSingleUnknown(object), msg...)
	}

	if !ok {
		var order string
		switch option {
		case 1:
			order = "increasing"
		case 0:
			order = "equal"
		case -1:
			order = "decreasing"
		}
		Fail(t, fmt.Sprintf("The 'object' !!is not %s!!.", order), NewObjectsSingleUnknown(object), msg...)
	}
}

func AssertRegexpHelper(t testRunner, regex any, txt any, shouldMatch bool, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	regexString := fmt.Sprint(regex)
	txtString := fmt.Sprint(txt)
	match, _ := regexp.MatchString(regexString, txtString)
	if shouldMatch != match {
		failText := "!!does not match!! the string."
		if !shouldMatch {
			failText = "!!does match!! the string, but !!should not!!."
		}
		Fail(t, "The regex pattern "+failText, Objects{
			NewObjectsSingleNamed("Regex Pattern", regexString+"\n")[0],
			NewObjectsSingleNamed("String", txtString+"\n")[0],
		}, msg...)
	}
}

// AssertDirEmptyHelper checks for io.EOF to determine if directory empty or not
func AssertDirEmptyHelper(t testRunner, dir string) bool {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	f, err := os.Open(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return true
		}
		Fail(t, "Error opening directory specified", NewObjectsSingleNamed("dir", dir), err.Error())
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	return errors.Is(err, io.EOF)
}

func IsList(list any) bool {
	kind := reflect.TypeOf(list).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return false
	}

	return true
}

func HasSameElements(expected any, actual any) bool {
	if assert.Nil(expected) || assert.Nil(actual) {
		return expected == actual
	}

	if !IsList(expected) || !IsList(actual) {
		return false
	}

	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	expectedLen := expectedValue.Len()
	actualLen := actualValue.Len()

	visited := make([]bool, actualLen)

	var extraA, extraB []any
	for i := 0; i < expectedLen; i++ {
		element := expectedValue.Index(i).Interface()
		found := false
		for j := 0; j < actualLen; j++ {
			if visited[j] {
				continue
			}
			if IsEqual(actualValue.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			extraA = append(extraA, element)
		}
	}

	for j := 0; j < actualLen; j++ {
		if visited[j] {
			continue
		}
		extraB = append(extraB, actualValue.Index(j).Interface())
	}

	if len(extraA) == 0 && len(extraB) == 0 {
		return true
	}

	return false
}

func IsSubset(t testRunner, list any, subset any) bool {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Nil(subset) {
		return true
	}

	if !IsList(list) {
		Fail(t, "The first argument is not a list.", NewObjectsSingleNamed("First argument", list))
	}

	if !IsList(subset) {
		Fail(t, "The second argument is not a list.", NewObjectsSingleNamed("Second argument", subset))
	}

	subsetValue := reflect.ValueOf(subset)

	for i := 0; i < subsetValue.Len(); i++ {
		element := subsetValue.Index(i).Interface()
		contains := DoesContain(list, element)

		if !contains {
			return false
		}
	}

	return true
}
