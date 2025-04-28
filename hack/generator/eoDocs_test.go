package generator_test

import (
	"dsa/hack/generator"
	"testing"
)

func Test_CreateAndDeleteDocsBoilerplateInTemp(t *testing.T) {
	targetPath := t.TempDir()

	if err := generator.EoDocsCreate(targetPath); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := generator.EoDocsDelete(targetPath); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_DocsTemplateDataHasSmallKebabCaseTitle_TwoWords(t *testing.T) {
	expected := "missing-number"

	actual := generator.NewDocsTemplateData("Missing Number")

	if actual.SmallTitle != expected {
		t.Errorf("expected %s, got % s", expected, actual.SmallTitle)
	}
}

func Test_DocsTemplateDataHasSmallKebabCaseTitle_OneWord(t *testing.T) {
	expected := "arrays"

	actual := generator.NewDocsTemplateData("Arrays")

	if actual.SmallTitle != expected {
		t.Errorf("expected %s, got % s", expected, actual.SmallTitle)
	}
}
