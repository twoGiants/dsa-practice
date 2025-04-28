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
