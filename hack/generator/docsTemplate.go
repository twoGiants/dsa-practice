package generator

type DocsTemplate struct {
	conf Config
	sl   StorageLocation
}

func NewDocsTemplate(conf Config, sl StorageLocation) DocsTemplate {
	return DocsTemplate{conf, sl}
}

func (dt DocsTemplate) Content() (string, error) {
	return dt.sl.Load(dt.conf.From)
}
