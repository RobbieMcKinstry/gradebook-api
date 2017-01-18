package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Bug struct {
	Name        string
	Description string
}

func main() {

	bytes, err := ioutil.ReadFile("./checkstyle_checks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileContents := string(bytes)
	lines := strings.Split(fileContents, "\n")

	var bugs = []Bug{}

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		fields := strings.SplitN(line, " ", 2)
		newBug := Bug{}
		newBug.Name = strings.TrimSpace(fields[0])
		newBug.Description = strings.TrimSpace(fields[1])
		bugs = append(bugs, newBug)
	}

	t := template.Must(template.New("goa").Parse(tmpl))

	f, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(f, bugs)
	if err != nil {
		log.Fatal(err)
	}

}

var tmpl = `
{{ range . }}
Attribute("{{.Name}}", Boolean, "{{.Description}}")
{{end}}
`
