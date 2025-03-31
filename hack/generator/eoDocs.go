package generator

func EoDocs() error {
	filesystem := NewFilesystem()
	config := Config{
		To:   "temp",
		From: "hack/boilerplate/docs.gotmpl",
	}

	docsBoilerplate := NewDocsBoilerplate(
		config,
		NewDocsTemplate(config, filesystem),
		NewDocsTemplateData("Missing Number", "missing-number"),
	)

	return docsBoilerplate.Save(filesystem)
}
