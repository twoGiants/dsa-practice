package generator

func EoDocsCreate() error {
	filesystem := NewFilesystem()
	config := Config{
		To:   "../../temp",
		From: "../boilerplate/docs.gotmpl",
	}

	docsBoilerplate := NewDocsBoilerplate(
		config,
		NewDocsTemplate(config, filesystem),
		NewDocsTemplateData("Missing Number", "missing-number"),
	)

	return docsBoilerplate.Save(filesystem)
}
