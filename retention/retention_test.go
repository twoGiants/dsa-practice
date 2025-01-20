package retention_test

import (
	"dsa/retention"
	"slices"
	"testing"
)

func Test_FindDuplicates_inArrayOfEightWithTwoDuplicates_findsTwoAndThree(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := retention.FindDuplicates(input)
	expected := []int{2, 3}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_FindDuplicates_inArrayOfThreeWithOneDuplicate_findsOne(t *testing.T) {
	input := []int{1, 1, 2}
	result := retention.FindDuplicates(input)
	expected := []int{1}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_FindDuplicates_inArrayWithOneElement_findsNone(t *testing.T) {
	input := []int{1}
	result := retention.FindDuplicates(input)
	expected := []int{}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_Construct2DArray_convertArrayOfFourElements_intoATwoByTwoMatrix(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := retention.Construct2DArray(input, 2, 2)
	expected := [][]int{{1, 2}, {3, 4}}

	if len(result) != 2 {
		t.Fatal("expected length 2, but got", len(result))
	}

	if !slices.Equal(result[0], expected[0]) || !slices.Equal(result[1], expected[1]) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_Construct2DArray_convertArrayOfThreeElements_intoAOneByThreeMatrix(t *testing.T) {
	input := []int{1, 2, 3}
	result := retention.Construct2DArray(input, 1, 3)
	expected := [][]int{{1, 2, 3}}

	if len(result) != 1 {
		t.Fatal("expected length 1, but got", len(result))
	}

	if !slices.Equal(result[0], expected[0]) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_Construct2DArray_doNotConvertArrayOfThree_ifDimensionsAreOneByOne(t *testing.T) {
	input := []int{1, 2}
	result := retention.Construct2DArray(input, 1, 1)

	if len(result) != 0 {
		t.Fatal("expected length 0, but got", len(result))
	}
}

func Test_SingleNumber_arrayWithOneDoubleAndOneSingle_returnsOne(t *testing.T) {
	input := []int{2, 2, 1}
	result := retention.SingleNumber(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_SingleNumber_arrayWithTwoDoublesAndOneSingle_returnsFour(t *testing.T) {
	input := []int{4, 1, 2, 1, 2}
	result := retention.SingleNumber(input)
	expected := 4
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_SingleNumber_arrayWithOneEntry_returnsThatEntry(t *testing.T) {
	input := []int{1}
	result := retention.SingleNumber(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_FindDisappearedNumbers_arrayOfEight_disappearedAreFiveAndSix(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := retention.FindDisappearedNumbers(input)
	expected := []int{5, 6}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_FindDisappearedNumbers_arrayOfTwo_disappearedIsTwo(t *testing.T) {
	input := []int{1, 1}
	result := retention.FindDisappearedNumbers(input)
	expected := []int{2}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_MissingNumber_arrayOfTree_returnsMissingZero(t *testing.T) {
	input := []int{1, 2, 3}
	result := retention.MissingNumber(input)
	expected := 0
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_MissingNumber_arrayOfTwo_returnsMissingOne(t *testing.T) {
	input := []int{0, 2}
	result := retention.MissingNumber(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_MissingNumberXor_arrayOfTree_returnsMissingZero(t *testing.T) {
	input := []int{1, 2, 3}
	result := retention.MissingNumberXor(input)
	expected := 0
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_MissingNumberXor_arrayOfTwo_returnsMissingOne(t *testing.T) {
	input := []int{0, 2}
	result := retention.MissingNumberXor(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_TwoSum_arrayOfFourWithTargetSeven_returnsZeroAndOne(t *testing.T) {
	input := []int{3, 4, 5, 6}
	target := 7
	result := retention.TwoSum(input, target)
	expected := []int{0, 1}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_TwoSum_arrayOfThreeWithTargetTen_returnsZeroAndTwo(t *testing.T) {
	input := []int{4, 5, 6}
	target := 10
	result := retention.TwoSum(input, target)
	expected := []int{0, 2}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_TwoSum_arrayOfTwoWithTargetTen_returnsZeroAndOne(t *testing.T) {
	input := []int{5, 5}
	target := 10
	result := retention.TwoSum(input, target)
	expected := []int{0, 1}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_HasDuplicate_arrayOfFourWithDuplicate_returnsTrue(t *testing.T) {
	input := []int{1, 2, 3, 3}
	result := retention.HasDuplicate(input)
	if !result {
		t.Fatalf("expected %v to have duplicate, but didn't", input)
	}
}

func Test_HasDuplicate_arrayOfFourWithoutDuplicate_returnsFalse(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := retention.HasDuplicate(input)
	if result {
		t.Fatalf("expected %v to have no duplicate, but did", input)
	}
}

func Test_HasDuplicate_arrayOfOne_returnsFalse(t *testing.T) {
	input := []int{1}
	result := retention.HasDuplicate(input)
	if result {
		t.Fatalf("expected %v to have no duplicate, but did", input)
	}
}
