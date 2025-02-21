package generator

import (
	"os"
	"strings"
	"text/template"
)

func Docs(notParsed, title string) (string, error) {
	docsTmpl, err := template.New("docs").Parse(notParsed)
	if err != nil {
		return "", err
	}

	small := strings.ToLower(title)

	if len(strings.Split(small, " ")) > 1 {
		small = strings.ReplaceAll(small, " ", "-")
	}

	data := struct {
		Title      string
		SmallTitle string
	}{
		Title:      title,
		SmallTitle: small,
	}

	var result strings.Builder
	if err := docsTmpl.Execute(&result, data); err != nil {
		return "", err
	}

	return result.String(), nil
}

func LoadDocs(path string) (string, error) {
	result, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(result), err
}

func Sample() {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
