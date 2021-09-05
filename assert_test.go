package testza_test

import (
	"errors"
	"fmt"
	"go/types"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/pterm/pterm"

	. "github.com/MarvinJWendt/testza"
)

type assertionTestStruct struct {
	Name string
	Age  int
	Meta assertionTestStructNested
}

type assertionTestStructNested struct {
	ID    int
	Admin bool
}

func randomString() string {
	return MockInputStringGenerateRandom(1, rand.Intn(10))[0]
}

func generateStruct() (ret assertionTestStruct) {
	ret.Name = randomString()
	ret.Age = rand.Intn(40)
	ret.Meta.ID = rand.Intn(10_000)
	ret.Meta.Admin = rand.Intn(1) == 1

	return
}

func testEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	t.Run("Equal", func(t *testing.T) {
		t.Helper()

		AssertEqual(t, expected, actual)
	})

	t.Run("EqualValues", func(t *testing.T) {
		t.Helper()

		AssertEqualValues(t, expected, actual)
	})
}

func TestAssertKindOf(t *testing.T) {
	tests := []struct {
		kind  reflect.Kind
		value interface{}
	}{
		{kind: reflect.String, value: "Hello, World!"},
		{kind: reflect.Int, value: 1337},
		{kind: reflect.Bool, value: true},
		{kind: reflect.Bool, value: false},
		{kind: reflect.Map, value: make(map[string]int)},
		{kind: reflect.Func, value: func() {}},
		{kind: reflect.Struct, value: testing.T{}},
		{kind: reflect.Slice, value: []string{}},
		{kind: reflect.Slice, value: []int{}},
		{kind: reflect.Slice, value: []float64{}},
		{kind: reflect.Slice, value: []float32{}},
		{kind: reflect.Float64, value: 13.37},
		{kind: reflect.Float32, value: float32(13.37)},
	}
	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			AssertKindOf(t, test.kind, test.value)
		})
	}
}

func TestAssertKindOf_fails(t *testing.T) {
	tests := []struct {
		kind  reflect.Kind
		value interface{}
	}{
		{kind: reflect.Int, value: "Hello, World!"},
		{kind: reflect.Bool, value: 1337},
		{kind: reflect.Float64, value: true},
		{kind: reflect.Float32, value: false},
		{kind: reflect.Struct, value: make(map[string]int)},
		{kind: reflect.Map, value: func() {}},
		{kind: reflect.Slice, value: testing.T{}},
		{kind: reflect.Struct, value: []string{}},
		{kind: reflect.Array, value: []int{}},
		{kind: reflect.Chan, value: []float64{}},
		{kind: reflect.Complex64, value: []float32{}},
		{kind: reflect.Complex128, value: 13.37},
		{kind: reflect.Float64, value: float32(13.37)},
	}
	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertKindOf(t, test.kind, test.value)
			})
		})
	}
}

func TestAssertNotKindOf(t *testing.T) {
	tests := []struct {
		kind  reflect.Kind
		value interface{}
	}{
		{kind: reflect.Int, value: "Hello, World!"},
		{kind: reflect.Bool, value: 1337},
		{kind: reflect.Float64, value: true},
		{kind: reflect.Float32, value: false},
		{kind: reflect.Struct, value: make(map[string]int)},
		{kind: reflect.Map, value: func() {}},
		{kind: reflect.Slice, value: testing.T{}},
		{kind: reflect.Struct, value: []string{}},
		{kind: reflect.Array, value: []int{}},
		{kind: reflect.Chan, value: []float64{}},
		{kind: reflect.Complex64, value: []float32{}},
		{kind: reflect.Complex128, value: 13.37},
		{kind: reflect.Float64, value: float32(13.37)},
	}
	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			AssertNotKindOf(t, test.kind, test.value)
		})
	}
}

