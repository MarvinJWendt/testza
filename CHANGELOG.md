<a name="unreleased"></a>
## [Unreleased]

### Features
- **assert:** added `AssertFileExists` and `AssertNoFileExists`


<a name="v0.2.9"></a>
## [v0.2.9] - 2021-09-21
### Code Refactoring
- **assert:** renamed internal variable
- **assert:** moved `AssertCompareHelper` to assertion_helper.go
- **assert:** moved getter functions to internal package
- **internal:** moved `AssertRegexpHelper` to right file


<a name="v0.2.8"></a>
## [v0.2.8] - 2021-09-19
### Features
- **assert:** added `AssertRegexp` and `AssertNotRegexp`
- **assert:** added `AssertIncreasing` and `AssertDecreasing`
- **assert:** added `AssertLen`
- **assert:** added error message to `AssertNoError`

### Code Refactoring
- **assert:** added regexp tests for `AssertRegexp` & `AssertNotRegexp`
- **assert:** changed parameter for `Assert(Not)Regexp` to interface
- **assert:** changed parameter for `Assert(Not)Regexp` to interface
- **assert:** changed parameter for `Assert(Not)Regexp` to interface
- **assert:** rearranged parameters
- **assert:** change fail message
- **assert:** rearranged imports
- **assert:** removed unused break statements
- **assert:** renamed variable in `AssertLen`


<a name="v0.2.7"></a>
## [v0.2.7] - 2021-08-27
### Features
- **assert:** added `AssertErrorIs` and `AssertNotErrorIs`

### Test
- **assert:** added tests for `AssertErrorIs` and `AssertNotErrorIs`


<a name="v0.2.6"></a>
## [v0.2.6] - 2021-08-24
### Features
- **capture:** add `CaptureStdoutAndStderr`

### Code Refactoring
- rename variable in `MockInputXXXModify` functions


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


[Unreleased]: https://github.com/MarvinJWendt/testza/compare/v0.2.9...HEAD
[v0.2.9]: https://github.com/MarvinJWendt/testza/compare/v0.2.8...v0.2.9
[v0.2.8]: https://github.com/MarvinJWendt/testza/compare/v0.2.7...v0.2.8
[v0.2.7]: https://github.com/MarvinJWendt/testza/compare/v0.2.6...v0.2.7
[v0.2.6]: https://github.com/MarvinJWendt/testza/compare/v0.2.5...v0.2.6
[v0.2.5]: https://github.com/MarvinJWendt/testza/compare/v0.2.4...v0.2.5
[v0.2.4]: https://github.com/MarvinJWendt/testza/compare/v0.2.3...v0.2.4
[v0.2.3]: https://github.com/MarvinJWendt/testza/compare/v0.2.2...v0.2.3
[v0.2.2]: https://github.com/MarvinJWendt/testza/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/MarvinJWendt/testza/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/MarvinJWendt/testza/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/MarvinJWendt/testza/compare/v0.0.3...v0.1.0
[v0.0.3]: https://github.com/MarvinJWendt/testza/compare/v0.0.2...v0.0.3
[v0.0.2]: https://github.com/MarvinJWendt/testza/compare/v0.0.1...v0.0.2
