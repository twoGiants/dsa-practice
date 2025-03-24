package assert_test

import (
	"dsa/utils/assert"
	"testing"
)

func Test_FirstDimensionLengthEqual(t *testing.T) {
	mockT := new(testing.T)

	var tests = []struct {
		name     string
		actual   [][]int
		expected [][]int
		result   bool
	}{
		{"two 2D slices of equal length", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := assert.FirstDimensionLengthEqual(mockT, tt.actual, tt.expected)

			if result != tt.result {
				t.Errorf("expected %t, got %t", tt.result, result)
			}
		})
	}
}
