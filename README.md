<h1 align="center">testza 🍕</h1>
<p align="center">Testza is like pizza for Go - you could live without it, but why should you?</p>

<p align="center">

<a href="https://github.com/MarvinJWendt/testza/releases">
<img src="https://img.shields.io/github/v/release/MarvinJWendt/testza?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/MarvinJWendt/testza" target="_blank">
<img src="https://img.shields.io/github/actions/workflow/status/MarvinJWendt/testza/go.yml?label=tests&style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/MarvinJWendt/testza" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/MarvinJWendt/testza?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/MarvinJWendt/testza">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-2956-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>
  
<a href="https://pkg.go.dev/github.com/MarvinJWendt/testza" target="_blank">
<img src="https://pkg.go.dev/badge/github.com/MarvinJWendt/testza.svg" alt="Go Reference">
</a>

<a href="https://goreportcard.com/report/github.com/MarvinJWendt/testza" target="_blank">
<img src="https://goreportcard.com/badge/github.com/MarvinJWendt/testza?style=flat-square" alt="Go report">
</a>  

</p>


---

<p align="center">
<strong><a href="https://github.com/MarvinJWendt/testza#-installation">Get The Module</a></strong>
|
<strong><a href="https://github.com/MarvinJWendt/testza#-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<img align="right" height="400" alt="Screenshot of an example test message" src="https://user-images.githubusercontent.com/31022056/161153895-e772bc61-b751-407f-b526-8f6a66d8f8d5.png" />

<br/>

<p align="center">
<a href="https://discord.gg/vE2dNkfAmF">
<img width="300" src="https://user-images.githubusercontent.com/31022056/158916278-4504b838-7ecb-4ab9-a900-7dc002aade78.png" alt="Join us on Discord!" />
<br/>
<b>Join us on Discord for support, discussions, updates and general talk!</b>
</a>
</p>

## 📦 Installation

```console
# Execute this command inside your project
go get github.com/MarvinJWendt/testza
```
<br/>
<br/>

## 📝 Description

Testza is a full-featured testing framework for Go.
It integrates with the default test runner, so you can use it with the standard `go test` tool.
Testza contains easy to use methods, like assertions, output capturing, fuzzing, and much more.

The main goal of testza is to provide an easy and fun experience writing tests and providing a nice, user-friendly output.
Even developers who never used testza, will get into it quickly.

## ⭐ Features

| Feature            | Description                                                                                                                                          |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| Assertions         | Assertions allow you to quickly check objects for expected values.                                                                                   |
| Fuzzing            | Fuzzing allows you to check functions against sets of generated input parameters.<br/>A couple lines of test code can run thousands of sanity tests. |
| Output Capture     | Capture and validate output written to the terminal.<br/>Perfect for CLI tools.                                                                      |
| Snapshots          | Snapshot objects between test runs, to ensure a consistent behaviour.                                                                                |
| Clean Output       | Clean and colorful output provides you the needed details in an easy-to-understand format.                                                           |
| System Information | Testza prints information about the system on startup.<br/> You can quickly figure out what's wrong, when a user submits an issue.                   |
| Well Documented    | Every function of testza is well documented and contains an example to make usage super easy.                                                        |
| Customizable       | Testza features customizable settings, if you want to change something.                                                                              |
| Test flags         | You can configure testza via flags too!<br/>That makes it super simple to change test runs, or output, without touching code!                        |

## 🚀 Getting Started

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

// - Fuzzing -
// Testing a function that accepts email addresses as a parameter:

// Testset of many different email addresses
emailAddresses := testza.FuzzStringEmailAddresses()

// Run a test for every string in the test set
testza.FuzzStringRunTests(t, emailAddresses, func(t *testing.T, index int, str string) {
  user, domain, err := internal.ParseEmailAddress(str) // Use your function
  testza.AssertNoError(t, err) // Assert that your function does not return an error
  testza.AssertNotZero(t, user) // Assert that the user is returned
  testza.AssertNotZero(t, domain) // Assert that the domain is returned
})

