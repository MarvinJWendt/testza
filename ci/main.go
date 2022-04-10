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

var Categories = []Category{
	{Name: "Settings", Prefix: "Set"},
	{Name: "Assert", Prefix: "Assert"},
	{Name: "Capture", Prefix: "Capture"},
	{Name: "Fuzz Utils", Prefix: "FuzzUtil"},
	{Name: "Fuzz Booleans", Prefix: "FuzzBool"},
	{Name: "Fuzz Strings", Prefix: "FuzzString"},
	{Name: "Fuzz Float64s", Prefix: "FuzzFloat64"},
	{Name: "Fuzz Integers", Prefix: "FuzzInt"},
	{Name: "Snapshot", Prefix: "Snapshot"},
}

func main() {
	pterm.EnableDebugMessages()

	goDoc = getGoDoc()
	parseGoDoc()

	writeBetween("README.md", "docs", getMarkdown())
}

type Category struct {
	Name   string
	Prefix string
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

	Functions = append(Functions, lastFunc)

	for i, f := range Functions {
		if strings.TrimSpace(f.Head) == "" {
			continue
		}

		var re = regexp.MustCompile(`(?m)func (?P<name>[a-zA-Z0-9]*)`)
		Functions[i].Name = regexGroupsToMap(re, f.Head)["name"]

		var newBody string
		for _, v := range strings.Split(f.Body, "\n") {
			newBody += strings.TrimPrefix(v, "    ") + "\n"
		}
		Functions[i].Body = strings.TrimRight(newBody, "\n")
	}

	pterm.Debug.Printfln("Found %d functions:", len(Functions))
	for _, f := range Functions {
		pterm.Debug.Println(f.Name)
	}
}

func pathToMarkdownLink(path string) string {
	path = strings.ReplaceAll(path, " ", "-")
	path = strings.ReplaceAll(path, ".", "")

	return path
}

func getCategoryOfFunctionName(name string) (c Category) {
	for _, category := range Categories {
		if strings.HasPrefix(name, category.Prefix) {
			c = category
		}
	}

	return
}

func getMarkdown() (md string) {
	var lastCategory Category

	md += `<table>
  <tr>
    <th>Module</th>
    <th>Methods</th>
  </tr>`

	for _, category := range Categories {
		path := category.Name
		path = pathToMarkdownLink(path)
		md += "<tr>\n"
		md += fmt.Sprintf(`<td><a href="https://github.com/MarvinJWendt/testza#%s">%s</a></td>`+"\n", path, category.Name)

		md += "<td>\n\n<details>\n<summary>Click to expand</summary>\n\n"
		for _, f := range Functions {
			if strings.TrimSpace(f.Head) == "" {
				continue
			}

			if strings.HasPrefix(f.Name, category.Prefix) {
				md += fmt.Sprintf("  - [%s](https://github.com/MarvinJWendt/testza#%s)\n", f.Name, pathToMarkdownLink(f.Name))
			}
		}
		md += "</td>\n\n</details>\n\n"
		md += "</tr>\n"
	}

	md += "</table>\n\n"

	for _, f := range Functions {
		if strings.TrimSpace(f.Head) == "" {
			continue
		}

		category := getCategoryOfFunctionName(f.Name)
		if category != lastCategory {
			md += fmt.Sprintf("### %s\n\n", category.Name)
		}
		lastCategory = category

		md += "#### " + f.Name + "\n\n"
		var re = regexp.MustCompile(`(?m)func \(.*?\)`)
		md += "```go\n" + re.ReplaceAllString(f.Head, "func") + "\n```\n\n"
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
