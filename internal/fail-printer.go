package internal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
)

var Highlight = pterm.NewStyle(pterm.FgLightRed).Sprint

type Object struct {
	Name      string
	NameStyle *pterm.Style
	Data      interface{}
	DataStyle *pterm.Style
	Raw       bool
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

func NewObjectsUnknown(objs ...interface{}) Objects {
	return Objects{
		{
			Name:      "Object",
			NameStyle: pterm.NewStyle(pterm.FgYellow),
			Data:      objs,
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

	if len(objects) == 2 {
		if strings.Count(spew.Sdump(objects[0].Data), "\n")+strings.Count(spew.Sdump(objects[1].Data), "\n") > 4 &&
			objects[0].Name == "Expected" &&
			objects[1].Name == "Actual" {
			objects = append(Objects{{
				Name:      "Difference",
				NameStyle: pterm.NewStyle(pterm.FgYellow),
				Data:      GetDifference(objects[0].Data, objects[1].Data),
				Raw:       true,
			}}, objects...)
		}
	}

	for i, v := range objects {
		if v.DataStyle == nil {
			objects[i].DataStyle = pterm.NewStyle()
		}

		if v.NameStyle == nil {
			objects[i].NameStyle = pterm.NewStyle()
		}
	}

	for _, v := range objects {
		if v.NameStyle == nil {
			v.NameStyle = pterm.NewStyle(pterm.FgCyan)
		}
		message += "\n" + v.NameStyle.Add(*pterm.NewStyle(pterm.Bold)).Sprint(v.Name+":") + "\n"
		if !v.Raw {
			message += v.DataStyle.Sprint(spew.Sdump(v.Data))
		} else {
			message += fmt.Sprint(v.Data)
		}
	}

	newMessage := "\n"
	lines := strings.Split(message, "\n")
	for i, line := range lines {
		if i > 0 && i < len(lines)-1 {
			newMessage += pterm.FgGray.Sprintf("%4d| ", i) + line + "\n" + pterm.Reset.Sprint()
		}
	}
	message = "\n" + newMessage + "\n"

	return message
}

func Fail(t testingT, message string, objects Objects, args ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	t.Error(FailS(message, objects, args...))
}
