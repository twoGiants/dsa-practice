package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/top-k-frequent-elements"
	"dsa/utils/assert"
	"slices"
	"testing"
)

func Test_TopKFrequent(t *testing.T) {
	var topKTests = []struct {
		name     string
		inputArr []int
		inputK   int
		expected []int
	}{
		{"array [1,2,2,3,3,3] with k=2", []int{1, 2, 2, 3, 3, 3}, 2, []int{3, 2}},
		{"array [1,1,1,2,2,3] with k=2", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"array [7,7] with k=1", []int{7, 7}, 1, []int{7}},
		{"array [1] with k=1", []int{1}, 1, []int{1}},
	}

	for _, tt := range topKTests {
		t.Run(tt.name, func(t *testing.T) {

			result := arrays.TopKFrequent(tt.inputArr, tt.inputK)

			assert.SlicesEqual(t, result, tt.expected, tt.inputArr)
		})
	}
}

func Test_TopKFrequent_nonDeterministic(t *testing.T) {
	inArr := []int{1, 2, 3}
	inK := 1
	expected := []int{1, 2, 3}

	result := arrays.TopKFrequent(inArr, inK)

	if !slices.Contains(expected, result[0]) {
		t.Errorf("expected one of %v, got %d", expected, result[0])
	}
}
