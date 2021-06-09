<h1 align="center">AtomicGo | testutil</h1>

<p align="center">

<a href="https://github.com/atomicgo/testutil/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/testutil?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/testutil" target="_blank">
<img src="https://img.shields.io/github/workflow/status/atomicgo/testutil/Go?label=tests&style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/testutil" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/testutil?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/testutil">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-0-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://github.com/atomicgo/testutil/issues">
<img src="https://img.shields.io/github/issues/atomicgo/testutil.svg?style=flat-square" alt="Issues">
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

</p>

---

<p align="center">
<strong><a href="#install">Get The Module</a></strong>
|
<strong><a href="https://pkg.go.dev/github.com/atomicgo/testutil#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

## Description

Package testutil contains util functions for writing tests in Go.

## Install

```console
# Execute this command inside your project
go get -u github.com/atomicgo/testutil
```

```go
// Add this to your imports
import "github.com/atomicgo/testutil"
```

## Usage

```go
var AssertHelper assert
```

#### type InputHelper

```go
type InputHelper struct {
	Strings StringsHelper
}
```


```go
var Input InputHelper
```
Input contains test sets, which you can pass to a function as input parameters
and validate the output.

#### type StringsHelper

```go
type StringsHelper struct{}
```


#### func (StringsHelper) All

```go
func (s StringsHelper) All() (ret []string)
```
All contains all string test sets plus ten generated random StringsHelper.

#### func (StringsHelper) GenerateRandom

```go
func (s StringsHelper) GenerateRandom(length, count int) (result []string)
```
GenerateRandom returns random StringsHelper in a test set.

#### func (StringsHelper) HtmlTags

```go
func (s StringsHelper) HtmlTags() []string
```
HtmlTags returns a test set with html tags.

#### func (StringsHelper) Limit

```go
func (s StringsHelper) Limit(testSet []string, max int) []string
```
Limit limits a test set in size.

#### func (StringsHelper) Modify

```go
func (s StringsHelper) Modify(inputSlice []string, f func(index int, value string) string) (ret []string)
```
Modify returns a modified version of a test set.

#### func (StringsHelper) RunTests

```go
func (s StringsHelper) RunTests(t TestingT, testSet []string, testFunc func(t *testing.T, index int, str string))
```
RunTests runs tests with a specific test set.

#### func (StringsHelper) Usernames

```go
func (s StringsHelper) Usernames() []string
```
Usernames returns a test set with usernames.

#### type TestingT

```go
type TestingT interface {
	Error(args ...interface{})
}
```

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
