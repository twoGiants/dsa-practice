package generator

type DocsTemplate struct {
	from  string
	store Store
}

func NewDocsTemplate(path string, store Store) DocsTemplate {
	return DocsTemplate{path, store}
}

func (dt DocsTemplate) Content() (string, error) {
	return dt.store.Load(dt.from)
}