func TestAssertNotKindOf_fails(t *testing.T) {
	tests := []struct {
		kind  reflect.Kind
		value interface{}
	}{
		{kind: reflect.String, value: "Hello, World!"},
		{kind: reflect.Int, value: 1337},
		{kind: reflect.Bool, value: true},
		{kind: reflect.Bool, value: false},
		{kind: reflect.Map, value: make(map[string]int)},
		{kind: reflect.Func, value: func() {}},
		{kind: reflect.Struct, value: testing.T{}},
		{kind: reflect.Slice, value: []string{}},
		{kind: reflect.Slice, value: []int{}},
		{kind: reflect.Slice, value: []float64{}},
		{kind: reflect.Slice, value: []float32{}},
		{kind: reflect.Float64, value: 13.37},
		{kind: reflect.Float32, value: float32(13.37)},
	}
	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertNotKindOf(t, test.kind, test.value)
			})
		})
	}
}

func TestAssertNumeric(t *testing.T) {
	var numbers []interface{}

	for i := 0; i < 10; i++ {
		numbers = append(numbers,
			i,
			int8(i),
			int16(i),
			int32(i),
			int64(i),
			uint(i),
			uint8(i),
			uint16(i),
			uint32(i),
			uint64(i),
			float32(i)+0.25,
			float32(i)+0.5,
			float64(i)+0.25,
			float64(i)+0.5,
			complex64(complex(float64(i), 2)),
			complex(float64(i), 2),
		)
	}

	for _, number := range numbers {
		t.Run(pterm.Sprintf("Type=%s;Value=%#v", reflect.TypeOf(number).Kind().String(), number), func(t *testing.T) {
			AssertNumeric(t, number)
		})
	}
}

func TestAssertNumeric_fails(t *testing.T) {
	noNumbers := []interface{}{"Hello, World!", true, false}
	for _, number := range noNumbers {
		t.Run(pterm.Sprintf("Type=%s;Value=%#v", reflect.TypeOf(number).Kind().String(), number), func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertNumeric(t, number)
			})
		})
	}
}

func TestAssertNotNumeric_fails(t *testing.T) {
	var numbers []interface{}

	for i := 0; i < 10; i++ {
		numbers = append(numbers,
			i,
			int8(i),
			int16(i),
			int32(i),
			int64(i),
			uint(i),
			uint8(i),
			uint16(i),
			uint32(i),
			uint64(i),
			float32(i)+0.25,
			float32(i)+0.5,
			float64(i)+0.25,
			float64(i)+0.5,
			complex64(complex(float64(i), 2)),
			complex(float64(i), 2),
		)
	}

	for _, number := range numbers {
		t.Run(pterm.Sprintf("Type=%s;Value=%#v", reflect.TypeOf(number).Kind().String(), number), func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertNotNumeric(t, number)
			})
		})
	}
}

func TestAssertZero(t *testing.T) {
	var zeroValues []interface{}
	zeroValues = append(zeroValues, 0, "", false, nil)

	for i, value := range zeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			AssertZero(t, value)
		})
	}
}

func TestAssertZero_fails(t *testing.T) {
	var zeroValues []interface{}
	zeroValues = append(zeroValues, true, "asd", 123, 1.5, 'a')

	for i, value := range zeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertZero(t, value)
			})
		})
	}
}

func TestAssertNotZero(t *testing.T) {
	var zeroValues []interface{}
	zeroValues = append(zeroValues, true, "asd", 123, 1.5, 'a')

	for i, value := range zeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			AssertNotZero(t, value)
		})
	}
}

func TestAssertNotZero_fails(t *testing.T) {
	var zeroValues []interface{}
	zeroValues = append(zeroValues, 0, "", false, nil)

	for i, value := range zeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertNotZero(t, value)
			})
		})
	}
}

func TestAssertEqual(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
			t.Helper()

			testEqual(t, str, str)
		})
	})

	t.Run("Nil and Nil are equal", func(t *testing.T) {
		AssertEqual(t, nil, nil)
	})

	t.Run("Equal pointers to same struct", func(t *testing.T) {
		s := generateStruct()
		AssertEqual(t, &s, &s)
	})

	t.Run("Equal dereference to same pointer struct", func(t *testing.T) {
		s := generateStruct()
		s2 := &s
		AssertEqual(t, *s2, *s2)
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			testEqual(t, s, s)
		}
	})
}

