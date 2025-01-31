package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/rotate-image"
	"dsa/utils/assert"
	"testing"
)

func Test_RotateTwoByTwoMatrix(t *testing.T) {
	input := [][]int{
		{1, 2},
		{3, 4},
	}
	copy := [][]int{
		{1, 2},
		{3, 4},
	}
	expected := [][]int{
		{3, 1},
		{4, 2},
	}

	arrays.Rotate(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_RotateThreeByThreeMatrix(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	copy := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	arrays.Rotate(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_RotateFourByFourMatrix(t *testing.T) {
	input := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	copy := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	expected := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	arrays.Rotate(input)

	assert.MatrixEqual(t, input, expected, copy)
}
