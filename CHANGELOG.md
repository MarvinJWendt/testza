<a name="unreleased"></a>
## [Unreleased]

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


[Unreleased]: https://github.com/MarvinJWendt/testza/compare/v0.0.1...HEAD
