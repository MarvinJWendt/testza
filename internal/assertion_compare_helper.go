package internal

import (
	"fmt"
	"reflect"
)

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
		var txt string
		switch option {
		case 1:
			txt = "increasing"
		case 0:
			txt = "equal"
		case -1:
			txt = "decreasing"
		}
		Fail(t, fmt.Sprintf("The 'object' !!is not %s!!.", txt), NewObjectsSingleUnknown(object), msg...)
	}
}
