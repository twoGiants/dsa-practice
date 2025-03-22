package generator

func EoDocs() {
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

	docsBoilerplate.Save(filesystem)
}
