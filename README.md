<h1 align="center">testza 🍕</h1>
<p align="center">Testza is like pizza for Go - you could live without it, but why should you?</p>

<p align="center">

<a href="https://github.com/MarvinJWendt/testza/releases">
<img src="https://img.shields.io/github/v/release/MarvinJWendt/testza?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/MarvinJWendt/testza" target="_blank">
<img src="https://img.shields.io/github/workflow/status/MarvinJWendt/testza/Go?label=tests&style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/MarvinJWendt/testza" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/MarvinJWendt/testza?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/MarvinJWendt/testza">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-2527-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>
  
<a href="https://pkg.go.dev/github.com/MarvinJWendt/testza" target="_blank">
<img src="https://pkg.go.dev/badge/github.com/MarvinJWendt/testza.svg" alt="Go Reference">
</a>

<a href="https://goreportcard.com/report/github.com/MarvinJWendt/testza" target="_blank">
<img src="https://goreportcard.com/badge/github.com/MarvinJWendt/testza" alt="Go report">
</a>  

</p>


---

<p align="center">
<strong><a href="#install">Get The Module</a></strong>
|
<strong><a href="https://github.com/MarvinJWendt/testza#documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<img align="right" height="400" alt="Screenshot of an example test message" src="https://user-images.githubusercontent.com/31022056/124531029-ea31b780-de0d-11eb-8984-74e679f84aec.png" />

## Installation

```console
# Execute this command inside your project
go get github.com/MarvinJWendt/testza
```

## Description

Testza is a full-featured testing framework for Go.
It integrates with the default test runner, so you can use it with the standard `go test` tool.
Testza contains easy to use methods, like assertions, output capturing, mocking, and much more.

The main goal of testza is to provide an easy and fun experience writing tests and providing a nice, user-friendly output.
Even developers who never used testza, will get into it quickly.

## Getting Started

See the examples below for a quick introduction!

```go
// --- Some Examples ---

// - Some assertions -
testza.AssertTrue(t, true) // -> Pass
testza.AssertNoError(t, err) // -> Pass
testza.AssertEqual(t, object, object) // -> Pass
// ...

// - Testing console output -
// Test the output of your CLI tool easily!
terminalOutput, _ := testza.CaptureStdout(func(w io.Writer) error {fmt.Println("Hello"); return nil})
testza.AssertEqual(t, terminalOutput, "Hello\n") // -> Pass

// - Mocking -
// Testing a function that accepts email addresses as a parameter:

// Testset of many different email addresses
emailAddresses := testza.MockStringEmailAddresses()

// Run a test for every string in the test set
testza.MockStringRunTests(t, emailAddresses, func(t *testing.T, index int, str string) {
  user, domain, err := internal.ParseEmailAddress(str) // Use your function
  testza.AssertNoError(t, err) // Assert that your function does not return an error
  testza.AssertNotZero(t, user) // Assert that the user is returned
  testza.AssertNotZero(t, domain) // Assert that the domain is returned
})

// And that's just a few examples of what you can do with Testza!
```

## Documentation

<!-- docs:start -->
<table>
  <tr>
    <th>Module</th>
    <th>Methods</th>
  </tr><tr>
