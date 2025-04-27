package generator

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
)

type StorageLocation interface {
	Load(from string) (string, error)
	Save(data, to string) error
	Delete(from string) error
}

type DocsBoilerplate struct {
	conf     Config
	tmpl     DocsFile
	tmplData DocsTemplateData
}

func NewDocsBoilerplate(conf Config, tmpl DocsFile, tmplData DocsTemplateData) *DocsBoilerplate {
	return &DocsBoilerplate{conf, tmpl, tmplData}
}

func (dob *DocsBoilerplate) Content() (string, error) {
	tmplFileContent, err := dob.tmpl.Content()
	if err != nil {
		return "", err
	}

	tmpl, err := template.New(dob.tmplData.Title).Parse(tmplFileContent)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, dob.tmplData); err != nil {
		return "", err
	}

	return result.String(), nil
}

func (dob *DocsBoilerplate) Save(fs Filesystem) error {
	fmt.Println("Hello")
	data, err := dob.Content()
	if err != nil {
		return err
	}

	fmt.Println(dob.docsFilePath())

	return fs.Save(data, dob.docsFilePath())
}

func (dob *DocsBoilerplate) docsFilePath() string {
	docsFileName := dob.smallKebab() + ".md"
	return filepath.Join(dob.conf.To, docsFileName)
}

func (dob *DocsBoilerplate) smallKebab() string {
	small := strings.ToLower(dob.tmplData.Title)

	if len(strings.Split(small, " ")) > 1 {
		small = strings.ReplaceAll(small, " ", "-")
	}
	return small
}
