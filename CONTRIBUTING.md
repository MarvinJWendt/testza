# Contributing to Testza

This guide will explain how you can contribute to testza.

## Function Naming

Every function in testza follows a specific name scheme.
As you can see in the [documentations Table of Contents](https://github.com/MarvinJWendt/testza#documentation), testza has many different modules.
The first word of a new function, is always the module name (eg. `Assert`, `Capture`, `Snapshot`, etc.).
If the function is inside a nested sub-module, the module names will be chained as in `FuzzBool`.
The last word of the function name is the actual function name itself.

### Example

```
  ┌ Fuzz part of testza
  |   ┌ Fuzz different inputs
  |   |   ┌ Fuzz integer inputs
  |   |   |   ┌ Full set of integer input fuzzing
  |   |   |   |
FuzzIntFull()
```

## File Naming

File names should describe a single module, and the content should be the functions of the module.
Tests should have the same file name with `_test` as a suffix.

## Writing Tests

Every function of testza has to be tested. As testza is a test framework, it is a convenient choice to write all tests with testza.
The tests should be consistent, so it's best to look at a few tests before writing your own.

## Documenting Functions

Each function of testza must be well documented.
The documentation in the README file is automatically generated from the comments of the functions.
Functions are documented according to the Go standard and should convey what the function does and when to use it.
At the end of each function there should be an example of how the function can be used.

## Adding Examples to Functions

At the end of every function documentation, there must be an example. 
The example must be separated by the actual documentation by a blank line.
The first line of the example is `// Example:`. Now the examples follow.
Every line of the actual example must be indeted by 2 spaces after the slashes (`//  testa.ExampleFunction(t, usage...)`).
A full example might look like this:

```go
// AssertEqual asserts that two objects are equal.
//
// Example:
//  testza.AssertEqual(t, "Hello, World!", "Hello, World!")
//  testza.AssertEqual(t, true, true)
```
