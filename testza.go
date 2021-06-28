package testza

// Use util functions of testza.
var Use Helper

// Helper contains every util function in a structured format for easy usage.
type Helper struct {
	Assert AssertHelper
	Mock   MockHelper
}

// AssertHelper contains assertion functions.
type AssertHelper struct {
}

// MockHelper contains helper functions for different input types.
type MockHelper struct {
	Strings StringsHelper
}