func TestAssertEqual_fails(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertEqual(t, str, str+" addon")
			})
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			s2 := s
			s2.Name += " addon"
			t.Run("NotEqual", func(t *testing.T) {
				AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
					AssertEqual(t, s, s2)
				})
			})
			t.Run("NotEqualValues", func(t *testing.T) {
				AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
					AssertEqualValues(t, s, s2)
				})
			})
		}
	})
}

func TestAssertNotEqual(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
			AssertNotEqual(t, str, str+" addon")
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			s2 := s
			s2.Name += " addon"
			t.Run("NotEqual", func(t *testing.T) {
				AssertNotEqual(t, s, s2)
			})
			t.Run("NotEqualValues", func(t *testing.T) {
				AssertNotEqualValues(t, s, s2)
			})
		}
	})
}

func TestAssertNotEqual_fails(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
			t.Helper()

			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertNotEqual(t, str, str)
			})
		})
	})

	t.Run("Nil and Nil are equal", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			AssertNotEqual(t, nil, nil)
		})
	})

	t.Run("Equal pointers to same struct", func(t *testing.T) {
		s := generateStruct()
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			AssertNotEqual(t, &s, &s)
		})
	})

	t.Run("Equal dereference to same pointer struct", func(t *testing.T) {
		s := generateStruct()
		s2 := &s
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			AssertNotEqual(t, *s2, *s2)
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertNotEqual(t, s, s)
			})
		}
	})
}

func TestAssertEqualValues(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
			AssertEqualValues(t, str, str)
		})
	})
}

func TestAssertEqualValues_fails(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertEqualValues(t, str, str+" addon")
			})
		})
	})
}

func TestAssertNotEqualValues(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
			AssertNotEqualValues(t, str, str+" addon")
		})
	})
}

func TestAssertNotEqualValues_fails(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertNotEqualValues(t, str, str)
			})
		})
	})
}

func TestAssertTrue(t *testing.T) {
	AssertTrue(t, true)
}

func TestAssertTrue_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertTrue(t, false)
	})
}

func TestAssertFalse(t *testing.T) {
	AssertFalse(t, false)
}

func TestAssertFalse_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertFalse(t, true)
	})
}

func TestAssertImplements(t *testing.T) {
	t.Run("ConstImplementsStringer", func(t *testing.T) {
		AssertImplements(t, (*fmt.Stringer)(nil), new(types.Const))
	})
}

func TestAssertImplements_fails(t *testing.T) {
	t.Run("assertionTestStructNotImplementsFmtStringer", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			AssertImplements(t, (*fmt.Stringer)(nil), new(assertionTestStruct))
		})
	})
}

func TestAssertNotImplements(t *testing.T) {
	t.Run("assertionTestStructNotImplementsFmtStringer", func(t *testing.T) {
		AssertNotImplements(t, (*fmt.Stringer)(nil), new(assertionTestStruct))
	})
}

func TestAssertNotImplements_fails(t *testing.T) {
	t.Run("ConstImplementsStringer", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			AssertNotImplements(t, (*fmt.Stringer)(nil), new(types.Const))
		})
	})
}

func TestAssertContains(t *testing.T) {
	s := []string{"Hello", "World"}

	AssertContains(t, s, "World")
}

func TestAssertContains_fails(t *testing.T) {
	s := []string{"Hello", "World"}

	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertContains(t, s, "asdasd")
	})
}

func TestAssertNotContains(t *testing.T) {
	s := []string{"Hello", "World"}

	AssertNotContains(t, s, "asdasd")
}

func TestAssertNotContains_fails(t *testing.T) {
	s := []string{"Hello", "World"}
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertNotContains(t, s, "World")
	})
}

func TestAssertPanics(t *testing.T) {
	AssertPanics(t, func() {
		panic("TestPanic")
	})
}

func TestAssertPanics_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertPanics(t, func() {
			// If we do nothing here it can't panic ;)
		})
	})
}

func TestAssertNotPanics(t *testing.T) {
	AssertNotPanics(t, func() {
		// If we do nothing here it can't panic ;)
	})
}

func TestAssertNotPanics_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertNotPanics(t, func() {
			panic("TestPanic")
		})
	})
}

