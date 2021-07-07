package testza

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Use util functions of testza directly in the default testing system of Go.
// Methods in here integrate directly with the default Go testing system, and give detailed output.
// The methods will trigger the test to fail, if your expected behavior didn't occur.
var Use Helper

// Helper contains helper methods for the Go testing system.
// Methods in here integrate directly with the default Go testing system, and give detailed output.
// The methods will trigger the test to fail, if your expected behavior didn't occur.
type Helper struct {
	Assert  AssertHelper
	Mock    MockHelper
	Capture CaptureHelper
}

// AssertHelper contains assertion methods for the Go testing system.
type AssertHelper struct {
}

type MockInputsHelper struct {
	Strings  MockInputsStringsHelper
	Floats64 MockInputsFloats64Helper
	Ints     MockInputsIntsHelper
	Bools    MockInputsBoolsHelper
}

// MockHelper contains mocking methods for the Go testing system.
// Do not use this struct directly, use the methods in Use.
type MockHelper struct {
	Inputs MockInputsHelper
}