// And that's just a few examples of what you can do with Testza!
```

## 📚 Documentation

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
  - [SetDiffContextLines](https://github.com/MarvinJWendt/testza#SetDiffContextLines)
  - [SetLineNumbersEnabled](https://github.com/MarvinJWendt/testza#SetLineNumbersEnabled)
  - [SetRandomSeed](https://github.com/MarvinJWendt/testza#SetRandomSeed)
  - [SetShowStartupMessage](https://github.com/MarvinJWendt/testza#SetShowStartupMessage)
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
  - [AssertDecreasing](https://github.com/MarvinJWendt/testza#AssertDecreasing)
  - [AssertDirEmpty](https://github.com/MarvinJWendt/testza#AssertDirEmpty)
  - [AssertDirExists](https://github.com/MarvinJWendt/testza#AssertDirExists)
  - [AssertDirNotEmpty](https://github.com/MarvinJWendt/testza#AssertDirNotEmpty)
  - [AssertEqual](https://github.com/MarvinJWendt/testza#AssertEqual)
  - [AssertEqualValues](https://github.com/MarvinJWendt/testza#AssertEqualValues)
  - [AssertErrorIs](https://github.com/MarvinJWendt/testza#AssertErrorIs)
  - [AssertFalse](https://github.com/MarvinJWendt/testza#AssertFalse)
  - [AssertFileExists](https://github.com/MarvinJWendt/testza#AssertFileExists)
  - [AssertGreater](https://github.com/MarvinJWendt/testza#AssertGreater)
  - [AssertGreaterOrEqual](https://github.com/MarvinJWendt/testza#AssertGreaterOrEqual)
  - [AssertImplements](https://github.com/MarvinJWendt/testza#AssertImplements)
  - [AssertInRange](https://github.com/MarvinJWendt/testza#AssertInRange)
  - [AssertIncreasing](https://github.com/MarvinJWendt/testza#AssertIncreasing)
  - [AssertKindOf](https://github.com/MarvinJWendt/testza#AssertKindOf)
  - [AssertLen](https://github.com/MarvinJWendt/testza#AssertLen)
  - [AssertLess](https://github.com/MarvinJWendt/testza#AssertLess)
  - [AssertLessOrEqual](https://github.com/MarvinJWendt/testza#AssertLessOrEqual)
  - [AssertNil](https://github.com/MarvinJWendt/testza#AssertNil)
  - [AssertNoDirExists](https://github.com/MarvinJWendt/testza#AssertNoDirExists)
  - [AssertNoError](https://github.com/MarvinJWendt/testza#AssertNoError)
  - [AssertNoFileExists](https://github.com/MarvinJWendt/testza#AssertNoFileExists)
  - [AssertNoSubset](https://github.com/MarvinJWendt/testza#AssertNoSubset)
  - [AssertNotCompletesIn](https://github.com/MarvinJWendt/testza#AssertNotCompletesIn)
  - [AssertNotContains](https://github.com/MarvinJWendt/testza#AssertNotContains)
  - [AssertNotEqual](https://github.com/MarvinJWendt/testza#AssertNotEqual)
  - [AssertNotEqualValues](https://github.com/MarvinJWendt/testza#AssertNotEqualValues)
  - [AssertNotErrorIs](https://github.com/MarvinJWendt/testza#AssertNotErrorIs)
  - [AssertNotImplements](https://github.com/MarvinJWendt/testza#AssertNotImplements)
  - [AssertNotInRange](https://github.com/MarvinJWendt/testza#AssertNotInRange)
  - [AssertNotKindOf](https://github.com/MarvinJWendt/testza#AssertNotKindOf)
  - [AssertNotNil](https://github.com/MarvinJWendt/testza#AssertNotNil)
  - [AssertNotNumeric](https://github.com/MarvinJWendt/testza#AssertNotNumeric)
  - [AssertNotPanics](https://github.com/MarvinJWendt/testza#AssertNotPanics)
  - [AssertNotRegexp](https://github.com/MarvinJWendt/testza#AssertNotRegexp)
  - [AssertNotSameElements](https://github.com/MarvinJWendt/testza#AssertNotSameElements)
  - [AssertNotUnique](https://github.com/MarvinJWendt/testza#AssertNotUnique)
  - [AssertNotZero](https://github.com/MarvinJWendt/testza#AssertNotZero)
  - [AssertNumeric](https://github.com/MarvinJWendt/testza#AssertNumeric)
  - [AssertPanics](https://github.com/MarvinJWendt/testza#AssertPanics)
  - [AssertRegexp](https://github.com/MarvinJWendt/testza#AssertRegexp)
  - [AssertSameElements](https://github.com/MarvinJWendt/testza#AssertSameElements)
  - [AssertSubset](https://github.com/MarvinJWendt/testza#AssertSubset)
  - [AssertTestFails](https://github.com/MarvinJWendt/testza#AssertTestFails)
  - [AssertTrue](https://github.com/MarvinJWendt/testza#AssertTrue)
  - [AssertUnique](https://github.com/MarvinJWendt/testza#AssertUnique)
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
  - [CaptureStdoutAndStderr](https://github.com/MarvinJWendt/testza#CaptureStdoutAndStderr)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Fuzz-Utils">Fuzz Utils</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzUtilDistinctSet](https://github.com/MarvinJWendt/testza#FuzzUtilDistinctSet)
  - [FuzzUtilLimitSet](https://github.com/MarvinJWendt/testza#FuzzUtilLimitSet)
  - [FuzzUtilMergeSets](https://github.com/MarvinJWendt/testza#FuzzUtilMergeSets)
  - [FuzzUtilModifySet](https://github.com/MarvinJWendt/testza#FuzzUtilModifySet)
  - [FuzzUtilRunTests](https://github.com/MarvinJWendt/testza#FuzzUtilRunTests)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Fuzz-Booleans">Fuzz Booleans</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzBoolFull](https://github.com/MarvinJWendt/testza#FuzzBoolFull)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Fuzz-Strings">Fuzz Strings</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzStringEmailAddresses](https://github.com/MarvinJWendt/testza#FuzzStringEmailAddresses)
  - [FuzzStringEmpty](https://github.com/MarvinJWendt/testza#FuzzStringEmpty)
  - [FuzzStringFull](https://github.com/MarvinJWendt/testza#FuzzStringFull)
  - [FuzzStringGenerateRandom](https://github.com/MarvinJWendt/testza#FuzzStringGenerateRandom)
  - [FuzzStringHtmlTags](https://github.com/MarvinJWendt/testza#FuzzStringHtmlTags)
  - [FuzzStringLong](https://github.com/MarvinJWendt/testza#FuzzStringLong)
  - [FuzzStringNumeric](https://github.com/MarvinJWendt/testza#FuzzStringNumeric)
  - [FuzzStringUsernames](https://github.com/MarvinJWendt/testza#FuzzStringUsernames)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Fuzz-Float64s">Fuzz Float64s</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzFloat64Full](https://github.com/MarvinJWendt/testza#FuzzFloat64Full)
  - [FuzzFloat64GenerateRandomNegative](https://github.com/MarvinJWendt/testza#FuzzFloat64GenerateRandomNegative)
  - [FuzzFloat64GenerateRandomPositive](https://github.com/MarvinJWendt/testza#FuzzFloat64GenerateRandomPositive)
  - [FuzzFloat64GenerateRandomRange](https://github.com/MarvinJWendt/testza#FuzzFloat64GenerateRandomRange)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/MarvinJWendt/testza#Fuzz-Integers">Fuzz Integers</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzIntFull](https://github.com/MarvinJWendt/testza#FuzzIntFull)
  - [FuzzIntGenerateRandomNegative](https://github.com/MarvinJWendt/testza#FuzzIntGenerateRandomNegative)
  - [FuzzIntGenerateRandomPositive](https://github.com/MarvinJWendt/testza#FuzzIntGenerateRandomPositive)
  - [FuzzIntGenerateRandomRange](https://github.com/MarvinJWendt/testza#FuzzIntGenerateRandomRange)
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
func AssertCompletesIn(t testRunner, duration time.Duration, f func(), msg ...any)
```