func TestAssertNil(t *testing.T) {
	AssertNil(t, nil)
}

func TestAssertNil_pointer(t *testing.T) {
	var a *int = nil
	AssertNil(t, a)
}

func TestAssertNil_interface(t *testing.T) {
	var i *int = nil
	var a interface{} = i
	AssertNil(t, a)
}

func TestAssertNil_pointer_struct(t *testing.T) {
	var a *assertionTestStruct
	AssertNil(t, a)
}

func TestAssertNil_fails(t *testing.T) {
	objectsNotNil := []interface{}{"asd", 0, false, true, 'c', "", []int{1, 2, 3}}

	for _, v := range objectsNotNil {
		t.Run(pterm.Sprintf("Value=%#v", v), func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertNil(t, v)
			})
		})
	}
}

func TestAssertNotNil(t *testing.T) {
	objectsNotNil := []interface{}{"asd", 0, false, true, 'c', "", []int{1, 2, 3}}

	for _, v := range objectsNotNil {
		t.Run(pterm.Sprintf("Value=%#v", v), func(t *testing.T) {
			AssertNotNil(t, v)
		})
	}
}

func TestAssertNotNil_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertNotNil(t, nil)
	})
}

func TestAssertCompletesIn(t *testing.T) {
	AssertCompletesIn(t, 50*time.Millisecond, func() {
		time.Sleep(5 * time.Millisecond)
	})
}

func TestAssertCompletesIn_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertCompletesIn(t, 10*time.Microsecond, func() {
			time.Sleep(15 * time.Millisecond)
		})
	})
}

func TestAssertNotCompletesIn(t *testing.T) {
	AssertNotCompletesIn(t, 10*time.Microsecond, func() {
		time.Sleep(15 * time.Millisecond)
	})
}

func TestAssertNotCompletesIn_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertNotCompletesIn(t, 50*time.Millisecond, func() {
			time.Sleep(5 * time.Millisecond)
		})
	})
}

func TestAssertNoError(t *testing.T) {
	AssertNoError(t, nil)
}

func TestAssertNoError_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertNoError(t, errors.New("hello error"))
	})
}

func TestAssertGreater(t *testing.T) {
	AssertGreater(t, 2, 1)
	AssertGreater(t, 5, 4)
}

func TestAssertGreater_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertGreater(t, 1, 2)
	})

	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertGreater(t, 4, 5)
	})
}

func TestAssertLess(t *testing.T) {
	AssertLess(t, 1, 2)
	AssertLess(t, 4, 5)
}

func TestAssertLess_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertLess(t, 2, 1)
	})

	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertLess(t, 5, 4)
	})
}

func TestAssertTestFails(t *testing.T) {
	t.Run("Wrong assertion should fail", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			AssertTrue(t, false)
		})
	})

	t.Run(".Error should make the test pass", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Error()
		})
	})

	t.Run(".Errorf should make the test pass", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Errorf("")
		})
	})

	t.Run(".Fail should make the test pass", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Fail()
		})
	})

	t.Run(".FailNow should make the test pass", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			t.FailNow()
		})
	})

	t.Run(".Fatal should make the test pass", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Fatal()
		})
	})

	t.Run(".Fatalf should make the test pass", func(t *testing.T) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Fatalf("")
		})
	})
}

func TestAssertTestFails_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
			AssertTrue(t, true)
		})
	})
}

func TestAssertErrorIs(t *testing.T) {
	var testErr = errors.New("hello world")
	var testErrWrapped = fmt.Errorf("test err: %w", testErr)
	AssertErrorIs(t, testErrWrapped, testErr)
}

func TestAssertErrorIs_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		var testErr = errors.New("hello world")
		var test2Err = errors.New("hello world 2")
		var testErrWrapped = fmt.Errorf("test err: %w", testErr)
		AssertErrorIs(t, testErrWrapped, test2Err)
	})
}

func TestAssertNotErrorIs(t *testing.T) {
	var testErr = errors.New("hello world")
	var test2Err = errors.New("hello world 2")
	var testErrWrapped = fmt.Errorf("test err: %w", testErr)
	AssertNotErrorIs(t, testErrWrapped, test2Err)
}

