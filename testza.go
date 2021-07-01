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
// If you want to get the results of those functions, you can use the methods in Getter via testza.Get.
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

// MockHelper contains mocking methods for the Go testing system.
type MockHelper struct {
	Strings  StringsHelper
	Floats64 Floats64Helper
	Ints     IntsHelper
}
