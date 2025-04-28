package generator

type DocsTemplate struct {
	sourcePath string
	store      Store
}

func NewDocsTemplate(sourcePath string, store Store) DocsTemplate {
	return DocsTemplate{sourcePath, store}
}

func (dt DocsTemplate) Content() (string, error) {
	return dt.store.Load(dt.sourcePath)
}
