<a name="unreleased"></a>
## [Unreleased]

### Features
- **capture:** add `CaptureStdoutAndStderr`
- **mock:** add `Stdin` mocking

### Code Refactoring
- **mock-stdin:** remove unused test runner from `MockStdinString` function
- **mock-stdin:** wrap errors


<a name="v0.2.5"></a>
## [v0.2.5] - 2021-08-23
### Bug Fixes
- **assert:** fix `AssertNil` and `AssertNotNil` on non basic types

### Test
- **assert:** add tests for `AssertNil` with non basic types


<a name="v0.2.4"></a>
## [v0.2.4] - 2021-08-19
### Features
- add configuratable options

### Test
- **configuration:** add tests to configuration values


<a name="v0.2.3"></a>
## [v0.2.3] - 2021-08-17
### Bug Fixes
- **assert:** get correctly formatted error messages on mocked test runner


<a name="v0.2.2"></a>
## [v0.2.2] - 2021-08-16

<a name="v0.2.1"></a>
## [v0.2.1] - 2021-08-12
### Features
- **snapshot:** add snapshot functionality

### Test
- delete invalid named snapshot
- **snapshot:** make tests pass on linux
- **snapshot:** add more error tests to Snapshot
- **snapshot:** add Snapshot tests

### Code Refactoring
- fix linting and make 1.15.X tests pass
- **snapshot:** make linebreaks consistent


<a name="v0.2.0"></a>
## [v0.2.0] - 2021-07-21
### Features
- **assert:** add `AssertTestFails`

### Test
- rename test functions to new structure
- **assert:** test that assertions fail when they should
- **assert:** cleanup tests for `AssertTestFails`
- **assert:** add more tests to `AssertEqual`
- **assert:** add tests for `AssertTestFails`
- **mock:** add test for `MockInputIntModify`
- **mock:** add test for `MockInputFloat64Modify`
- **mock:** add test for `MockInputFloat64GenerateRandomPositive` and `MockInputFloat64GenerateRandomNegative`
- **mock:** add test for `MockInputBoolModify`

### Code Refactoring
- move tests into own package
- **assert:** change `AssertPanic` to `AssertPanics`

### BREAKING CHANGE

change `AssertPanic` to `AssertPanics`


<a name="v0.1.0"></a>
## [v0.1.0] - 2021-07-19
### Code Refactoring
- rewrite CI to new structure
- change structure

### BREAKING CHANGE

Functions have a new structure.


<a name="v0.0.3"></a>
## [v0.0.3] - 2021-07-17
### Features
- **mock-input-strings:** add more email addresses


<a name="v0.0.2"></a>
## [v0.0.2] - 2021-07-14
### Features
- **mock-inputs-string:** add `Long` test set

### Test
- **mock-input-strings:** add tests for `Limit`


<a name="v0.0.1"></a>
## v0.0.1 - 2021-07-13
### Features
- only show difference if the objects are named `Expected` and `Actual`
- add `Getter` and assertion methods
- add integer mocking
- add difference to failure messages that have two objects
- add line numbers to failure messages
- add `Use.Assert.KindOf` and `Use.Assert.NotKindOf`
- add float64 mocking
- add `Use.Assert.Number` and `Use.Assert.NotNumber`
- add more styling to test errors
- upload code
- **assert:** add `CompletesIn` and `NotCompletesIn`
- **assert:** add `Nil` and `NotNil` assertions
- **assert:** add `Panic` and `NotPanic`
- **assert:** add `NoError`
- **assert:** add `Greater` and `Less`
- **assert:** add `Contains`
- **capture:** add `Stderr`
- **capture:** add `Stdout`
- **mock:** add boolean mockings
- **mock-ints:** add `GenerateRandomRange`
- **mock-strings:** add numeric strings set
- **mock-strings:** add email addresses and empty string sets

### Bug Fixes
- check both errors
- check both errors
- fix failure objects with no `DataStyle`
- remove blank line at the end of failure messages
- fix Assert naming

### Test
- do nothing in `NotPanic` test
- **assert:** try to fix macOS weirdness
- **assert:** more buffer for `NotCompletesIn` test
- **assert:** add `True` and `False` tests
- **mock-string:** add tests to `GenerateRandom`

### Code Refactoring
- tidy project
- rename `testingT` to `testRunner`
- move input mocking methods to `Use.Mock.Inputs`
- add easier methods for unknown objects
- move from atomicgo to MarvinJWendt
- export Helper structs
- add `Use` variable and `custom_readme` setting
- remove unused import alias
- add errors.go
- replace all `testutil` with `testza`
- change import path to `testza`
- change package name to `testza`
- rename `Input` to `Mock`
- set up template
- **assert:** rename `Number` to `Numeric`
- **internal:** rewrite `Fail` for a nicer output


[Unreleased]: https://github.com/MarvinJWendt/testza/compare/v0.2.5...HEAD
[v0.2.5]: https://github.com/MarvinJWendt/testza/compare/v0.2.4...v0.2.5
[v0.2.4]: https://github.com/MarvinJWendt/testza/compare/v0.2.3...v0.2.4
[v0.2.3]: https://github.com/MarvinJWendt/testza/compare/v0.2.2...v0.2.3
[v0.2.2]: https://github.com/MarvinJWendt/testza/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/MarvinJWendt/testza/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/MarvinJWendt/testza/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/MarvinJWendt/testza/compare/v0.0.3...v0.1.0
[v0.0.3]: https://github.com/MarvinJWendt/testza/compare/v0.0.2...v0.0.3
[v0.0.2]: https://github.com/MarvinJWendt/testza/compare/v0.0.1...v0.0.2
