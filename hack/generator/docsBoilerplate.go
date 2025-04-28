package generator

import (
	"path/filepath"
	"strings"
	"text/template"
)

type Store interface {
	Load(from string) (string, error)
	Save(data, to string) error
	Delete(from string) error
}

type DocsBoilerplate struct {
	targetPath string
	tmpl       DocsFile
	tmplData   DocsTemplateData
}

func NewDocsBoilerplate(
	targetPath string,
	tmpl DocsFile,
	tmplData DocsTemplateData,
) *DocsBoilerplate {
	return &DocsBoilerplate{targetPath, tmpl, tmplData}
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

func (dob *DocsBoilerplate) Save(store Store) error {
	data, err := dob.Content()
	if err != nil {
		return err
	}

	return store.Save(data, dob.docsFilePath())
}

func (dob *DocsBoilerplate) Delete(store Store) error {
	return store.Delete(dob.docsFilePath())
}

func (dob *DocsBoilerplate) docsFilePath() string {
	docsFileName := dob.tmplData.SmallTitle + ".md"
	return filepath.Join(dob.targetPath, docsFileName)
}
