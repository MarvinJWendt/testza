package internal

import (
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pterm/pterm"
)

var Highlight = pterm.NewStyle(pterm.FgLightRed).Sprint

type Object struct {
	Name      string
	NameStyle *pterm.Style
	Data      any
	DataStyle *pterm.Style
	Raw       bool
}

type Objects []Object

func NewObjectsExpectedActual(expected, actual any) Objects {
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

func NewObjectsExpectedActualWithDiff(expected, actual any) Objects {
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
		NewDiffObject(expected, actual),
	}
}

func NewObjectsUnknown(objs ...any) Objects {
	return Objects{
		{
			Name:      "Object",
			NameStyle: pterm.NewStyle(pterm.FgYellow),
			Data:      objs,
		},
	}
}

func NewObjectsSingleUnknown(obj any) Objects {
	return Objects{
		{
			Name:      "Object",
			NameStyle: pterm.NewStyle(pterm.FgMagenta),
			Data:      obj,
		},
	}
}

func NewObjectsSingleNamed(name string, obj any) Objects {
	return Objects{
		{
			Name:      name,
			NameStyle: pterm.NewStyle(pterm.FgMagenta),
			Data:      obj,
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

func FailS(message string, objects Objects, args ...any) string {
	message = ModifyWrappedText(message, "!!", func(wrappedText string) string {
		return Highlight(wrappedText)
	})

	if len(args) > 0 {
		message += "\n\n" + pterm.FgMagenta.Sprint("Message: ") + pterm.Sprintf(pterm.Sprint(args[0]), args[1:]...) + "\n"
	}

	if !strings.HasSuffix(message, "\n") {
		message += "\n"
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
			message += v.DataStyle.Sprint(v.Data)
		}
	}

	// Prepend line numbers
	newMessage := "\n"
	lines := strings.Split(message, "\n")
	for i, line := range lines {
		if !(i == 0 && strings.TrimSpace(line) == "") && i < len(lines)-1 {
			if LineNumbersEnabled {
				newMessage += pterm.FgGray.Sprintf("%4d| ", i+1) + line + "\n" + pterm.Reset.Sprint()
			} else {
				newMessage += line + "\n"
			}
		}
	}
	message = "\n" + newMessage + "\n"

	return message
}

func Fail(t testRunner, message string, objects Objects, args ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	t.Error(FailS(message, objects, args...))
}
