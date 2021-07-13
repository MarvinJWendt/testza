<a name="unreleased"></a>
## [Unreleased]

### Ci
- add changelog generation


<a name="v0.0.1"></a>
## v0.0.1 - 2021-07-13
### Bug Fixes
- check both errors
- check both errors
- fix failure objects with no `DataStyle`
- remove blank line at the end of failure messages
- fix Assert naming

### Chore
- set minimum Go version to 1.15
- **mock:** rename `Input` to `Mock`
- **settings:** update repo description

### Ci
- add ACCESS_TOKEN to `actions/checkout`
- give the custom CI access to ACCESS_TOKEN
- do not include parent struct in function head
- disable some linting checks
- disable `exhaustive` checks
- add CodeQL analysis
- add newline after ToC
- fix module path for `Mock.Ints`
- temporarily remove `Mock` from module list, as it has no methods
- generate ToC for docs
- simplify regex
- add custom CI-System
- disable verbose output in Go tests
- add matrix tests for different Go versions
- disable `thelper` check
- **codeql:** disable cron job
- **golangci:** bump version

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

### Documentation Changes
- update package docs
- add CONTRIBUTING.md
- add CODE_OF_CONDUCT.md
- update readme
- fix spelling
- update license year
- **assert:** add comments to assertion functions
- **readme:** added cheese
- **readme:** add slogan
- **readme:** add syntax example
- **readme:** styling
- **readme:** styling
- **readme:** better description
- **readme:** add screenshot
- **readme:** experiment with screenshot position
- **readme:** experiment with screenshot position
- **readme:** update documentation link
- **readme:** make example syntax easier to understand

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

### Style
- fix linting issues

### Test
- do nothing in `NotPanic` test
- **assert:** try to fix macOS weirdness
- **assert:** more buffer for `NotCompletesIn` test
- **assert:** add `True` and `False` tests
- **mock-string:** add tests to `GenerateRandom`


[Unreleased]: https://github.com/atomicgo/atomicgo/compare/v0.0.1...HEAD
