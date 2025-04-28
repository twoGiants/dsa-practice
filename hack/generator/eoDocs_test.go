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

func Test_DocsTemplateDataSmallKebabTitle(t *testing.T) {
	var testCases = []struct {
		name,
		title,
		expected string
	}{
		{
			name:     "single word title",
			title:    "Arrays",
			expected: "arrays",
		},
		{
			name:     "two words title",
			title:    "Missing Number",
			expected: "missing-number",
		},
		{
			name:     "single word title with number",
			title:    "1D",
			expected: "1d",
		},
		{
			name:     "two words with one number",
			title:    "1D Transformer",
			expected: "1d-transformer",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := generator.NewDocsTemplateData(tc.title)

			if actual.SmallTitle != tc.expected {
				t.Errorf("expected %s, got % s", tc.expected, actual.SmallTitle)
			}
		})
	}
}