func TestAssertNotErrorIs_fails(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		var testErr = errors.New("hello world")
		var testErrWrapped = fmt.Errorf("test err: %w", testErr)
		AssertNotErrorIs(t, testErrWrapped, testErr)
	})
}

func TestAssertLen(t *testing.T) {
	tests := []struct {
		object interface{}
		expLen int
	}{
		{object: "", expLen: 0},
		{object: "a", expLen: 1},
		{object: "123", expLen: 3},
		{object: []string{"", "", "", ""}, expLen: 4},
		{object: []string{}, expLen: 0},
		{object: []int{1, 2, 3, 4, 5}, expLen: 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.object), func(t *testing.T) {
			AssertLen(t, test.object, test.expLen)
		})
	}
}

func TestAssertLen_fails(t *testing.T) {
	tests := []struct {
		object interface{}
		expLen int
	}{
		{object: "", expLen: 1},
		{object: "a", expLen: 2},
		{object: "123", expLen: 4},
		{object: []string{"", "", "", ""}, expLen: 5},
		{object: []string{}, expLen: 1},
		{object: []int{1, 2, 3, 4, 5}, expLen: 4},
		{object: 1, expLen: 4},
		{object: 0.1, expLen: 4},
		{object: generateStruct(), expLen: 1},
		{object: interface{}(nil), expLen: 1},
		{object: true, expLen: 0},
		{object: uint(137), expLen: 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.object), func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertLen(t, test.object, test.expLen)
			})
		})
	}
}

func TestAssertLen_full(t *testing.T) {
	MockInputStringRunTests(t, MockInputStringFull(), func(t *testing.T, index int, str string) {
		AssertLen(t, str, len(str))
	})
}

func TestAssertIsIncreasing(t *testing.T) {
	tests := []struct {
		name   string
		object interface{}
	}{
		{name: "int", object: []int{1, 2, 3, 4, 5, 6}},
		{name: "int2", object: []int{1, 7, 10, 11, 134, 700432}},
		{name: "int_negative", object: []int{-2, -1, 0, 1, 2, 3, 4, 5, 6}},
		{name: "int8", object: []int8{1, 2, 3, 4, 5, 6}},
		{name: "int16", object: []int16{1, 2, 3, 4, 5, 6}},
		{name: "int32", object: []int32{1, 2, 3, 4, 5, 6}},
		{name: "int64", object: []int64{1, 2, 3, 4, 5, 6}},
		{name: "uint", object: []uint{1, 2, 3, 4, 5, 6}},
		{name: "uint8", object: []uint8{1, 2, 3, 4, 5, 6}},
		{name: "uint16", object: []uint16{1, 2, 3, 4, 5, 6}},
		{name: "uint32", object: []uint32{1, 2, 3, 4, 5, 6}},
		{name: "uint64", object: []uint64{1, 2, 3, 4, 5, 6}},
		{name: "float32", object: []float32{1.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "float64", object: []float64{1.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			AssertIncreasing(t, test.object)
		})
	}
}

func TestAssertIsIncreasing_fail(t *testing.T) {
	tests := []struct {
		name   string
		object interface{}
	}{
		{name: "int_empty", object: []int{}},
		{name: "int_one", object: []int{4}},
		{name: "int", object: []int{4, 2, 3, 4, 5, 6}},
		{name: "int8", object: []int8{4, 2, 3, 4, 5, 6}},
		{name: "int16", object: []int16{4, 2, 3, 4, 5, 6}},
		{name: "int32", object: []int32{4, 2, 3, 4, 5, 6}},
		{name: "int64", object: []int64{4, 2, 3, 4, 5, 6}},
		{name: "uint", object: []uint{4, 2, 3, 4, 5, 6}},
		{name: "uint8", object: []uint8{4, 2, 3, 4, 5, 6}},
		{name: "uint16", object: []uint16{4, 2, 3, 4, 5, 6}},
		{name: "uint32", object: []uint32{4, 2, 3, 4, 5, 6}},
		{name: "uint64", object: []uint64{4, 2, 3, 4, 5, 6}},
		{name: "float32", object: []float32{4.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "float64", object: []float64{4.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "complex64", object: []complex64{complex64(complex(float64(4), 2)), complex64(complex(float64(2), 2)), complex64(complex(float64(4), 2))}},
		{name: "complex128", object: []complex128{complex(400.0, 2), complex(133.4, 2), complex(144.4, 2)}},
		{name: "stringSlice", object: []string{"a"}},
		{name: "string", object: "a"},
		{name: "bool", object: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertIncreasing(t, test.object)
			})
		})
	}
}

