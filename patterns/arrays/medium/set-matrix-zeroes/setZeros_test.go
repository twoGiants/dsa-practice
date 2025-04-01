package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/set-matrix-zeroes"
	"dsa/utils/assert"
	"testing"
)

func Test_twoByTwoMatrixWithZeroInFirstRowFirstColumn_firstRowAndColumnBecomesZero(t *testing.T) {
	input := [][]int{
		{0, 1},
		{1, 1},
	}
	copy := [][]int{
		{0, 1},
		{1, 1},
	}
	expected := [][]int{
		{0, 0},
		{0, 1},
	}

	arrays.SetZeroes(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_threeByThreeMatrixWithOneZeroInTheMiddle_secondRowAndSecondColumnBecomesZero(
	t *testing.T,
) {
	input := [][]int{
		{1, 2, 3},
		{4, 0, 5},
		{6, 7, 8},
	}
	copy := [][]int{
		{1, 2, 3},
		{4, 0, 5},
		{6, 7, 8},
	}
	expected := [][]int{
		{1, 0, 3},
		{0, 0, 0},
		{6, 0, 8},
	}

	arrays.SetZeroes(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_threeByThreeMatrixFilledWithOnesAndOneZeroInTheMiddle_secondRowAndSecondColumnBecomesZero(
	t *testing.T,
) {
	input := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	copy := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	expected := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}

	arrays.SetZeroes(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_threeByFourMatrixWithTwoZerosInTheFirstRow_firstRowFirstColumnAndLastColumnBecomesZero(
	t *testing.T,
) {
	input := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	copy := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0},
	}

	arrays.SetZeroes(input)

	assert.MatrixEqual(t, input, expected, copy)
}
