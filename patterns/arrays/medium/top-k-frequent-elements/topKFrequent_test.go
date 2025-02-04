package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/top-k-frequent-elements"
	"dsa/utils/assert"
	"testing"
)

func Test_TopKFrequent(t *testing.T) {
	var topKTests = []struct {
		name     string
		inputArr []int
		inputK   int
		expected []int
	}{
		{"array of six with k=2", []int{1, 2, 2, 3, 3, 3}, 2, []int{3, 2}},
		{"array of two with k=1", []int{7, 7}, 1, []int{7}},
	}

	for _, tt := range topKTests {
		t.Run(tt.name, func(t *testing.T) {

			result := arrays.TopKFrequent(tt.inputArr, tt.inputK)

			assert.SlicesEqual(t, result, tt.expected, tt.inputArr)
		})
	}
}
