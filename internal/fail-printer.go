package internal

import (
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
)

var Highlight = pterm.NewStyle(pterm.Bold, pterm.FgLightRed).Sprint

type Objects map[string]interface{}

func NewObjectsExpectedActual(expected, actual interface{}) Objects {
	return Objects{
		"expected": expected,
		"actual":   actual,
	}
}

func NewObjectsSingleUnknown(obj interface{}) Objects {
	return Objects{"object": obj}
}

func NewObjectsSingleNamed(name string, obj interface{}) Objects {
	return Objects{name: obj}
}

func ModifyWrappedText(text, wrappingString string, modifier func(wrappedText string) string) string {
	r := regexp.MustCompile(wrappingString + "(.*?)" + wrappingString)

	res := r.ReplaceAllStringFunc(text, func(s string) string {
		return modifier(s)
	})

	return strings.ReplaceAll(res, wrappingString, "")
}

func FailS(message string, objects Objects, args ...interface{}) string {
	message = ModifyWrappedText(message, "!!", func(wrappedText string) string {
		return Highlight(wrappedText)
	})

	message = pterm.Sprint(args...) + "\n" + message

	if !strings.HasSuffix(message, "\n") {
		message += "\n"
	}

	for k, v := range objects {
		message += pterm.Sprintf("\n%s:\n%s", k, spew.Sdump(v))
	}

	return message
}

func Fail(t testingT, message string, objects Objects, args ...interface{}) {
	t.Error(FailS(message, objects, args...))
}
