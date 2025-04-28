package generator

func EoDocsCreate(targetPath string) error {
	filesystem := NewFilesystem()
	config := Config{
		To:   targetPath,
		From: "../boilerplate/docs.gotmpl",
	}

	docsBoilerplate := NewDocsBoilerplate(
		config.To,
		NewDocsTemplate(config.From, filesystem),
		NewDocsTemplateData("Missing Number", "missing-number"),
	)

	return docsBoilerplate.Save(filesystem)
}

func EoDocsDelete(targetPath string) error {
	filesystem := NewFilesystem()
	config := Config{
		To:   targetPath,
		From: "../boilerplate/docs.gotmpl",
	}

	docsBoilerplate := NewDocsBoilerplate(
		config.To,
		NewDocsTemplate(config.From, filesystem),
		NewDocsTemplateData("Missing Number", "missing-number"),
	)

	return docsBoilerplate.Delete(filesystem)
}
