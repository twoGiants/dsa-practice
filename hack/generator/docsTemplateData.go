package generator

import "strings"

type DocsTemplateData struct {
	Title, SmallTitle string
}

func NewDocsTemplateData(title string) DocsTemplateData {
	result := DocsTemplateData{}
	result.Title = title
	result.SmallTitle = smallKebab(title)
	return result
}

func smallKebab(word string) string {
	result := strings.ToLower(word)

	if len(strings.Split(result, " ")) > 1 {
		result = strings.ReplaceAll(result, " ", "-")
	}

	return result
}
