


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
<strong><a href="https://pkg.go.dev/github.com/MarvinJWendt/testza#section-documentation" target="_blank">Documentation</a></strong>
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


<!-- docs:end -->

---

> Made with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) and contributors! |
> [MarvinJWendt.com](https://marvinjwendt.com)



