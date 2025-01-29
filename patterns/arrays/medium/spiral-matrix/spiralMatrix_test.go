package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/spiral-matrix"
	"dsa/utils/assert"
	"testing"
)

func Test_twoByTwoMatrix_spiralStartsWithOneAndEndsWithFive(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

	result := arrays.SpiralOrder(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_threeByFourMatrix_spiralStartsWithOneAndEndsWithSeven(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	result := arrays.SpiralOrder(input)

	assert.SlicesEqual(t, result, expected, input)
}
