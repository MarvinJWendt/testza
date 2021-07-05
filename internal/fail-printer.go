package internal

import (
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
)

var Highlight = pterm.NewStyle(pterm.Bold, pterm.FgLightRed).Sprint

type Object struct {
	Name      string
	NameStyle *pterm.Style
	Data      interface{}
	DataStyle *pterm.Style
}

type Objects []Object

func NewObjectsExpectedActual(expected, actual interface{}) Objects {
	return Objects{
		{
			Name:      "Expected",
			NameStyle: pterm.NewStyle(pterm.FgLightGreen),
			Data:      expected,
			DataStyle: pterm.NewStyle(pterm.FgGreen),
		},
		{
			Name:      "Actual",
			NameStyle: pterm.NewStyle(pterm.FgLightRed),
			Data:      actual,
			DataStyle: pterm.NewStyle(pterm.FgRed),
		},
	}
}

func NewObjectsSingleUnknown(obj interface{}) Objects {
	return Objects{
		{
			Name:      "Object",
			Data:      obj,
			DataStyle: pterm.NewStyle(pterm.FgYellow),
		},
	}
}

func NewObjectsSingleNamed(name string, obj interface{}) Objects {
	return Objects{
		{
			Name: name,
			Data: obj,
		},
	}
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

	for _, v := range objects {
		if v.NameStyle == nil {
			v.NameStyle = pterm.NewStyle(pterm.FgCyan)
		}
		message += pterm.Sprintf("\n%s\n%s", v.NameStyle.Add(*pterm.NewStyle(pterm.Bold)).Sprint(v.Name+":"), v.DataStyle.Sprint(spew.Sdump(v.Data)))
	}

	message = "\n" + strings.Join(strings.Split(message, "\n"), "\n"+pterm.FgRed.Sprint("| "))

	return message
}

func Fail(t testingT, message string, objects Objects, args ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	t.Error(FailS(message, objects, args...))
}
