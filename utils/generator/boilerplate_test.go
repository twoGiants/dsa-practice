package generator_test

import (
	"dsa/utils/generator"
	"testing"
)

func Test_GenerateBoilerplateDocs(t *testing.T) {
	var docsTests = []struct {
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
