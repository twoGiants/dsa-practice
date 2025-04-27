package generator_test

import (
	"dsa/hack/generator"
	"testing"
)

func Test_CreateDocsBoilerplateInTemp(t *testing.T) {
	if err := generator.EoDocsCreate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_DeleteDocsBoilerplateFromTemp(t *testing.T) {
	if err := generator.EoDocsDelete(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
