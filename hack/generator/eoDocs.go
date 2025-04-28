package generator

func EoDocs(rootDir string, cli InputSource) error {
	request, err := NewRequest(cli)
	if err != nil {
		return err
	}
	filesystem := NewFilesystem()
	tmplSourcePath := "../boilerplate/docs.gotmpl"

	docsBoilerplate := NewDocsBoilerplate(
		request.TargetPath(rootDir),
		NewDocsTemplate(tmplSourcePath, filesystem),
		request.TemplateData(),
	)

	if request.Create() {
		return docsBoilerplate.Save(filesystem)
	}

	return docsBoilerplate.Delete(filesystem)
}
