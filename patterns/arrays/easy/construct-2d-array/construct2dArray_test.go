package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/construct-2d-array"
	"dsa/utils/assert"
	"slices"
	"testing"
)

func Test_convertArrayOfFourElements_intoATwoByTwoMatrix(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := [][]int{{1, 2}, {3, 4}}

	result := arrays.Construct2DArray(input, 2, 2)

	assert.FirstDimensionLengthEqual(t, result, expected)

	if !slices.Equal(result[0], expected[0]) || !slices.Equal(result[1], expected[1]) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_convertArrayOfThreeElements_intoAOneByThreeMatrix(t *testing.T) {
	input := []int{1, 2, 3}
	expected := [][]int{{1, 2, 3}}

	result := arrays.Construct2DArray(input, 1, 3)

	assert.FirstDimensionLengthEqual(t, result, expected)

	if !slices.Equal(result[0], expected[0]) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_doNotConvertArrayOfThree_ifDimensionsAreOneByOne(t *testing.T) {
	input := []int{1, 2}
	expected := [][]int{}

	result := arrays.Construct2DArray(input, 1, 1)

	assert.FirstDimensionLengthEqual(t, result, expected)
}
