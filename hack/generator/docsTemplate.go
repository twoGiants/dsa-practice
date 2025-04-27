package generator

type DocsTemplate struct {
	conf  Config
	store Store
}

func NewDocsTemplate(c Config, store Store) DocsTemplate {
	return DocsTemplate{c, store}
}

func (dt DocsTemplate) Content() (string, error) {
	return dt.store.Load(dt.conf.From)
}