AssertCompletesIn asserts that a function completes in a given time.
Use this function to test that functions do not take too long to complete.

NOTE: Every system takes a different amount of time to complete a function.
Do not set the duration too low, if you want consistent results.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertCompletesIn(t, 2 * time.Second, func() {
    	// some code that should take less than 2 seconds...
    }) // => PASS

#### AssertContains

```go
func AssertContains(t testRunner, object, element any, msg ...any)
```

AssertContains asserts that a string/list/array/slice/map contains the
specified element.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertContains(t, []int{1,2,3}, 2)
    testza.AssertContains(t, []string{"Hello", "World"}, "World")
    testza.AssertContains(t, "Hello, World!", "World")

#### AssertDecreasing

```go
func AssertDecreasing(t testRunner, object any, msg ...any)
```

AssertDecreasing asserts that the values in a slice are decreasing. the test
fails if the values are not in a slice or if the values are not comparable.

Valid input kinds are: []int, []int8, []int16, []int32, []int64, []uint,
[]uint8, []uint16, []uint32, []uint64, []float32, []float64.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertDecreasing(t, []int{1000, 137, 2, 1})
    testza.AssertDecreasing(t, []float32{13.5, 7, 0.1, -10.3})

#### AssertDirEmpty

```go
func AssertDirEmpty(t testRunner, dir string, msg ...any)
```

AssertDirEmpty asserts that a directory is empty. The test will pass when
the directory is empty or does not exist.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertDirEmpty(t, "FolderName")

#### AssertDirExists

```go
func AssertDirExists(t testRunner, dir string, msg ...any)
```

AssertDirExists asserts that a directory exists. The test will pass when the
directory exists, and it's visible to the current user. The test will fail,
if the path points to a file.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertDirExists(t, "FolderName")

#### AssertDirNotEmpty

```go
func AssertDirNotEmpty(t testRunner, dir string, msg ...any)
```

AssertDirNotEmpty asserts that a directory is not empty The test will pass
when the directory is not empty and will fail if the directory does not
exist.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertDirNotEmpty(t, "FolderName")

#### AssertEqual

```go
func AssertEqual(t testRunner, expected any, actual any, msg ...any)
```

AssertEqual asserts that two objects are equal.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertEqual(t, "Hello, World!", "Hello, World!")
    testza.AssertEqual(t, true, true)

#### AssertEqualValues

```go
func AssertEqualValues(t testRunner, expected any, actual any, msg ...any)
```

AssertEqualValues asserts that two objects have equal values. The order of
the values is also validated.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

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

#### AssertErrorIs

```go
func AssertErrorIs(t testRunner, err, target error, msg ...any)
```

AssertErrorIs asserts that target is inside the error chain of err.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    var testErr = errors.New("hello world")
    var testErrWrapped = fmt.Errorf("test err: %w", testErr)
    testza.AssertErrorIs(t, testErrWrapped ,testErr)

#### AssertFalse

```go
func AssertFalse(t testRunner, value any, msg ...any)
```

