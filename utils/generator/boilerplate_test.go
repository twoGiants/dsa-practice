package generator_test

import (
	"dsa/utils/generator"
	"os"
	"strings"
	"testing"
)

func Test_GenerateBoilerplateDocs(t *testing.T) {
	docsTests := []struct {
		name, pattern, input, expected string
	}{
		{
			"single word title",
			"Duplicate",
			"{{.Title}} [{{.Title}} solution.]({{.SmallTitle}}-solution.md)",
			"Duplicate [Duplicate solution.](duplicate-solution.md)",
		},
		{
			"two words title",
			"Missing Number",
			"{{.Title}} [{{.Title}} solution.]({{.SmallTitle}}-solution.md)",
			"Missing Number [Missing Number solution.](missing-number-solution.md)",
		},
		{
			"two words title",
			"Missing Number",
			"{{.Title}} [{{.Title}} solution.]({{.SmallTitle}}-solution.md)",
			"Missing Number [Missing Number solution.](missing-number-solution.md)",
		},
		{
			"three words title",
			"Set Matrix Zeroes",
			"{{.Title}} [{{.Title}} solution.]({{.SmallTitle}}-solution.md)",
			"Set Matrix Zeroes [Set Matrix Zeroes solution.](set-matrix-zeroes-solution.md)",
		},
		{
			"two words, one with number title",
			"Convert 1D",
			"{{.Title}} [{{.Title}} solution.]({{.SmallTitle}}-solution.md)",
			"Convert 1D [Convert 1D solution.](convert-1d-solution.md)",
		},
		{
			"three words, one with number title",
			"Convert 1D Array",
			"{{.Title}} [{{.Title}} solution.]({{.SmallTitle}}-solution.md)",
			"Convert 1D Array [Convert 1D Array solution.](convert-1d-array-solution.md)",
		},
	}

	for _, tt := range docsTests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := generator.Docs(tt.input, tt.pattern)
			if err != nil {
				t.Fatal("unexpected error", err)
			}

			if tt.expected != actual {
				t.Fatalf("expected: %s, got: %s", tt.expected, actual)
			}
		})
	}
}

// Integration unit test - not faking fs
func Test_LoadDocsTemplate(t *testing.T) {
	tmplPath := "docs.gotmpl"
	expected := "{{.Title}}"

	actual, err := generator.LoadDocs(tmplPath)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if !strings.Contains(actual, expected) {
		t.Fatalf("expected docs to contain %s but it didn't, docs \n%s", expected, actual)
	}
}

// Integration unit test - not faking fs
func Test_StoreDocsBoilerplate(t *testing.T) {
	path := t.TempDir() + "/docs.md"
	boilerplate := "Missing Number [Missing Number solution.](missing-number-solution.md)"

	t.Log(path)

	err := generator.StoreDocs(boilerplate, path)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	stored, err := generator.LoadDocs(path)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if boilerplate != stored {
		t.Fatalf("expected: %s, got: %s", boilerplate, stored)
	}
}

// Integration unit test - not faking fs
func Test_CreateExerciseDirectories(t *testing.T) {
	path := t.TempDir() + "/arrays/easy/missing-numbers"

	t.Log(path)

	if err := generator.CreateExerciseDirectories(path); err != nil {
		t.Fatal("unexpected error", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if !info.IsDir() {
		t.Fatalf("expected path \"%s\" to exist but it didn't", path)
	}
}

// Integration unit test - not faking fs
func Test_e2e_CreateMissingNumbersBoilerplateMd(t *testing.T) {
	pattern, difficulty, title := "arrays", "easy", "Missing Numbers"
	rootPath := t.TempDir()
	tmplPath := "docs.gotmpl"
	expectedPath := rootPath + "/arrays/easy/missing-numbers/missing-numbers.md"

	if err := generator.DocsBoilerplate(rootPath, tmplPath, pattern, difficulty, title); err != nil {
		t.Fatal("unexpected error", err)
	}

	actual, err := generator.LoadDocs(expectedPath)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	t.Log(actual)

	if !strings.Contains(actual, title) {
		t.Fatalf("expected title \"%s\", got \n%s", title, actual)
	}
}
