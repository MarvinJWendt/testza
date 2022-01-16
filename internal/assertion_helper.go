package internal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// IsKind returns if an object is kind of a specific kind.
func IsKind(expectedKind reflect.Kind, value interface{}) bool {
	return reflect.TypeOf(value).Kind() == expectedKind
}

// IsNil checks if an object is nil.
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

// IsNumber checks if the value is of a numeric kind.
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

// IsZero checks if a value is the zero value of its type.
func IsZero(value interface{}) bool {
	return value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

// IsEqual checks if two objects are equal.
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

// HasEqualValues checks if two objects have equal values.
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

// DoesImplement checks if an objects implements an interface.
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

// DoesContain checks that ab objects contains an element.
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

// AssertCompareHelper option: 1 = increasing, 0 = equal, -1 = decreasing
func AssertCompareHelper(t testRunner, object interface{}, option int, msg ...interface{}) {
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

func AssertRegexpHelper(t testRunner, regex interface{}, txt interface{}, shouldMatch bool, msg ...interface{}) {
	regexString := fmt.Sprint(regex)
	txtString := fmt.Sprint(txt)
	match, _ := regexp.MatchString(regexString, txtString)
	if shouldMatch != match {
		failText := "!!does not match!! the string."
		if !shouldMatch {
			failText = "!!does match!! the string, but !!should not!!."
		}
		Fail(t, "The regex pattern "+failText, Objects{
			{
				Name:      "Regex Pattern",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      regexString + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
			{
				Name:      "String",
				NameStyle: pterm.NewStyle(pterm.FgRed),
				Data:      txtString + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
		}, msg...)
	}
}

// AssertDirEmptyHelper checks for io.EOF to determine if directory empty or not
func AssertDirEmptyHelper(t testRunner, dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		Fail(t, "Error opening directory specified", NewObjectsSingleNamed("dir", dir))
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	return errors.Is(err, io.EOF)
}

func IsList(list interface{}) bool {
	kind := reflect.TypeOf(list).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return false
	}

	return true
}

func HasSameElements(expected interface{}, actual interface{}) bool {
	if IsNil(expected) || IsNil(actual) {
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

	var extraA, extraB []interface{}
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

func IsSubset(t testRunner, list interface{}, subset interface{}) bool {
	if IsNil(subset) {
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
