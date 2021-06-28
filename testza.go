package testza

// Use util functions of testza directly in the default testing system of Go.
// Methods in here integrate directly with the default Go testing system, and give detailed output.
// The methods will trigger the test to fail, if your expected behaviour didn't occur.
// If you want to get the results of those functions, you can use the methods in Getter via testza.Get.
var Use Helper

// Get contains methods, that can be used for testing different scenarios.
// Methods in here return the raw value and do not integrate with the Go testing system.
// You can use the Getter methods directly in your code.
// If you want to use testza in Go test files, you should use Helper via testza.Use.
var Get Getter

// Helper contains helper methods for the Go testing system.
// Methods in here integrate directly with the default Go testing system, and give detailed output.
// The methods will trigger the test to fail, if your expected behaviour didn't occur.
// If you want to get the results of those functions, you can use the methods in Getter via testza.Get.
type Helper struct {
	Assert AssertHelper
	Mock   MockHelper
}

// Getter contains methods, that can be used for testing different scenarios.
// Methods in here return the raw value and do not integrate with the Go testing system.
// You can use the Getter methods directly in your code.
// If you want to use testza in Go test files, you should use Helper via testza.Use.
type Getter struct {
	Assert AssertGetter
}

// AssertHelper contains assertion methods for the Go testing system.
type AssertHelper struct {
}

// AssertGetter contains assertion getter methods, that can be used everywhere.
type AssertGetter struct {
}

// MockHelper contains mocking methods for the Go testing system.
type MockHelper struct {
	Strings StringsHelper
}
