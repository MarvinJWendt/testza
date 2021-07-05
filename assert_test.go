package testza

import (
	"fmt"
	"go/types"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/pterm/pterm"
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
	return Use.Mock.Strings.GenerateRandom(1, rand.Intn(10))[0]
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

		Use.Assert.Equal(t, expected, actual)
	})

	t.Run("EqualValues", func(t *testing.T) {
		t.Helper()

		Use.Assert.EqualValues(t, expected, actual)
	})
}

func TestAssertHelper_KindOf(t *testing.T) {
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
			Use.Assert.KindOf(t, test.kind, test.value)
		})
	}
}

func TestAssertHelper_NotKindOf(t *testing.T) {
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
			Use.Assert.NotKindOf(t, test.kind, test.value)
		})
	}
}

func TestAssertHelper_Numeric(t *testing.T) {
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
			Use.Assert.Numeric(t, number)
		})
	}
}

func TestAssertHelper_NotNumber(t *testing.T) {
	noNumbers := []interface{}{"Hello, World!", true, false}
	for _, number := range noNumbers {
		t.Run(pterm.Sprintf("Type=%s;Value=%#v", reflect.TypeOf(number).Kind().String(), number), func(t *testing.T) {
			Use.Assert.NotNumeric(t, number)
		})
	}
}

func TestAssertHelper_Zero(t *testing.T) {
	var zeroValues []interface{}
	zeroValues = append(zeroValues, 0, "", false, nil)

	for i, value := range zeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			Use.Assert.Zero(t, value)
		})
	}
}

func TestAssertHelper_NotZero(t *testing.T) {
	var zeroValues []interface{}
	zeroValues = append(zeroValues, true, "asd", 123, 1.5, 'a')

	for i, value := range zeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			Use.Assert.NotZero(t, value)
		})
	}
}

func TestAssertHelper_Equal(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		Use.Mock.Strings.RunTests(t, Use.Mock.Strings.Full(), func(t *testing.T, index int, str string) {
			t.Helper()

			testEqual(t, str, str)
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			testEqual(t, s, s)
		}
	})
}

func TestAssertHelper_NotEqual(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		Use.Mock.Strings.RunTests(t, Use.Mock.Strings.Full(), func(t *testing.T, index int, str string) {
			Use.Assert.NotEqual(t, str, str+" addon")
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			s2 := s
			s2.Name += " addon"
			t.Run("NotEqual", func(t *testing.T) {
				Use.Assert.NotEqual(t, s, s2)
			})
			t.Run("NotEqualValues", func(t *testing.T) {
				Use.Assert.NotEqualValues(t, s, s2)
			})
		}
	})
}

func TestAssert_EqualValues(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		Use.Mock.Strings.RunTests(t, Use.Mock.Strings.Full(), func(t *testing.T, index int, str string) {
			Use.Assert.EqualValues(t, str, str)
		})
	})
}

func TestAssert_NotEqualValues(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		Use.Mock.Strings.RunTests(t, Use.Mock.Strings.Full(), func(t *testing.T, index int, str string) {
			Use.Assert.NotEqualValues(t, str, str+" addon")
		})
	})
}

func TestAssertHelper_True(t *testing.T) {
	Use.Assert.True(t, true)
}

func TestAssertHelper_False(t *testing.T) {
	Use.Assert.False(t, false)
}

func TestAssert_Implements(t *testing.T) {
	t.Run("ConstImplementsStringer", func(t *testing.T) {
		Use.Assert.Implements(t, (*fmt.Stringer)(nil), new(types.Const))
	})
}

func TestAssert_NotImplements(t *testing.T) {
	t.Run("assertionTestStructNotImplementsFmtStringer", func(t *testing.T) {
		Use.Assert.NotImplements(t, (*fmt.Stringer)(nil), new(assertionTestStruct))
	})
}

func TestAssertHelper_Contains(t *testing.T) {
	s := []string{"Hello", "World"}

	Use.Assert.Contains(t, s, "World")
}

func TestAssertHelper_NotContains(t *testing.T) {
	s := []string{"Hello", "World"}

	Use.Assert.NotContains(t, s, "asdasd")
}

func TestAssertHelper_Panic(t *testing.T) {
	Use.Assert.Panic(t, func() {
		panic("TestPanic")
	})
}

func TestAssertHelper_NotPanic(t *testing.T) {
	Use.Assert.NotPanic(t, func() {
		// If we do nothing here it can't panic ;)
	})
}

func TestAssertHelper_Nil(t *testing.T) {
	Use.Assert.Nil(t, nil)
}

func TestAssertHelper_NotNil(t *testing.T) {
	objectsNotNil := []interface{}{"asd", 0, false, true, 'c', "", []int{1, 2, 3}}

	for _, v := range objectsNotNil {
		t.Run(pterm.Sprintf("Value=%#v", v), func(t *testing.T) {
			Use.Assert.NotNil(t, v)
		})
	}
}

func TestAssertHelper_CompletesIn(t *testing.T) {
	Use.Assert.CompletesIn(t, 50*time.Millisecond, func() {
		time.Sleep(5 * time.Millisecond)
	})
}

func TestAssertHelper_NotCompletesIn(t *testing.T) {
	Use.Assert.NotCompletesIn(t, 10*time.Microsecond, func() {
		time.Sleep(15 * time.Millisecond)
	})
}
