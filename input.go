package testutil

// Input contains test sets, which you can pass to a function as input parameters and validate the output.
var Input InputHelper

type InputHelper struct {
	Strings StringsHelper
}
