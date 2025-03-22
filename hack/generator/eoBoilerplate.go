package generator

import (
	"strings"
	"text/template"
)

type Filesystem interface {
	Load(from string) (string, error)
	Save(data, to string) error
}

type Config struct {
	To   string
	From string
}

type OnDiskDocsBoilerplate struct {
	conf Config
	docs DocsBoilerplate
}

func NewOnDiskDocsBoilerplate(c Config, d DocsBoilerplate) OnDiskDocsBoilerplate {
	return OnDiskDocsBoilerplate{c, d}
}

func (o OnDiskDocsBoilerplate) Save(f Filesystem) error {
	data, err := o.docs.Content()
	if err != nil {
		return err
	}

	return f.Save(data, o.conf.To)
}

type DocsBoilerplate struct {
	tmpl     InMemoryDocsTemplate
	tmplData DocsTemplateData
}

func NewDocsTemplate(i InMemoryDocsTemplate, d DocsTemplateData) DocsBoilerplate {
	return DocsBoilerplate{i, d}
}

func (d DocsBoilerplate) Content() (string, error) {
	tmplFileContent, err := d.tmpl.Content()
	if err != nil {
		return "", err
	}

	tmpl, err := template.New(d.tmplData.Title).Parse(tmplFileContent)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, d.tmplData); err != nil {
		return "", err
	}

	return result.String(), nil
}

type InMemoryDocsTemplate struct {
	conf       Config
	filesystem Filesystem
}

func NewInMemoryFile(c Config, f Filesystem) InMemoryDocsTemplate {
	return InMemoryDocsTemplate{c, f}
}

func (i InMemoryDocsTemplate) Content() (string, error) {
	return i.filesystem.Load(i.conf.From)
}

type DocsTemplateData struct {
	Title, SmallTitle string
}

func NewDocsTemplateData(t, s string) DocsTemplateData {
	return DocsTemplateData{t, s}
}
