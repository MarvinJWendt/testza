<h1 align="center">testza 🍕</h1>
<p align="center">Testza is like pizza for Go - you could life without it, but why should you?</p>

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
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-847-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://github.com/MarvinJWendt/testza/issues">
<img src="https://img.shields.io/github/issues/MarvinJWendt/testza.svg?style=flat-square" alt="Issues">
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
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

Testza is structured a bit differently than you might be used to in Go, but we think that it makes writing tests as easy and efficient as possible.  
After all, writing tests should be very simple and should not require you to study a whole framework.  
That's why we made testza to integrate perfectly with your IDE.
You don't even have to lookup the documentation, as testza is self-explanatory.

## Getting Started

Testza is very IDE friendly and was made to integrate with your IDE to increase your productivity.  

```go
   ┌ Testza package
   |    ┌ Container for all testza modules
   |    |     ┌ The module you want to use (Press Ctrl+Space to see a list of all modules)
testza.Use.XXXXXXX


// --- Some Examples ---

// - Some assertions -
testza.Use.Assert.True(t, true) // -> Pass
testza.Use.Assert.NoError(t, err) // -> Pass
testza.Use.Assert.Equal(t, object, object) // -> Pass
// ...

// - Testing console output -
// Test the output of your CLI tool easily!
terminalOutput, _ := testza.Use.Capture.Stdout(func(w io.Writer) error {fmt.Println("Hello"); return nil})
testza.Use.Assert.Equal(t, terminalOutput, "Hello\n") // -> Pass

// - Mocking -
// Testing a function that accepts email addresses as a parameter:

// Testset of many different email addresses
emailAddresses := testza.Use.Mock.Strings.EmailAddresses()

// Run a test for every string in the test set
testza.Use.Mock.Strings.RunTests(t, emailAddresses, func(t *testing.T, index int, str string) {
  user, domain, err := internal.ParseEmailAddress(str) // Use your function
  testza.Use.Assert.NoError(t, err) // Assert that your function does not return an error
  testza.Use.Assert.NotZero(t, user) // Assert that the user is returned
  testza.Use.Assert.NotZero(t, domain) // Assert that the domain is returned
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
<td><a href="https://github.com/MarvinJWendt/testza#Assert">Assert</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [CompletesIn](https://github.com/MarvinJWendt/testza#testzaUseAssertCompletesIn)
  - [Contains](https://github.com/MarvinJWendt/testza#testzaUseAssertContains)
  - [Equal](https://github.com/MarvinJWendt/testza#testzaUseAssertEqual)
  - [EqualValues](https://github.com/MarvinJWendt/testza#testzaUseAssertEqualValues)
  - [False](https://github.com/MarvinJWendt/testza#testzaUseAssertFalse)
  - [Greater](https://github.com/MarvinJWendt/testza#testzaUseAssertGreater)
  - [Implements](https://github.com/MarvinJWendt/testza#testzaUseAssertImplements)
  - [KindOf](https://github.com/MarvinJWendt/testza#testzaUseAssertKindOf)
  - [Less](https://github.com/MarvinJWendt/testza#testzaUseAssertLess)
  - [Nil](https://github.com/MarvinJWendt/testza#testzaUseAssertNil)
  - [NoError](https://github.com/MarvinJWendt/testza#testzaUseAssertNoError)
  - [NotCompletesIn](https://github.com/MarvinJWendt/testza#testzaUseAssertNotCompletesIn)
  - [NotContains](https://github.com/MarvinJWendt/testza#testzaUseAssertNotContains)
  - [NotEqual](https://github.com/MarvinJWendt/testza#testzaUseAssertNotEqual)
  - [NotEqualValues](https://github.com/MarvinJWendt/testza#testzaUseAssertNotEqualValues)
  - [NotImplements](https://github.com/MarvinJWendt/testza#testzaUseAssertNotImplements)
  - [NotKindOf](https://github.com/MarvinJWendt/testza#testzaUseAssertNotKindOf)
  - [NotNil](https://github.com/MarvinJWendt/testza#testzaUseAssertNotNil)
  - [NotNumeric](https://github.com/MarvinJWendt/testza#testzaUseAssertNotNumeric)
  - [NotPanic](https://github.com/MarvinJWendt/testza#testzaUseAssertNotPanic)
  - [NotZero](https://github.com/MarvinJWendt/testza#testzaUseAssertNotZero)
  - [Numeric](https://github.com/MarvinJWendt/testza#testzaUseAssertNumeric)
  - [Panic](https://github.com/MarvinJWendt/testza#testzaUseAssertPanic)
  - [True](https://github.com/MarvinJWendt/testza#testzaUseAssertTrue)
  - [Zero](https://github.com/MarvinJWendt/testza#testzaUseAssertZero)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Capture">Capture</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [Stderr](https://github.com/MarvinJWendt/testza#testzaUseCaptureStderr)
  - [Stdout](https://github.com/MarvinJWendt/testza#testzaUseCaptureStdout)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#MockStrings">Mock.Strings</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [EmailAddresses](https://github.com/MarvinJWendt/testza#testzaUseMockStringsEmailAddresses)
  - [Empty](https://github.com/MarvinJWendt/testza#testzaUseMockStringsEmpty)
  - [Full](https://github.com/MarvinJWendt/testza#testzaUseMockStringsFull)
  - [GenerateRandom](https://github.com/MarvinJWendt/testza#testzaUseMockStringsGenerateRandom)
  - [HtmlTags](https://github.com/MarvinJWendt/testza#testzaUseMockStringsHtmlTags)
  - [Limit](https://github.com/MarvinJWendt/testza#testzaUseMockStringsLimit)
  - [Modify](https://github.com/MarvinJWendt/testza#testzaUseMockStringsModify)
  - [Numeric](https://github.com/MarvinJWendt/testza#testzaUseMockStringsNumeric)
  - [RunTests](https://github.com/MarvinJWendt/testza#testzaUseMockStringsRunTests)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#MockBools">Mock.Bools</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [Full](https://github.com/MarvinJWendt/testza#testzaUseMockBoolsFull)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#MockFloats64">Mock.Floats64</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [Full](https://github.com/MarvinJWendt/testza#testzaUseMockFloats64Full)
  - [GenerateRandomNegative](https://github.com/MarvinJWendt/testza#testzaUseMockFloats64GenerateRandomNegative)
  - [GenerateRandomPositive](https://github.com/MarvinJWendt/testza#testzaUseMockFloats64GenerateRandomPositive)
  - [GenerateRandomRange](https://github.com/MarvinJWendt/testza#testzaUseMockFloats64GenerateRandomRange)
  - [Modify](https://github.com/MarvinJWendt/testza#testzaUseMockFloats64Modify)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#MockInts">Mock.Ints</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [Full](https://github.com/MarvinJWendt/testza#testzaUseMockIntsFull)
  - [GenerateRandomNegative](https://github.com/MarvinJWendt/testza#testzaUseMockIntsGenerateRandomNegative)
  - [GenerateRandomPositive](https://github.com/MarvinJWendt/testza#testzaUseMockIntsGenerateRandomPositive)
  - [GenerateRandomRange](https://github.com/MarvinJWendt/testza#testzaUseMockIntsGenerateRandomRange)
  - [Modify](https://github.com/MarvinJWendt/testza#testzaUseMockIntsModify)
</td>

</details>

</tr>
</table>

### Assert

#### testza.Use.Assert.CompletesIn

```go
func (a AssertHelper) CompletesIn(t testingT, duration time.Duration, f func(), msg ...interface{})
```

CompletesIn asserts that a function completes in a given time. Use this
function to test that functions do not take too long to complete.

NOTE: Every system takes a different amount of time to complete a function.
Do not set the duration too low, if you want consistent results.

#### testza.Use.Assert.Contains

```go
func (a AssertHelper) Contains(t testingT, object, element interface{}, msg ...interface{})
```



#### testza.Use.Assert.Equal

```go
func (a AssertHelper) Equal(t testingT, expected interface{}, actual interface{}, msg ...interface{})
```

Equal asserts that two objects are equal.

#### testza.Use.Assert.EqualValues

```go
func (a AssertHelper) EqualValues(t testingT, expected interface{}, actual interface{}, msg ...interface{})
```

EqualValues asserts that two objects have equal values.

#### testza.Use.Assert.False

```go
func (a AssertHelper) False(t testingT, value interface{}, msg ...interface{})
```

False asserts that an expression or object resolves to false.

#### testza.Use.Assert.Greater

```go
func (a AssertHelper) Greater(t testingT, object1, object2 interface{}, msg ...interface{})
```

Greater asserts that the first object is greater than the second.

#### testza.Use.Assert.Implements

```go
func (a AssertHelper) Implements(t testingT, interfaceObject, object interface{}, msg ...interface{})
```

Implements checks if an objects implements an interface.

    testza.Use.Assert.Implements(t, (*YourInterface)(nil), new(YourObject))
    testza.Use.Assert.Implements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass

#### testza.Use.Assert.KindOf

```go
func (a AssertHelper) KindOf(t testingT, expectedKind reflect.Kind, object interface{}, msg ...interface{})
```

KindOf asserts that the object is a type of kind exptectedKind.

#### testza.Use.Assert.Less

```go
func (a AssertHelper) Less(t testingT, object1, object2 interface{}, msg ...interface{})
```

Less asserts that the first object is less than the second.

#### testza.Use.Assert.Nil

```go
func (a AssertHelper) Nil(t testingT, object interface{}, msg ...interface{})
```

Nil asserts that an object is nil.

#### testza.Use.Assert.NoError

```go
func (a AssertHelper) NoError(t testingT, err interface{}, msg ...interface{})
```

NoError asserts that an error is nil.

#### testza.Use.Assert.NotCompletesIn

```go
func (a AssertHelper) NotCompletesIn(t testingT, duration time.Duration, f func(), msg ...interface{})
```

NotCompletesIn asserts that a function does not complete in a given time.
Use this function to test that functions do not complete to quickly. For
example if your database connection completes in under a millisecond, there
might be something wrong.

NOTE: Every system takes a different amount of time to complete a function.
Do not set the duration too high, if you want consistent results.

#### testza.Use.Assert.NotContains

```go
func (a AssertHelper) NotContains(t testingT, object, element interface{}, msg ...interface{})
```



#### testza.Use.Assert.NotEqual

```go
func (a AssertHelper) NotEqual(t testingT, expected interface{}, actual interface{}, msg ...interface{})
```

NotEqual asserts that two objects are not equal.

#### testza.Use.Assert.NotEqualValues

```go
func (a AssertHelper) NotEqualValues(t testingT, expected interface{}, actual interface{}, msg ...interface{})
```

NotEqualValues asserts that two objects do not have equal values.

#### testza.Use.Assert.NotImplements

```go
func (a AssertHelper) NotImplements(t testingT, interfaceObject, object interface{}, msg ...interface{})
```

NotImplements checks if an object does not implement an interface.

    testza.Use.Assert.NotImplements(t, (*YourInterface)(nil), new(YourObject))
    testza.Use.Assert.NotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.

#### testza.Use.Assert.NotKindOf

```go
func (a AssertHelper) NotKindOf(t testingT, kind reflect.Kind, object interface{}, msg ...interface{})
```

NotKindOf asserts that the object is not a type of kind `kind`.

#### testza.Use.Assert.NotNil

```go
func (a AssertHelper) NotNil(t testingT, object interface{}, msg ...interface{})
```

NotNil assertsthat an object is not nil.

#### testza.Use.Assert.NotNumeric

```go
func (a AssertHelper) NotNumeric(t testingT, object interface{}, msg ...interface{})
```

Number checks if the object is not a numeric type. Numeric types are: Int,
Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32,
Uint64, Complex64 and Complex128.

#### testza.Use.Assert.NotPanic

```go
func (a AssertHelper) NotPanic(t testingT, f func(), msg ...interface{})
```

NotPanic asserts that a function does not panic.

#### testza.Use.Assert.NotZero

```go
func (a AssertHelper) NotZero(t testingT, value interface{}, msg ...interface{})
```

NotZero asserts that the value is not the zero value for it's type.

    assert.NotZero(t, 1337)
    assert.NotZero(t, true)
    assert.NotZero(t, "Hello, World")

#### testza.Use.Assert.Numeric

```go
func (a AssertHelper) Numeric(t testingT, object interface{}, msg ...interface{})
```

Numeric asserts that the object is a numeric type. Numeric types are: Int,
Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32,
Uint64, Complex64 and Complex128.

#### testza.Use.Assert.Panic

```go
func (a AssertHelper) Panic(t testingT, f func(), msg ...interface{})
```

Panic asserts that a function panics.

#### testza.Use.Assert.True

```go
func (a AssertHelper) True(t testingT, value interface{}, msg ...interface{})
```

True asserts that an expression or object resolves to true.

#### testza.Use.Assert.Zero

```go
func (a AssertHelper) Zero(t testingT, value interface{}, msg ...interface{})
```

Zero asserts that the value is the zero value for it's type.

    assert.Zero(t, 0)
    assert.Zero(t, false)
    assert.Zero(t, "")

### Mock.Bools

#### testza.Use.Mock.Bools.Full

```go
func (BoolsHelper) Full() []bool
```

Full returns true and false in a boolean slice.

### Capture

#### testza.Use.Capture.Stderr

```go
func (h *CaptureHelper) Stderr(capture func(w io.Writer) error) (string, error)
```

Stderr captures everything written to stderr from a specific function. You
can use this method in tests, to validate that your functions writes a
string to the terminal.

#### testza.Use.Capture.Stdout

```go
func (h *CaptureHelper) Stdout(capture func(w io.Writer) error) (string, error)
```

Stdout captures everything written to stdout from a specific function. You
can use this method in tests, to validate that your functions writes a
string to the terminal.

### Mock.Floats64

#### testza.Use.Mock.Floats64.Full

```go
func (h Floats64Helper) Full() (floats []float64)
```



#### testza.Use.Mock.Floats64.GenerateRandomNegative

```go
func (h Floats64Helper) GenerateRandomNegative(count int, min float64) (floats []float64)
```

GenerateRandomNegative generates random negative integers with a minimum of
min. If the minimum is positive, it will be converted to a negative number.
If it is set to 0, there is no limit.

#### testza.Use.Mock.Floats64.GenerateRandomPositive

```go
func (h Floats64Helper) GenerateRandomPositive(count int, max float64) (floats []float64)
```

GenerateRandomPositive generates random positive integers with a maximum of
max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### testza.Use.Mock.Floats64.GenerateRandomRange

```go
func (h Floats64Helper) GenerateRandomRange(count int, min, max float64) (floats []float64)
```

GenerateRandomRange generates random positive integers with a maximum of
max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### testza.Use.Mock.Floats64.Modify

```go
func (h Floats64Helper) Modify(inputSlice []float64, f func(index int, value float64) float64) (floats []float64)
```

Modify returns a modified version of a test set.

### Mock.Ints

#### testza.Use.Mock.Ints.Full

```go
func (h IntsHelper) Full() (ints []int)
```

Full returns a combination of every integer testset and some random integers
(positive and negative).

#### testza.Use.Mock.Ints.GenerateRandomNegative

```go
func (h IntsHelper) GenerateRandomNegative(count, min int) (ints []int)
```

GenerateRandomNegative generates random negative integers with a minimum of
min. If the minimum is 0, or above, the maximum will be set to
math.MinInt64.

#### testza.Use.Mock.Ints.GenerateRandomPositive

```go
func (h IntsHelper) GenerateRandomPositive(count, max int) (ints []int)
```

GenerateRandomPositive generates random positive integers with a maximum of
max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### testza.Use.Mock.Ints.GenerateRandomRange

```go
func (h IntsHelper) GenerateRandomRange(count, min, max int) (ints []int)
```

GenerateRandomRange generates random integers with a range of min to max.

#### testza.Use.Mock.Ints.Modify

```go
func (h IntsHelper) Modify(inputSlice []int, f func(index int, value int) int) (ints []int)
```

Modify returns a modified version of a test set.

### Mock.Strings

#### testza.Use.Mock.Strings.EmailAddresses

```go
func (s StringsHelper) EmailAddresses() []string
```

EmailAddresses returns a test set with valid email addresses.

#### testza.Use.Mock.Strings.Empty

```go
func (s StringsHelper) Empty() []string
```

Empty returns a test set with a single empty string.

#### testza.Use.Mock.Strings.Full

```go
func (s StringsHelper) Full() (ret []string)
```

Full contains all string test sets plus ten generated random strings.

#### testza.Use.Mock.Strings.GenerateRandom

```go
func (s StringsHelper) GenerateRandom(count, length int) (result []string)
```

GenerateRandom returns random StringsHelper in a test set.

#### testza.Use.Mock.Strings.HtmlTags

```go
func (s StringsHelper) HtmlTags() []string
```

HtmlTags returns a test set with html tags.

#### testza.Use.Mock.Strings.Limit

```go
func (s StringsHelper) Limit(testSet []string, max int) []string
```

Limit limits a test set in size.

#### testza.Use.Mock.Strings.Modify

```go
func (s StringsHelper) Modify(inputSlice []string, f func(index int, value string) string) (ret []string)
```

Modify returns a modified version of a test set.

#### testza.Use.Mock.Strings.Numeric

```go
func (s StringsHelper) Numeric() []string
```

Numeric returns a test set with strings that are numeric. The highest number
in here is "9223372036854775807", which is equal to the maxmim int64.

#### testza.Use.Mock.Strings.RunTests

```go
func (s StringsHelper) RunTests(t testingT, testSet []string, testFunc func(t *testing.T, index int, str string))
```

RunTests runs tests with a specific test set.


<!-- docs:end -->

---

> Made with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) and contributors! |
> [MarvinJWendt.com](https://marvinjwendt.com)
