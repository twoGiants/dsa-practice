package generator

type DocsTemplateData struct {
	Title, SmallTitle string
}

func NewDocsTemplateData(t, s string) DocsTemplateData {
	return DocsTemplateData{t, s}
}
