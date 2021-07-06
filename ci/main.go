// Custom CI-System for https://github.com/MarvinJWendt/testza.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pterm/pterm"
)

var goDoc string

var Functions []Function

var Modules = []Module{
	{Name: "Assertions", StructName: "AssertHelper", Path: "testza.Use.Assert"},
	{Name: "Capture", StructName: "CaptureHelper", Path: "testza.Use.Capture"},
	{Name: "Mocking", StructName: "MockHelper", Path: "testza.Use.Mock"},
	{Name: "Mock Strings", StructName: "StringsHelper", Path: "testza.Use.Mock.Strings"},
	{Name: "Mock Booleans", StructName: "BoolsHelper", Path: "testza.Use.Mock.Bools"},
	{Name: "Mock Floats64", StructName: "Floats64Helper", Path: "testza.Use.Mock.Floats64"},
	{Name: "Mock Integers", StructName: "IntsHelper", Path: "testza.Use.Mock.Floats64"},
}

func main() {
	goDoc = getGoDoc()
	parseGoDoc()

	pterm.EnableDebugMessages()

	writeBetween("README.md", "docs", getMarkdown())
}

type Module struct {
	Name       string
	StructName string
	Path       string
}

type Function struct {
	Head string
	Body string
	Name string
	Path string
}

func getGoDoc() string {
	out, err := exec.Command("go", "doc", "-all").Output()
	pterm.Fatal.PrintOnError(err)

	return string(out)
}

func parseGoDoc() {
	var insideFunctionBlock bool
	var lastFunc Function

	for _, line := range strings.Split(goDoc, "\n") {
		if strings.HasPrefix(line, "func") {
			insideFunctionBlock = true
			Functions = append(Functions, lastFunc)
			lastFunc = Function{}
			lastFunc.Head = line
		} else if insideFunctionBlock {
			if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\n") || line == "" {
				lastFunc.Body += line + "\n"
			} else {
				insideFunctionBlock = false
			}
		}
	}

	for i, f := range Functions {
		if strings.TrimSpace(f.Head) == "" {
			continue
		}

		var re = regexp.MustCompile(`(?m)\)( (?P<name>.*?)\()`)
		Functions[i].Name = regexGroupsToMap(re, f.Head)["name"]

		var newBody string
		for _, v := range strings.Split(f.Body, "\n") {
			newBody += strings.TrimPrefix(v, "    ") + "\n"
		}
		Functions[i].Body = strings.TrimRight(newBody, "\n")
	}
}

func getModuleOfObject(head string) Module {
	head = strings.TrimLeft(head, "*")

	var re = regexp.MustCompile(`(?m)(?P<name>[a-zA-Z1-9]*)?\)`)
	parent := regexGroupsToMap(re, head)["name"]

	for _, module := range Modules {
		if module.StructName == parent {
			return module
		}
	}

	return Module{}
}

func getMarkdown() (md string) {
	var lastModule Module

	for _, f := range Functions {
		if strings.TrimSpace(f.Head) == "" {
			continue
		}

		module := getModuleOfObject(f.Head)
		if module.StructName != lastModule.StructName {
			md += fmt.Sprintf("### %s\n\n", module.Name)
		}
		lastModule = module

		md += "#### " + module.Path + "." + f.Name + "\n\n"
		md += "```go\n" + f.Head + "\n```\n\n"
		md += f.Body + "\n"

		md += "\n"
	}

	return
}

func regexGroupsToMap(r *regexp.Regexp, s string) map[string]string {
	names := r.SubexpNames()
	result := r.FindAllStringSubmatch(s, -1)
	m := map[string]string{}
	for i, n := range result[0] {
		m[names[i]] = n
	}
	return m
}

func writeBetween(file, name, insertText string) {
	out, err := os.ReadFile(file)
	pterm.Fatal.PrintOnError(err)
	original := string(out)

	beforeRegex := regexp.MustCompile(`(?ms).*<!-- ` + name + `:start -->`)
	afterRegex := regexp.MustCompile(`(?ms)<!-- ` + name + `:end -->.*`)
	before := beforeRegex.FindAllString(original, 1)[0]
	after := afterRegex.FindAllString(original, 1)[0]

	ret := before
	ret += "\n" + insertText + "\n"
	ret += after

	err = os.WriteFile(file, []byte(ret), 0600)
	pterm.Fatal.PrintOnError(err)
}
