package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/set-matrix-zeroes"
	"slices"
	"testing"
)

func Test_matrixWithOneZeroInFirstRow_firstRowAndColumnBecomesZero(t *testing.T) {
	input := [][]int{{0, 1}, {1, 1}}
	inputCopy := [][]int{{0, 1}, {1, 1}}

	arrays.SetZeroes(input)
	expected := [][]int{{0, 0}, {0, 1}}
	if !slices.Equal(input[0], expected[0]) || !slices.Equal(input[1], expected[1]) {
		t.Fatalf("for input %v expected output %v, but got %v", inputCopy, expected, input)
	}
}

func Test_matrixWithOneZeroInSecondRowAndFirstColumn_secondRowAndSecondColumnBecomesZero(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}}
	inputCopy := [][]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}}

	arrays.SetZeroes(input)
	expected := [][]int{{1, 0, 3}, {0, 0, 0}, {6, 0, 8}}
	if !slices.Equal(input[0], expected[0]) || !slices.Equal(input[1], expected[1]) || !slices.Equal(input[2], expected[2]) {
		t.Fatalf("for input %v expected output %v, but got %v", inputCopy, expected, input)
	}
}