AssertFalse asserts that an expression or object resolves to false.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertFalse(t, false)
    testza.AssertFalse(t, 1 == 2)
    testza.AssertFalse(t, 2 != 2)
    testza.AssertFalse(t, 1 > 5 && 4 < 0)

#### AssertFileExists

```go
func AssertFileExists(t testRunner, file string, msg ...any)
```

AssertFileExists asserts that a file exists.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertFileExists(t, "./test.txt")
    testza.AssertFileExists(t, "./config.yaml", "the config file is missing")

#### AssertGreater

```go
func AssertGreater(t testRunner, object1, object2 any, msg ...any)
```

AssertGreater asserts that the first object is greater than the second.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertGreater(t, 5, 1)
    testza.AssertGreater(t, 10, -10)

#### AssertGreaterOrEqual

```go
func AssertGreaterOrEqual(t testRunner, object1, object2 interface{}, msg ...interface{})
```

AssertGreaterOrEqual asserts that the first object is greater than or equal
to the second.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertGreaterOrEqual(t, 5, 1)
    testza.AssertGreaterOrEqual(t, 10, -10)

testza.AssertGreaterOrEqual(t, 10, 10)

#### AssertImplements

```go
func AssertImplements(t testRunner, interfaceObject, object any, msg ...any)
```

AssertImplements asserts that an objects implements an interface.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertImplements(t, (*YourInterface)(nil), new(YourObject))
    testza.AssertImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass

#### AssertInRange

```go
func AssertInRange[T number](t testRunner, value T, min T, max T, msg ...any)
```

AssertInRange asserts that the value is in the range.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertInRange(t, 5, 1, 10)

#### AssertIncreasing

```go
func AssertIncreasing(t testRunner, object any, msg ...any)
```

AssertIncreasing asserts that the values in a slice are increasing. the test
fails if the values are not in a slice or if the values are not comparable.

Valid input kinds are: []int, []int8, []int16, []int32, []int64, []uint,
[]uint8, []uint16, []uint32, []uint64, []float32, []float64.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertIncreasing(t, []int{1, 2, 137, 1000})
    testza.AssertIncreasing(t, []float32{-10.3, 0.1, 7, 13.5})

#### AssertKindOf

```go
func AssertKindOf(t testRunner, expectedKind reflect.Kind, object any, msg ...any)
```

AssertKindOf asserts that the object is a type of kind exptectedKind.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertKindOf(t, reflect.Slice, []int{1,2,3})
    testza.AssertKindOf(t, reflect.Slice, []string{"Hello", "World"})
    testza.AssertKindOf(t, reflect.Int, 1337)
    testza.AssertKindOf(t, reflect.Bool, true)
    testza.AssertKindOf(t, reflect.Map, map[string]bool{})

#### AssertLen

```go
func AssertLen(t testRunner, object any, length int, msg ...any)
```

AssertLen asserts that the length of an object is equal to the given length.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertLen(t, "abc", 3)
    testza.AssertLen(t, "Assert", 6)
    testza.AssertLen(t, []int{1, 2, 1337, 25}, 4)
    testza.AssertLen(t, map[string]int{"asd": 1, "test": 1337}, 2)

#### AssertLess

```go
func AssertLess(t testRunner, object1, object2 any, msg ...any)
```

AssertLess asserts that the first object is less than the second.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertLess(t, 1, 5)
    testza.AssertLess(t, -10, 10)

#### AssertLessOrEqual

```go
func AssertLessOrEqual(t testRunner, object1, object2 interface{}, msg ...interface{})
```

AssertLessOrEqual asserts that the first object is less than or equal to the
second.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertLessOrEqual(t, 1, 5)
    testza.AssertLessOrEqual(t, -10, 10)
    testza.AssertLessOrEqual(t, 1, 1)

#### AssertNil

```go
func AssertNil(t testRunner, object any, msg ...any)
```

AssertNil asserts that an object is nil.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNil(t, nil)

#### AssertNoDirExists

```go
func AssertNoDirExists(t testRunner, dir string, msg ...any)
```

AssertNoDirExists asserts that a directory does not exists. The test will
pass, if the path points to a file, as a directory with the same name,
cannot exist.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNoDirExists(t, "FolderName")

#### AssertNoError

```go
func AssertNoError(t testRunner, err error, msg ...any)
```

AssertNoError asserts that an error is nil.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    err := nil
    testza.AssertNoError(t, err)

#### AssertNoFileExists

```go
func AssertNoFileExists(t testRunner, file string, msg ...any)
```



#### AssertNoSubset

```go
func AssertNoSubset(t testRunner, list any, subset any, msg ...any)
```

AssertNoSubset asserts that the second parameter is not a subset of the
list. The order is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNoSubset(t, []int{1, 2, 3}, []int{1, 7})
    testza.AssertNoSubset(t, []string{"Hello", "World", "Test"}, []string{"Test", "John"})

#### AssertNotCompletesIn

```go
func AssertNotCompletesIn(t testRunner, duration time.Duration, f func(), msg ...any)
```

AssertNotCompletesIn asserts that a function does not complete in a given
time. Use this function to test that functions do not complete to quickly.
For example if your database connection completes in under a millisecond,
there might be something wrong.

