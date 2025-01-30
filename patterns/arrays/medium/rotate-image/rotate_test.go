package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/rotate-image"
	"slices"
	"testing"
)

func Test_Rotate(t *testing.T) {
	input := [][]int{}
	copy := [][]int{}
	expected := [][]int{}

	arrays.Rotate(input)

	if !slices.Equal(input[0], expected[0]) {
		t.Fatalf("for input %v expected output %v, but got %v", copy, expected, input)
	}
}