<td><a href="https://github.com/MarvinJWendt/testza#Settings">Settings</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [SetColorsEnabled](https://github.com/MarvinJWendt/testza#SetColorsEnabled)
  - [SetLineNumbersEnabled](https://github.com/MarvinJWendt/testza#SetLineNumbersEnabled)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Assert">Assert</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [AssertCompletesIn](https://github.com/MarvinJWendt/testza#AssertCompletesIn)
  - [AssertContains](https://github.com/MarvinJWendt/testza#AssertContains)
  - [AssertEqual](https://github.com/MarvinJWendt/testza#AssertEqual)
  - [AssertEqualValues](https://github.com/MarvinJWendt/testza#AssertEqualValues)
  - [AssertFalse](https://github.com/MarvinJWendt/testza#AssertFalse)
  - [AssertGreater](https://github.com/MarvinJWendt/testza#AssertGreater)
  - [AssertImplements](https://github.com/MarvinJWendt/testza#AssertImplements)
  - [AssertKindOf](https://github.com/MarvinJWendt/testza#AssertKindOf)
  - [AssertLess](https://github.com/MarvinJWendt/testza#AssertLess)
  - [AssertNil](https://github.com/MarvinJWendt/testza#AssertNil)
  - [AssertNoError](https://github.com/MarvinJWendt/testza#AssertNoError)
  - [AssertNotCompletesIn](https://github.com/MarvinJWendt/testza#AssertNotCompletesIn)
  - [AssertNotContains](https://github.com/MarvinJWendt/testza#AssertNotContains)
  - [AssertNotEqual](https://github.com/MarvinJWendt/testza#AssertNotEqual)
  - [AssertNotEqualValues](https://github.com/MarvinJWendt/testza#AssertNotEqualValues)
  - [AssertNotImplements](https://github.com/MarvinJWendt/testza#AssertNotImplements)
  - [AssertNotKindOf](https://github.com/MarvinJWendt/testza#AssertNotKindOf)
  - [AssertNotNil](https://github.com/MarvinJWendt/testza#AssertNotNil)
  - [AssertNotNumeric](https://github.com/MarvinJWendt/testza#AssertNotNumeric)
  - [AssertNotPanics](https://github.com/MarvinJWendt/testza#AssertNotPanics)
  - [AssertNotZero](https://github.com/MarvinJWendt/testza#AssertNotZero)
  - [AssertNumeric](https://github.com/MarvinJWendt/testza#AssertNumeric)
  - [AssertPanics](https://github.com/MarvinJWendt/testza#AssertPanics)
  - [AssertTestFails](https://github.com/MarvinJWendt/testza#AssertTestFails)
  - [AssertTrue](https://github.com/MarvinJWendt/testza#AssertTrue)
  - [AssertZero](https://github.com/MarvinJWendt/testza#AssertZero)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Capture">Capture</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [CaptureStderr](https://github.com/MarvinJWendt/testza#CaptureStderr)
  - [CaptureStdout](https://github.com/MarvinJWendt/testza#CaptureStdout)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Mock-Input-Bool">Mock Input Bool</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [MockInputBoolFull](https://github.com/MarvinJWendt/testza#MockInputBoolFull)
  - [MockInputBoolModify](https://github.com/MarvinJWendt/testza#MockInputBoolModify)
  - [MockInputBoolRunTests](https://github.com/MarvinJWendt/testza#MockInputBoolRunTests)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Mock-Input-String">Mock Input String</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [MockInputStringEmailAddresses](https://github.com/MarvinJWendt/testza#MockInputStringEmailAddresses)
  - [MockInputStringEmpty](https://github.com/MarvinJWendt/testza#MockInputStringEmpty)
  - [MockInputStringFull](https://github.com/MarvinJWendt/testza#MockInputStringFull)
  - [MockInputStringGenerateRandom](https://github.com/MarvinJWendt/testza#MockInputStringGenerateRandom)
  - [MockInputStringHtmlTags](https://github.com/MarvinJWendt/testza#MockInputStringHtmlTags)
  - [MockInputStringLimit](https://github.com/MarvinJWendt/testza#MockInputStringLimit)
  - [MockInputStringLong](https://github.com/MarvinJWendt/testza#MockInputStringLong)
  - [MockInputStringModify](https://github.com/MarvinJWendt/testza#MockInputStringModify)
  - [MockInputStringNumeric](https://github.com/MarvinJWendt/testza#MockInputStringNumeric)
  - [MockInputStringRunTests](https://github.com/MarvinJWendt/testza#MockInputStringRunTests)
  - [MockInputStringUsernames](https://github.com/MarvinJWendt/testza#MockInputStringUsernames)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Mock-Input-Float64">Mock Input Float64</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [MockInputFloat64Full](https://github.com/MarvinJWendt/testza#MockInputFloat64Full)
  - [MockInputFloat64GenerateRandomNegative](https://github.com/MarvinJWendt/testza#MockInputFloat64GenerateRandomNegative)
  - [MockInputFloat64GenerateRandomPositive](https://github.com/MarvinJWendt/testza#MockInputFloat64GenerateRandomPositive)
  - [MockInputFloat64GenerateRandomRange](https://github.com/MarvinJWendt/testza#MockInputFloat64GenerateRandomRange)
  - [MockInputFloat64Modify](https://github.com/MarvinJWendt/testza#MockInputFloat64Modify)
  - [MockInputFloat64RunTests](https://github.com/MarvinJWendt/testza#MockInputFloat64RunTests)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Mock-Input-Int">Mock Input Int</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [MockInputIntFull](https://github.com/MarvinJWendt/testza#MockInputIntFull)
  - [MockInputIntGenerateRandomNegative](https://github.com/MarvinJWendt/testza#MockInputIntGenerateRandomNegative)
  - [MockInputIntGenerateRandomPositive](https://github.com/MarvinJWendt/testza#MockInputIntGenerateRandomPositive)
  - [MockInputIntGenerateRandomRange](https://github.com/MarvinJWendt/testza#MockInputIntGenerateRandomRange)
  - [MockInputIntModify](https://github.com/MarvinJWendt/testza#MockInputIntModify)
  - [MockInputIntRunTests](https://github.com/MarvinJWendt/testza#MockInputIntRunTests)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Mock-Stdin">Mock Stdin</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [MockStdinString](https://github.com/MarvinJWendt/testza#MockStdinString)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Snapshot">Snapshot</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [SnapshotCreate](https://github.com/MarvinJWendt/testza#SnapshotCreate)
  - [SnapshotCreateOrValidate](https://github.com/MarvinJWendt/testza#SnapshotCreateOrValidate)
  - [SnapshotValidate](https://github.com/MarvinJWendt/testza#SnapshotValidate)
</td>

</details>

</tr>
</table>

### Assert

#### AssertCompletesIn

```go
func AssertCompletesIn(t testRunner, duration time.Duration, f func(), msg ...interface{})
```

AssertCompletesIn asserts that a function completes in a given time. Use
this function to test that functions do not take too long to complete.

NOTE: Every system takes a different amount of time to complete a function.
Do not set the duration too low, if you want consistent results.

Example:

    testza.AssertCompletesIn(t, 2 * time.Second, func() {
    	// some code that should take less than 2 seconds...
    }) // => PASS

#### AssertContains

```go
func AssertContains(t testRunner, object, element interface{}, msg ...interface{})
```

AssertContains asserts that a string/list/array/slice/map contains the
specified element.

Example:

    testza.AssertContains(t, []int{1,2,3}, 2)
    testza.AssertContains(t, []string{"Hello", "World"}, "World")
    testza.AssertContains(t, "Hello, World!", "World")

#### AssertEqual

```go
func AssertEqual(t testRunner, expected interface{}, actual interface{}, msg ...interface{})
```

AssertEqual asserts that two objects are equal.

Example:

    testza.AssertEqual(t, "Hello, World!", "Hello, World!")
    testza.AssertEqual(t, true, true)

#### AssertEqualValues

```go
func AssertEqualValues(t testRunner, expected interface{}, actual interface{}, msg ...interface{})
```

AssertEqualValues asserts that two objects have equal values. The order of
the values is also validated.

Example:

    testza.AssertEqualValues(t, []string{"Hello", "World"}, []string{"Hello", "World"})
    testza.AssertEqualValues(t, []int{1,2}, []int{1,2})
    testza.AssertEqualValues(t, []int{1,2}, []int{2,1}) // FAILS (wrong order)

Comparing struct values:

    person1 := Person{
      Name:   "Marvin Wendt",
      Age:    20,
      Gender: "male",
    }

    person2 := Person{
      Name:   "Marvin Wendt",
      Age:    20,
      Gender: "male",
    }

    testza.AssertEqualValues(t, person1, person2)

#### AssertFalse

```go
func AssertFalse(t testRunner, value interface{}, msg ...interface{})
```

AssertFalse asserts that an expression or object resolves to false.

Example:

    testza.AssertFalse(t, false)
    testza.AssertFalse(t, 1 == 2)
    testza.AssertFalse(t, 2 != 2)
    testza.AssertFalse(t, 1 > 5 && 4 < 0)

#### AssertGreater

```go
func AssertGreater(t testRunner, object1, object2 interface{}, msg ...interface{})
```

AssertGreater asserts that the first object is greater than the second.

Example:

    testza.AssertGreater(t, 5, 1)
    testza.AssertGreater(t, 10, -10)

#### AssertImplements

```go
func AssertImplements(t testRunner, interfaceObject, object interface{}, msg ...interface{})
```

AssertImplements asserts that an objects implements an interface.

Example:

    testza.AssertImplements(t, (*YourInterface)(nil), new(YourObject))
    testza.AssertImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass

#### AssertKindOf

```go
func AssertKindOf(t testRunner, expectedKind reflect.Kind, object interface{}, msg ...interface{})
```

AssertKindOf asserts that the object is a type of kind exptectedKind.

#### AssertLess

```go
func AssertLess(t testRunner, object1, object2 interface{}, msg ...interface{})
```

AssertLess asserts that the first object is less than the second.

Example:

    testza.AssertLess(t, 1, 5)
    testza.AssertLess(t, -10, 10)

#### AssertNil

```go
func AssertNil(t testRunner, object interface{}, msg ...interface{})
```

AssertNil asserts that an object is nil.

Example:

    testza.AssertNil(t, nil)

#### AssertNoError

```go
func AssertNoError(t testRunner, err interface{}, msg ...interface{})
```

AssertNoError asserts that an error is nil.

Example:

    err := nil
    testza.AssertNoError(t, err)

#### AssertNotCompletesIn

```go
func AssertNotCompletesIn(t testRunner, duration time.Duration, f func(), msg ...interface{})
```

AssertNotCompletesIn asserts that a function does not complete in a given
time. Use this function to test that functions do not complete to quickly.
For example if your database connection completes in under a millisecond,
there might be something wrong.

NOTE: Every system takes a different amount of time to complete a function.
Do not set the duration too high, if you want consistent results.

Example:

    testza.AssertNotCompletesIn(t, 2 * time.Second, func() {
    	// some code that should take more than 2 seconds...
    	time.Sleep(3 * time.Second)
    }) // => PASS

#### AssertNotContains

```go
func AssertNotContains(t testRunner, object, element interface{}, msg ...interface{})
```

AssertNotContains asserts that a string/list/array/slice/map does not
contain the specified element.

Example:

    testza.AssertNotContains(t, []string{"Hello", "World"}, "Spaceship")
    testza.AssertNotContains(t, "Hello, World!", "Spaceship")

#### AssertNotEqual

```go
func AssertNotEqual(t testRunner, expected interface{}, actual interface{}, msg ...interface{})
```

AssertNotEqual asserts that two objects are not equal.

Example:

    testza.AssertNotEqual(t, true, false)
    testza.AssertNotEqual(t, "Hello", "World")

#### AssertNotEqualValues

```go
func AssertNotEqualValues(t testRunner, expected interface{}, actual interface{}, msg ...interface{})
```

AssertNotEqualValues asserts that two objects do not have equal values.

Example:

    testza.AssertNotEqualValues(t, []int{1,2}, []int{3,4})

Comparing struct values:

    person1 := Person{
      Name:   "Marvin Wendt",
      Age:    20,
      Gender: "male",
    }

    person2 := Person{
      Name:   "Marvin Wendt",
      Age:    20,
      Gender: "female", // <-- CHANGED
    }

    testza.AssertNotEqualValues(t, person1, person2)

#### AssertNotImplements

```go
func AssertNotImplements(t testRunner, interfaceObject, object interface{}, msg ...interface{})
```

AssertNotImplements asserts that an object does not implement an interface.

Example:

    testza.AssertNotImplements(t, (*YourInterface)(nil), new(YourObject))
    testza.AssertNotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.

#### AssertNotKindOf

```go
func AssertNotKindOf(t testRunner, kind reflect.Kind, object interface{}, msg ...interface{})
```

AssertNotKindOf asserts that the object is not a type of kind `kind`.

#### AssertNotNil

```go
func AssertNotNil(t testRunner, object interface{}, msg ...interface{})
```

AssertNotNil asserts that an object is not nil.

Example:

    testza.AssertNotNil(t, true)
    testza.AssertNotNil(t, "Hello, World!")
    testza.AssertNotNil(t, 0)

#### AssertNotNumeric

```go
func AssertNotNumeric(t testRunner, object interface{}, msg ...interface{})
```

AssertNotNumeric checks if the object is not a numeric type. Numeric types
are: Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16,
Uint32, Uint64, Complex64 and Complex128.

Example:

    testza.AssertNotNumeric(t, true)
    testza.AssertNotNumeric(t, "123")

#### AssertNotPanics

```go
func AssertNotPanics(t testRunner, f func(), msg ...interface{})
```

AssertNotPanics asserts that a function does not panic.

Example:

    testza.AssertNotPanics(t, func() {
    	// some code that does not call a panic...
    }) // => PASS

#### AssertNotZero

```go
func AssertNotZero(t testRunner, value interface{}, msg ...interface{})
```

AssertNotZero asserts that the value is not the zero value for it's type.

Example:

    testza.AssertNotZero(t, 1337)
    testza.AssertNotZero(t, true)
    testza.AssertNotZero(t, "Hello, World")

#### AssertNumeric

```go
func AssertNumeric(t testRunner, object interface{}, msg ...interface{})
```

AssertNumeric asserts that the object is a numeric type. Numeric types are:
Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16,
Uint32, Uint64, Complex64 and Complex128.

Example:

    testza.AssertNumeric(t, 123)
    testza.AssertNumeric(t, 1.23)
    testza.AssertNumeric(t, uint(123))

#### AssertPanics

```go
func AssertPanics(t testRunner, f func(), msg ...interface{})
```

AssertPanics asserts that a function panics.

Example:

    testza.AssertPanics(t, func() {
    	// ...
    	panic("some panic")
    }) // => PASS

#### AssertTestFails

```go
func AssertTestFails(t testRunner, test func(t TestingPackageWithFailFunctions), msg ...interface{})
```

AssertTestFails asserts that a unit test fails. A unit test fails if one of
the following methods is called in the test function: Error, Errorf, Fail,
FailNow, Fatal, Fatalf

Example:

    testza.AssertTestFails(t, func(t testza.TestingPackageWithFailFunctions) {
    	testza.AssertTrue(t, false)
    }) // => Pass

    testza.AssertTestFails(t, func(t testza.TestingPackageWithFailFunctions) {
    	// ...
    	t.Fail() // Or any other failing method.
    }) // => Pass

#### AssertTrue

```go
func AssertTrue(t testRunner, value interface{}, msg ...interface{})
```

AssertTrue asserts that an expression or object resolves to true.

Example:

    testza.AssertTrue(t, true)
    testza.AssertTrue(t, 1 == 1)
    testza.AssertTrue(t, 2 != 3)
    testza.AssertTrue(t, 1 > 0 && 4 < 5)

#### AssertZero

```go
func AssertZero(t testRunner, value interface{}, msg ...interface{})
```

AssertZero asserts that the value is the zero value for it's type.

Example:

    testza.AssertZero(t, 0)
    testza.AssertZero(t, false)
    testza.AssertZero(t, "")

### Capture

#### CaptureStderr

```go
func CaptureStderr(capture func(w io.Writer) error) (string, error)
```

CaptureStderr captures everything written to stderr from a specific
function. You can use this method in tests, to validate that your functions
writes a string to the terminal.

#### CaptureStdout

```go
func CaptureStdout(capture func(w io.Writer) error) (string, error)
```

CaptureStdout captures everything written to stdout from a specific
function. You can use this method in tests, to validate that your functions
writes a string to the terminal.

### Mock Input Bool

#### MockInputBoolFull

```go
func MockInputBoolFull() []bool
```

MockInputBoolFull returns true and false in a boolean slice.

#### MockInputBoolModify

```go
func MockInputBoolModify(inputSlice []bool, f func(index int, value bool) bool) (floats []bool)
```

MockInputBoolModify returns a modified version of a test set.

#### MockInputBoolRunTests

```go
func MockInputBoolRunTests(t testRunner, testSet []bool, testFunc func(t *testing.T, index int, f bool))
```

MockInputBoolRunTests runs a test for every value in a testset. You can use
the value as input parameter for your functions, to sanity test against many
different cases. This ensures that your functions have a correct error
handling and enables you to test against hunderts of cases easily.

### Mock Input Float64

#### MockInputFloat64Full

```go
func MockInputFloat64Full() (floats []float64)
```

MockInputFloat64Full returns a combination of every float64 testset and some
random float64s (positive and negative).

#### MockInputFloat64GenerateRandomNegative

```go
func MockInputFloat64GenerateRandomNegative(count int, min float64) (floats []float64)
```

MockInputFloat64GenerateRandomNegative generates random negative integers
with a minimum of min. If the minimum is positive, it will be converted to a
negative number. If it is set to 0, there is no limit.

#### MockInputFloat64GenerateRandomPositive

```go
func MockInputFloat64GenerateRandomPositive(count int, max float64) (floats []float64)
```

MockInputFloat64GenerateRandomPositive generates random positive integers
with a maximum of max. If the maximum is 0, or below, the maximum will be
set to math.MaxInt64.

#### MockInputFloat64GenerateRandomRange

```go
func MockInputFloat64GenerateRandomRange(count int, min, max float64) (floats []float64)
```

MockInputFloat64GenerateRandomRange generates random positive integers with
a maximum of max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### MockInputFloat64Modify

```go
func MockInputFloat64Modify(inputSlice []float64, f func(index int, value float64) float64) (floats []float64)
```

MockInputFloat64Modify returns a modified version of a test set.

#### MockInputFloat64RunTests

```go
func MockInputFloat64RunTests(t testRunner, testSet []float64, testFunc func(t *testing.T, index int, f float64))
```

MockInputFloat64RunTests runs a test for every value in a testset. You can
use the value as input parameter for your functions, to sanity test against
many different cases. This ensures that your functions have a correct error
handling and enables you to test against hunderts of cases easily.

### Mock Input Int

#### MockInputIntFull

```go
func MockInputIntFull() (ints []int)
```

MockInputIntFull returns a combination of every integer testset and some
random integers (positive and negative).

#### MockInputIntGenerateRandomNegative

```go
func MockInputIntGenerateRandomNegative(count, min int) (ints []int)
```

MockInputIntGenerateRandomNegative generates random negative integers with a
minimum of min. If the minimum is 0, or above, the maximum will be set to
math.MinInt64.

#### MockInputIntGenerateRandomPositive

```go
func MockInputIntGenerateRandomPositive(count, max int) (ints []int)
```

MockInputIntGenerateRandomPositive generates random positive integers with a
maximum of max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### MockInputIntGenerateRandomRange

```go
func MockInputIntGenerateRandomRange(count, min, max int) (ints []int)
```

MockInputIntGenerateRandomRange generates random integers with a range of
min to max.

#### MockInputIntModify

```go
func MockInputIntModify(inputSlice []int, f func(index int, value int) int) (ints []int)
```

MockInputIntModify returns a modified version of a test set.

#### MockInputIntRunTests

```go
func MockInputIntRunTests(t testRunner, testSet []int, testFunc func(t *testing.T, index int, i int))
```

MockInputIntRunTests runs a test for every value in a testset. You can use
the value as input parameter for your functions, to sanity test against many
different cases. This ensures that your functions have a correct error
handling and enables you to test against hunderts of cases easily.

### Mock Input String

#### MockInputStringEmailAddresses

```go
func MockInputStringEmailAddresses() []string
```

MockInputStringEmailAddresses returns a test set with valid email addresses.

#### MockInputStringEmpty

```go
func MockInputStringEmpty() []string
```

MockInputStringEmpty returns a test set with a single empty string.

#### MockInputStringFull

```go
func MockInputStringFull() (ret []string)
```

MockInputStringFull contains all string test sets plus ten generated random
strings.

#### MockInputStringGenerateRandom

```go
func MockInputStringGenerateRandom(count, length int) (result []string)
```

MockInputStringGenerateRandom returns random StringsHelper in a test set.

#### MockInputStringHtmlTags

```go
func MockInputStringHtmlTags() []string
```

MockInputStringHtmlTags returns a test set with html tags.

#### MockInputStringLimit

```go
func MockInputStringLimit(testSet []string, max int) []string
```

MockInputStringLimit limits a test set in size.

#### MockInputStringLong

```go
func MockInputStringLong() (testSet []string)
```

MockInputStringLong returns a test set with long random strings. Returns: -
Random string (length: 25) - Random string (length: 50) - Random string
(length: 100) - Random string (length: 1,000) - Random string (length:
100,000)

#### MockInputStringModify

```go
func MockInputStringModify(inputSlice []string, f func(index int, value string) string) (ret []string)
```

MockInputStringModify returns a modified version of a test set.

#### MockInputStringNumeric

```go
func MockInputStringNumeric() []string
```

MockInputStringNumeric returns a test set with strings that are numeric. The
highest number in here is "9223372036854775807", which is equal to the
maxmim int64.

#### MockInputStringRunTests

```go
func MockInputStringRunTests(t testRunner, testSet []string, testFunc func(t *testing.T, index int, str string))
```

MockInputStringRunTests runs a test for every value in a testset. You can
use the value as input parameter for your functions, to sanity test against
many different cases. This ensures that your functions have a correct error
handling and enables you to test against hunderts of cases easily.

#### MockInputStringUsernames

```go
func MockInputStringUsernames() []string
```

MockInputStringUsernames returns a test set with usernames.

### Mock Stdin

#### MockStdinString

```go
func MockStdinString(t testRunner, f func() error) (output string, err error)
```



### Settings

#### SetColorsEnabled

```go
func SetColorsEnabled(enabled bool)
```

SetColorsEnabled controls if testza should print colored output. You should
use this in the init() method of the package, which contains your tests.

Example:

    init() {
      testza.SetColorsEnabled(false) // Disable colored output
      testza.SetColorsEnabled(true)  // Enable colored output
    }

#### SetLineNumbersEnabled

```go
func SetLineNumbersEnabled(enabled bool)
```

SetLineNumbersEnabled controls if line numbers should be printed in failing
tests. You should use this in the init() method of the package, which
contains your tests.

Example:

    init() {
      testza.SetLineNumbersEnabled(false) // Disable line numbers
      testza.SetLineNumbersEnabled(true)  // Enable line numbers
    }

### Snapshot

#### SnapshotCreate

```go
func SnapshotCreate(name string, snapshotObject interface{}) error
```

SnapshotCreate creates a snapshot of an object, which can be validated in
future test runs. Using this function directly will override previous
snapshots with the same name. You most likely want to use
SnapshotCreateOrValidate.

NOTICE: \r\n will be replaced with \n to make the files consistent between
operating systems.

Example:

    testza.SnapshotCreate(t.Name(), objectToBeSnapshotted)

#### SnapshotCreateOrValidate

```go
func SnapshotCreateOrValidate(t testRunner, name string, object interface{}, msg ...interface{}) error
```

SnapshotCreateOrValidate creates a snapshot of an object which can be used
in future test runs. It is good practice to name your snapshots the same as
the test they are created in. You can do that automatically by using
t.Name() as the second parameter, if you are using the inbuilt test system
of Go. If a snapshot already exists, the function will not create a new one,
but validate the exisiting one. To re-create a snapshot, you can delete the
according file in /testdata/snapshots/.

NOTICE: \r\n will be replaced with \n to make the files consistent between
operating systems.

Example:

    testza.SnapshotCreateOrValidate(t, t.Name(), object)
    testza.SnapshotCreateOrValidate(t, t.Name(), object, "Optional Message")

#### SnapshotValidate

```go
func SnapshotValidate(t testRunner, name string, actual interface{}, msg ...interface{}) error
```

SnapshotValidate validates an already exisiting snapshot of an object. You
most likely want to use SnapshotCreateOrValidate.

NOTICE: \r\n will be replaced with \n to make the files consistent between
operating systems.

Example:

    testza.SnapshotValidate(t, t.Name(), objectToBeValidated)
    testza.SnapshotValidate(t, t.Name(), objectToBeValidated, "Optional message")


<!-- docs:end -->

---

> Made with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) and contributors! |
> [MarvinJWendt.com](https://marvinjwendt.com)