NOTE: Every system takes a different amount of time to complete a function.
Do not set the duration too high, if you want consistent results.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotCompletesIn(t, 2 * time.Second, func() {
    	// some code that should take more than 2 seconds...
    	time.Sleep(3 * time.Second)
    }) // => PASS

#### AssertNotContains

```go
func AssertNotContains(t testRunner, object, element any, msg ...any)
```

AssertNotContains asserts that a string/list/array/slice/map does not
contain the specified element.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotContains(t, []string{"Hello", "World"}, "Spaceship")
    testza.AssertNotContains(t, "Hello, World!", "Spaceship")

#### AssertNotEqual

```go
func AssertNotEqual(t testRunner, expected any, actual any, msg ...any)
```

AssertNotEqual asserts that two objects are not equal.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotEqual(t, true, false)
    testza.AssertNotEqual(t, "Hello", "World")

#### AssertNotEqualValues

```go
func AssertNotEqualValues(t testRunner, expected any, actual any, msg ...any)
```

AssertNotEqualValues asserts that two objects do not have equal values.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

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

#### AssertNotErrorIs

```go
func AssertNotErrorIs(t testRunner, err, target error, msg ...any)
```

AssertNotErrorIs

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    var testErr = errors.New("hello world")
    var test2Err = errors.New("hello world 2")
    var testErrWrapped = fmt.Errorf("test err: %w", testErr)
    testza.AssertNotErrorIs(t, testErrWrapped, test2Err)

#### AssertNotImplements

```go
func AssertNotImplements(t testRunner, interfaceObject, object any, msg ...any)
```

AssertNotImplements asserts that an object does not implement an interface.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotImplements(t, (*YourInterface)(nil), new(YourObject))
    testza.AssertNotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.

#### AssertNotInRange

```go
func AssertNotInRange[T number](t testRunner, value T, min T, max T, msg ...any)
```

AssertNotInRange asserts that the value is not in the range.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotInRange(t, 5, 1, 10)

#### AssertNotKindOf

```go
func AssertNotKindOf(t testRunner, kind reflect.Kind, object any, msg ...any)
```

AssertNotKindOf asserts that the object is not a type of kind `kind`.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotKindOf(t, reflect.Slice, "Hello, World")
    testza.AssertNotKindOf(t, reflect.Slice, true)
    testza.AssertNotKindOf(t, reflect.Int, 13.37)
    testza.AssertNotKindOf(t, reflect.Bool, map[string]bool{})
    testza.AssertNotKindOf(t, reflect.Map, false)

#### AssertNotNil

```go
func AssertNotNil(t testRunner, object any, msg ...any)
```

AssertNotNil asserts that an object is not nil.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotNil(t, true)
    testza.AssertNotNil(t, "Hello, World!")
    testza.AssertNotNil(t, 0)

#### AssertNotNumeric

```go
func AssertNotNumeric(t testRunner, object any, msg ...any)
```

AssertNotNumeric checks if the object is not a numeric type. Numeric types
are: Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16,
Uint32, Uint64, Complex64 and Complex128.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotNumeric(t, true)
    testza.AssertNotNumeric(t, "123")

#### AssertNotPanics

```go
func AssertNotPanics(t testRunner, f func(), msg ...any)
```

AssertNotPanics asserts that a function does not panic.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotPanics(t, func() {
    	// some code that does not call a panic...
    }) // => PASS

#### AssertNotRegexp

```go
func AssertNotRegexp(t testRunner, regex any, txt any, msg ...any)
```

AssertNotRegexp asserts that a string does not match a given regexp.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotRegexp(t, "ab.*", "Hello, World!")

#### AssertNotSameElements

```go
func AssertNotSameElements(t testRunner, expected any, actual any, msg ...any)
```

AssertNotSameElements asserts that two slices contains same elements
(including pointers). The order is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

     testza.AssertNotSameElements(t, []string{"Hello", "World"}, []string{"Hello", "World", "World"})
     testza.AssertNotSameElements(t, []int{1,2}, []int{1,2,3})

     type A struct {
    	  a string
     }
     testza.AssertNotSameElements(t, []*A{{a: "A"}, {a: "B"}, {a: "C"}}, []*A{{a: "A"}, {a: "B"}, {a: "C"}, {a: "D"}})

#### AssertNotUnique

```go
func AssertNotUnique[elementType comparable](t testRunner, list []elementType, msg ...any)
```

AssertNotUnique asserts that the elements in a list are not unique.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotUnique(t, []int{1, 2, 3, 3})

#### AssertNotZero

```go
func AssertNotZero(t testRunner, value any, msg ...any)
```

AssertNotZero asserts that the value is not the zero value for it's type.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNotZero(t, 1337)
    testza.AssertNotZero(t, true)
    testza.AssertNotZero(t, "Hello, World")

#### AssertNumeric

```go
func AssertNumeric(t testRunner, object any, msg ...any)
```

AssertNumeric asserts that the object is a numeric type. Numeric types are:
Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16,
Uint32, Uint64, Complex64 and Complex128.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertNumeric(t, 123)
    testza.AssertNumeric(t, 1.23)
    testza.AssertNumeric(t, uint(123))

