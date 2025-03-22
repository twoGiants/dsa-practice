package generator

type DocsTemplate struct {
	conf Config
	fs   Filesystem
}

func NewDocsTemplate(conf Config, fs Filesystem) DocsTemplate {
	return DocsTemplate{conf, fs}
}

func (dt DocsTemplate) Content() (string, error) {
	return dt.fs.Load(dt.conf.From)
}
