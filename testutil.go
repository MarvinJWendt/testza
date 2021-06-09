package testutil

var Use Helper

type Helper struct {
	Assert AssertHelper
	Input  InputHelper
}

type AssertHelper struct{}

type InputHelper struct {
	Strings StringsHelper
}
