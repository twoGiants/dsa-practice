package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/construct-2d-array"
	"slices"
	"testing"
)

func Test_convertArrayOfFourElements_intoATwoByTwoMatrix(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := arrays.Construct2DArray(input, 2, 2)
	expected := [][]int{{1, 2}, {3, 4}}

	if len(result) != 2 {
		t.Fatal("expected length 2, but got", len(result))
	}

	if !slices.Equal(result[0], expected[0]) || !slices.Equal(result[1], expected[1]) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_convertArrayOfThreeElements_intoAOneByThreeMatrix(t *testing.T) {
	input := []int{1, 2, 3}
	result := arrays.Construct2DArray(input, 1, 3)
	expected := [][]int{{1, 2, 3}}

	if len(result) != 1 {
		t.Fatal("expected length 1, but got", len(result))
	}

	if !slices.Equal(result[0], expected[0]) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_doNotConvertArrayOfThree_ifDimensionsAreOneByOne(t *testing.T) {
	input := []int{1, 2}
	result := arrays.Construct2DArray(input, 1, 1)

	if len(result) != 0 {
		t.Fatal("expected length 0, but got", len(result))
	}
}