#### AssertPanics

```go
func AssertPanics(t testRunner, f func(), msg ...any)
```

AssertPanics asserts that a function panics.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertPanics(t, func() {
    	// ...
    	panic("some panic")
    }) // => PASS

#### AssertRegexp

```go
func AssertRegexp(t testRunner, regex any, txt any, msg ...any)
```

AssertRegexp asserts that a string matches a given regexp.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertRegexp(t, "^a.*c$", "abc")

#### AssertSameElements

```go
func AssertSameElements(t testRunner, expected any, actual any, msg ...any)
```

AssertSameElements asserts that two slices contains same elements (including
pointers). The order is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

     testza.AssertSameElements(t, []string{"Hello", "World"}, []string{"Hello", "World"})
     testza.AssertSameElements(t, []int{1,2,3}, []int{1,2,3})
     testza.AssertSameElements(t, []int{1,2}, []int{2,1})

     type A struct {
    	  a string
     }
     testza.AssertSameElements(t, []*A{{a: "A"}, {a: "B"}, {a: "C"}}, []*A{{a: "A"}, {a: "B"}, {a: "C"}})

#### AssertSubset

```go
func AssertSubset(t testRunner, list any, subset any, msg ...any)
```

AssertSubset asserts that the second parameter is a subset of the list.
The order is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertSubset(t, []int{1, 2, 3}, []int{1, 2})
    testza.AssertSubset(t, []string{"Hello", "World", "Test"}, []string{"Test", "World"})

#### AssertTestFails

```go
func AssertTestFails(t testRunner, test func(t TestingPackageWithFailFunctions), msg ...any)
```

