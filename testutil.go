package testutil

// Use util functions of testutil.
var Use Helper

// Helper contains every util function in a structured format for easy usage.
type Helper struct {
	Assert AssertHelper
	Input  InputHelper
}

// AssertHelper contains assertion functions.
type AssertHelper struct{}

// InputHelper contains helper functions for different input types.
type InputHelper struct {
	Strings StringsHelper
}