func TestAssertIsDecreasing(t *testing.T) {
	tests := []struct {
		name   string
		object interface{}
	}{
		{name: "int", object: []int{6, 5, 4, 3, 2, 1}},
		{name: "int2", object: []int{700432, 134, 11, 10, 7, 1}},
		{name: "int_negative", object: []int{6, 5, 4, 3, 2, 1, 0, -1, -2}},
		{name: "int8", object: []int8{6, 5, 4, 3, 2, 1}},
		{name: "int16", object: []int16{6, 5, 4, 3, 2, 1}},
		{name: "int32", object: []int32{6, 5, 4, 3, 2, 1}},
		{name: "int64", object: []int64{6, 5, 4, 3, 2, 1}},
		{name: "uint", object: []uint{6, 5, 4, 3, 2, 1}},
		{name: "uint8", object: []uint8{6, 5, 4, 3, 2, 1}},
		{name: "uint16", object: []uint16{6, 5, 4, 3, 2, 1}},
		{name: "uint32", object: []uint32{6, 5, 4, 3, 2, 1}},
		{name: "uint64", object: []uint64{6, 5, 4, 3, 2, 1}},
		{name: "float32", object: []float32{5.45623, 4.12345, 3.0, 2.1, 1.0}},
		{name: "float64", object: []float64{5.45623, 4.12345, 3.0, 2.1, 1.0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			AssertDecreasing(t, test.object)
		})
	}
}

func TestAssertIsDecreasing_fail(t *testing.T) {
	tests := []struct {
		name   string
		object interface{}
	}{
		{name: "int_empty", object: []int{}},
		{name: "int_one", object: []int{4}},
		{name: "int", object: []int{4, 2, 3, 4, 5, 6}},
		{name: "int8", object: []int8{4, 2, 3, 4, 5, 6}},
		{name: "int16", object: []int16{4, 2, 3, 4, 5, 6}},
		{name: "int32", object: []int32{4, 2, 3, 4, 5, 6}},
		{name: "int64", object: []int64{4, 2, 3, 4, 5, 6}},
		{name: "uint", object: []uint{4, 2, 3, 4, 5, 6}},
		{name: "uint8", object: []uint8{4, 2, 3, 4, 5, 6}},
		{name: "uint16", object: []uint16{4, 2, 3, 4, 5, 6}},
		{name: "uint32", object: []uint32{4, 2, 3, 4, 5, 6}},
		{name: "uint64", object: []uint64{4, 2, 3, 4, 5, 6}},
		{name: "float32", object: []float32{4.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "float64", object: []float64{4.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "complex64", object: []complex64{complex64(complex(float64(4), 2)), complex64(complex(float64(2), 2)), complex64(complex(float64(4), 2))}},
		{name: "complex128", object: []complex128{complex(400.0, 2), complex(133.4, 2), complex(144.4, 2)}},
		{name: "stringSlice", object: []string{"a"}},
		{name: "string", object: "a"},
		{name: "bool", object: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
				AssertDecreasing(t, test.object)
			})
		})
	}
}

func TestAssertRegexp(t *testing.T) {
	AssertRegexp(t, "p([a-z]+)ch", "peache")
}

func TestAssertRegexp_fail(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertRegexp(t, "p([a-z]+)ch", "peahe")
	})
}

func TestAssertNotRegexp(t *testing.T) {
	AssertNotRegexp(t, "p([a-z]+)ch", "peahe")
}

func TestAssertNotRegexp_fail(t *testing.T) {
	AssertTestFails(t, func(t TestingPackageWithFailFunctions) {
		AssertNotRegexp(t, "p([a-z]+)ch", "peache")
	})
}
