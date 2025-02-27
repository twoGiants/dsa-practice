package generator

import (
	"strings"
	"text/template"
)

type Boilerplate struct {
	tmpl  string
	title string
	value string
}

func (b *Boilerplate) StrValue() (string, error) {
	if err := b.generate(); err != nil {
		return "", err
	}

	return b.value, nil
}

func (b *Boilerplate) generate() error {
	tmpl, err := template.New(b.title).Parse(b.tmpl)
	if err != nil {
		return err
	}

	small := b.smallKebabTitle()

	data := struct {
		Title      string
		SmallTitle string
	}{
		Title:      b.title,
		SmallTitle: small,
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, data); err != nil {
		return err
	}

	b.value = result.String()

	return nil
}

func (b *Boilerplate) smallKebabTitle() string {
	small := strings.ToLower(b.title)

	if len(strings.Split(small, " ")) > 1 {
		small = strings.ReplaceAll(small, " ", "-")
	}

	return small
}