AssertTestFails asserts that a unit test fails. A unit test fails if one of
the following methods is called in the test function: Error, Errorf, Fail,
FailNow, Fatal, Fatalf

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

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
func AssertTrue(t testRunner, value any, msg ...any)
```

AssertTrue asserts that an expression or object resolves to true.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertTrue(t, true)
    testza.AssertTrue(t, 1 == 1)
    testza.AssertTrue(t, 2 != 3)
    testza.AssertTrue(t, 1 > 0 && 4 < 5)

#### AssertUnique

```go
func AssertUnique[elementType comparable](t testRunner, list []elementType, msg ...any)
```

AssertUnique asserts that the list contains only unique elements. The order
is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    testza.AssertUnique(t, []int{1, 2, 3})
    testza.AssertUnique(t, []string{"Hello", "World", "!"})

#### AssertZero

```go
func AssertZero(t testRunner, value any, msg ...any)
```

AssertZero asserts that the value is the zero value for it's type.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

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

Example:

    stderr, err := testza.CaptureStderr(func(w io.Writer) error {
    	_, err := fmt.Fprint(os.Stderr, "Hello, World!")
    	testza.AssertNoError(t, err)
    	return nil
    })

    testza.AssertNoError(t, err)
    testza.AssertEqual(t, "Hello, World!", stderr)

#### CaptureStdout

```go
func CaptureStdout(capture func(w io.Writer) error) (string, error)
```

CaptureStdout captures everything written to stdout from a specific
function. You can use this method in tests, to validate that your functions
writes a string to the terminal.

Example:

    stdout, err := testza.CaptureStdout(func(w io.Writer) error {
    	fmt.Println("Hello, World!")
    	return nil
    })

    testza.AssertNoError(t, err)
    testza.AssertEqual(t, "Hello, World!", stdout)

#### CaptureStdoutAndStderr

```go
func CaptureStdoutAndStderr(capture func(stdoutWriter, stderrWriter io.Writer) error) (stdout, stderr string, err error)
```

CaptureStdoutAndStderr captures everything written to stdout and stderr from
a specific function. You can use this method in tests, to validate that your
functions writes a string to the terminal.

Example:

    stdout, stderr, err := testza.CaptureStdoutAndStderr(func(stdoutWriter, stderrWriter io.Writer) error {
    	fmt.Fprint(os.Stdout, "Hello")
    	fmt.Fprint(os.Stderr, "World")
    	return nil
    })

    testza.AssertNoError(t, err)
    testza.AssertEqual(t, "Hello", stdout)
    testza.AssertEqual(t, "World", stderr)

### Fuzz Booleans

#### FuzzBoolFull

```go
func FuzzBoolFull() []bool
```

FuzzBoolFull returns true and false in a boolean slice.

### Fuzz Float64s

#### FuzzFloat64Full

```go
func FuzzFloat64Full() (floats []float64)
```

FuzzFloat64Full returns a combination of every float64 testset and some
random float64s (positive and negative).

#### FuzzFloat64GenerateRandomNegative

```go
func FuzzFloat64GenerateRandomNegative(count int, min float64) (floats []float64)
```

FuzzFloat64GenerateRandomNegative generates random negative integers with
a minimum of min. If the minimum is positive, it will be converted to a
negative number. If it is set to 0, there is no limit.

#### FuzzFloat64GenerateRandomPositive

```go
func FuzzFloat64GenerateRandomPositive(count int, max float64) (floats []float64)
```

FuzzFloat64GenerateRandomPositive generates random positive integers with
a maximum of max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### FuzzFloat64GenerateRandomRange

```go
func FuzzFloat64GenerateRandomRange(count int, min, max float64) (floats []float64)
```

FuzzFloat64GenerateRandomRange generates random positive integers with a
maximum of max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

### Fuzz Integers

#### FuzzIntFull

```go
func FuzzIntFull() (ints []int)
```

FuzzIntFull returns a combination of every integer testset and some random
integers (positive and negative).

#### FuzzIntGenerateRandomNegative

```go
func FuzzIntGenerateRandomNegative(count, min int) (ints []int)
```

FuzzIntGenerateRandomNegative generates random negative integers with a
minimum of min. If the minimum is 0, or above, the maximum will be set to
math.MinInt64.

#### FuzzIntGenerateRandomPositive

```go
func FuzzIntGenerateRandomPositive(count, max int) (ints []int)
```

FuzzIntGenerateRandomPositive generates random positive integers with a
maximum of max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### FuzzIntGenerateRandomRange

```go
func FuzzIntGenerateRandomRange(count, min, max int) (ints []int)
```

FuzzIntGenerateRandomRange generates random integers with a range of min to
max.

### Fuzz Strings

#### FuzzStringEmailAddresses

```go
func FuzzStringEmailAddresses() []string
```

FuzzStringEmailAddresses returns a test set with valid email addresses.
The addresses may look like they are invalid, but they are all conform to
RFC 2822 and could be used. You can use this test set to test your email
validation process.

#### FuzzStringEmpty

```go
func FuzzStringEmpty() []string
```

FuzzStringEmpty returns a test set with a single empty string.

#### FuzzStringFull

```go
func FuzzStringFull() (ret []string)
```

FuzzStringFull contains all string test sets plus ten generated random
strings. This test set is huge and should only be used if you want to make
sure that no string, at all, can crash a process.

#### FuzzStringGenerateRandom

```go
func FuzzStringGenerateRandom(count, length int) (result []string)
```

FuzzStringGenerateRandom returns random strings in a test set.

#### FuzzStringHtmlTags

```go
func FuzzStringHtmlTags() []string
```

FuzzStringHtmlTags returns a test set with different html tags.

Example:
  - <script>
  - <script>alert('XSS')</script>
  - <a href="https://github.com/MarvinJWendt/testza">link</a>

#### FuzzStringLong

```go
func FuzzStringLong() (testSet []string)
```

FuzzStringLong returns a test set with long random strings. Returns: [0]:
Random string (length: 25) [1]: Random string (length: 50) [2]: Random
string (length: 100) [3]: Random string (length: 1,000) [4]: Random string
(length: 100,000)

#### FuzzStringNumeric

```go
func FuzzStringNumeric() []string
```

FuzzStringNumeric returns a test set with strings that are numeric.
The highest number in here is "9223372036854775807", which is equal to the
maxmim int64.

#### FuzzStringUsernames

```go
func FuzzStringUsernames() []string
```

FuzzStringUsernames returns a test set with usernames.

### Fuzz Utils

#### FuzzUtilDistinctSet

```go
func FuzzUtilDistinctSet[setType comparable](testSet []setType) []setType
```

FuzzUtilDistinctSet returns a set with removed duplicates.

Example:

    uniqueSet := testza.FuzzUtilDistinctSet([]string{"A", "C", "A", "B", "A", "B", "C"})
    // uniqueSet => []string{"A", "C", "B"}

#### FuzzUtilLimitSet

```go
func FuzzUtilLimitSet[setType any](testSet []setType, max int) []setType
```

FuzzUtilLimitSet returns a random sample of a test set with a maximal size.

Example:

    limitedSet := testza.FuzzUtilLimitSet(testza.FuzzStringFull(), 10)

#### FuzzUtilMergeSets

```go
func FuzzUtilMergeSets[setType any](sets ...[]setType) (merged []setType)
```

FuzzUtilMergeSets merges multiple test sets into one. All test sets must
have the same type.

Example:

    mergedSet := testza.FuzzUtilMergeSets(testza.FuzzIntGenerateRandomNegative(3, 0), testza.FuzzIntGenerateRandomPositive(2, 0))

#### FuzzUtilModifySet

```go
func FuzzUtilModifySet[setType any](inputSet []setType, modifier func(index int, value setType) setType) (floats []setType)
```

FuzzUtilModifySet returns a modified version of a test set.

Example:

     modifiedSet := testza.FuzzUtilModifySet(testza.FuzzIntFull(), func(i int, value int) int {
    		return i * 2 // double every value in the test set
    	})

#### FuzzUtilRunTests

```go
func FuzzUtilRunTests[setType any](t testRunner, testSet []setType, testFunc func(t *testing.T, index int, f setType))
```

FuzzUtilRunTests runs a test for every value in a test set. You can use
the value as input parameter for your functions, to sanity test against
many different cases. This ensures that your functions have a correct error
handling and enables you to test against hundreds of cases easily.

Example:

    testza.FuzzUtilRunTests(t, testza.FuzzStringEmailAddresses(), func(t *testing.T, index int, emailAddress string) {
    	// Test logic
    	// err := YourFunction(emailAddress)
    	// testza.AssertNoError(t, err)
    	// ...
    })

### 

#### GetColorsEnabled

```go
func GetColorsEnabled() bool
```

GetColorsEnabled returns current value of ColorsEnabled setting.
ColorsEnabled controls if testza should print colored output.

#### GetDiffContextLines

```go
func GetDiffContextLines() int
```

GetDiffContextLines returns current value of DiffContextLines setting.
DiffContextLines setting controls how many lines are shown around a changed
diff line. If set to -1 it will show full diff.

#### GetLineNumbersEnabled

```go
func GetLineNumbersEnabled() bool
```

GetLineNumbersEnabled returns current value of LineNumbersEnabled setting.
LineNumbersEnabled controls if line numbers should be printed in failing
tests.

#### GetRandomSeed

```go
func GetRandomSeed() int64
```

GetRandomSeed returns current value of the random seed setting.

#### GetShowStartupMessage

```go
func GetShowStartupMessage() bool
```

GetShowStartupMessage returns current value of showStartupMessage setting.
showStartupMessage setting controls if the startup message should be
printed.

### Settings

#### SetColorsEnabled

```go
func SetColorsEnabled(enabled bool)
```

SetColorsEnabled controls if testza should print colored output. You should
use this in the init() method of the package, which contains your tests.

> This setting can also be set by the command line flag
--testza.disable-color.

Example:

    init() {
      testza.SetColorsEnabled(false) // Disable colored output
      testza.SetColorsEnabled(true)  // Enable colored output
    }

#### SetDiffContextLines

```go
func SetDiffContextLines(lines int)
```

SetDiffContextLines controls how many lines are shown around a changed diff
line. If set to -1 it will show full diff. You should use this in the init()
method of the package, which contains your tests.

> This setting can also be set by the command line flag
--testza.diff-context-lines.

Example:

    init() {
      testza.SetDiffContextLines(-1) // Show all diff lines
      testza.SetDiffContextLines(3)  // Show 3 lines around every changed line
    }

#### SetLineNumbersEnabled

```go
func SetLineNumbersEnabled(enabled bool)
```

SetLineNumbersEnabled controls if line numbers should be printed in
failing tests. You should use this in the init() method of the package,
which contains your tests.

> This setting can also be set by the command line flag
--testza.disable-line-numbers.

Example:

    init() {
      testza.SetLineNumbersEnabled(false) // Disable line numbers
      testza.SetLineNumbersEnabled(true)  // Enable line numbers
    }

#### SetRandomSeed

```go
func SetRandomSeed(seed int64)
```

SetRandomSeed sets the seed for the random generator used in testza.
Using the same seed will result in the same random sequences each time and
guarantee a reproducible test run. Use this setting, if you want a 100%
deterministic test. You should use this in the init() method of the package,
which contains your tests.

> This setting can also be set by the command line flag --testza.seed.

Example:

    init() {
      testza.SetRandomSeed(1337) // Set the seed to 1337
      testza.SetRandomSeed(time.Now().UnixNano()) // Set the seed back to the current time (default | non-deterministic)
    }

#### SetShowStartupMessage

```go
func SetShowStartupMessage(show bool)
```

SetShowStartupMessage controls if the startup message should be printed.
You should use this in the init() method of the package, which contains your
tests.

> This setting can also be set by the command line flag
--testza.disable-startup-message.

Example:

    init() {
      testza.SetShowStartupMessage(false) // Disable the startup message
      testza.SetShowStartupMessage(true)  // Enable the startup message
    }

### Snapshot

#### SnapshotCreate

```go
func SnapshotCreate(name string, snapshotObject any) error
```

SnapshotCreate creates a snapshot of an object, which can be validated
in future test runs. Using this function directly will override
previous snapshots with the same name. You most likely want to use
SnapshotCreateOrValidate.

NOTICE: \r\n will be replaced with \n to make the files consistent between
operating systems.

Example:

    testza.SnapshotCreate(t.Name(), objectToBeSnapshotted)

#### SnapshotCreateOrValidate

```go
func SnapshotCreateOrValidate(t testRunner, name string, object any, msg ...any) error
```

SnapshotCreateOrValidate creates a snapshot of an object which can be used
in future test runs. It is good practice to name your snapshots the same
as the test they are created in. You can do that automatically by using
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
func SnapshotValidate(t testRunner, name string, actual any, msg ...any) error
```

SnapshotValidate validates an already exisiting snapshot of an object.
You most likely want to use SnapshotCreateOrValidate.

NOTICE: \r\n will be replaced with \n to make the files consistent between
operating systems.

Example:

    testza.SnapshotValidate(t, t.Name(), objectToBeValidated)
    testza.SnapshotValidate(t, t.Name(), objectToBeValidated, "Optional message")


<!-- docs:end -->

---

> Made with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) and contributors! |
> [MarvinJWendt.com](https://marvinjwendt.com)
