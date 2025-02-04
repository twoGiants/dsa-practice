package retention_test

import (
	"dsa/retention"
	"dsa/utils/assert"
	"slices"
	"testing"
)

func Test_GroupAnagrams(t *testing.T) {
	var anagramTests = []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			"six words",
			[]string{"act", "pots", "tops", "cat", "stop", "hat"},
			[][]string{{"hat"}, {"act", "cat"}, {"pots", "stop", "tops"}},
		},
		{"x", []string{"x"}, [][]string{{"x"}}},
		{"", []string{""}, [][]string{{""}}},
	}

	for _, at := range anagramTests {
		t.Run(at.name, func(t *testing.T) {

			result := retention.GroupAnagrams(at.input)

			assert.SortGroups(result)

			for i := 0; i < len(at.expected); i++ {
				if !slices.Equal(result[i], at.expected[i]) {
					t.Errorf("for input %v expected output %v, but got %v", at.input, at.expected[i], result[i])
				}
			}
		})
	}
}

func Test_ValidAnagram_IsAnagram(t *testing.T) {
	var anagramTests = []struct {
		s1, s2   string
		expected bool
	}{
		{"racecar", "carrace", true},
		{"jar", "jam", false},
	}

	for _, at := range anagramTests {
		t.Run(at.s1+"."+at.s2, func(t *testing.T) {
			result := retention.IsAnagram(at.s1, at.s2)

			if result != at.expected {
				t.Errorf("for input '%s', '%s', expected output %t, but got %t", at.s1, at.s2, at.expected, result)
			}
		})
	}
}

func Test_LongestConsecutive_arrayOfEight_longestSequenceIsSeven(t *testing.T) {
	input := []int{0, 3, 2, 5, 4, 6, 1, 1}
	expected := 7

	result := retention.LongestConsecutive(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_LongestConsecutive_arrayOfSeven_longestSequenceIsFour(t *testing.T) {
	input := []int{2, 20, 4, 10, 3, 4, 5}
	expected := 4

	result := retention.LongestConsecutive(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_SpiralOrder_oneByThreeMatrix_spiralIsThoseThreeNumbers(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
	}
	expected := []int{1, 2, 3}

	result := retention.SpiralOrder(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_SpiralOrder_threeByOneMatrix_spiralIsThoseThreeNumbers(t *testing.T) {
	input := [][]int{
		{1},
		{2},
		{3},
	}
	expected := []int{1, 2, 3}

	result := retention.SpiralOrder(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_SpiralOrder_twoByTwoMatrix_spiralStartsWithOneAndEndsWithFive(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

	result := retention.SpiralOrder(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_SpiralOrder_threeByFourMatrix_spiralStartsWithOneAndEndsWithSeven(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	result := retention.SpiralOrder(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_RotateImage_twoByTwoMatrix(t *testing.T) {
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

	retention.RotateImage(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_RotateImage_threeByThreeMatrix(t *testing.T) {
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

	retention.RotateImage(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_RotateImage_fourByFourMatrix(t *testing.T) {
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

	retention.RotateImage(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_SetZeroes_matrixWithOneZerosInLeftAndRightTopCorners_firstRowAndColumnAndLastColumnZeroes(t *testing.T) {
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

	retention.SetZeroes(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_SetZeroes_matrixWithOneZeroInTheMiddle_secondRowAndColumnZeroes(t *testing.T) {
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

	retention.SetZeroes(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_SetZeroes_matrixWithOneZeroInRowAndColumnOne_firstRowAndColumnZeroes(t *testing.T) {
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

	retention.SetZeroes(input)

	assert.MatrixEqual(t, input, expected, copy)
}

func Test_ProductExceptSelf_arrayOfFiveElements_produces_0_m6_0_0_0(t *testing.T) {
	input := []int{-1, 0, 1, 2, 3}
	expected := []int{0, -6, 0, 0, 0}

	result := retention.ProductExceptSelf(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_ProductExceptSelf_arrayOfFourElements_produces_48_24_12_8(t *testing.T) {
	input := []int{1, 2, 4, 6}
	expected := []int{48, 24, 12, 8}

	result := retention.ProductExceptSelf(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_FindDuplicates_inArrayOfEightWithTwoDuplicates_findsTwoAndThree(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{2, 3}

	result := retention.FindDuplicates(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_FindDuplicates_inArrayOfThreeWithOneDuplicate_findsOne(t *testing.T) {
	input := []int{1, 1, 2}
	expected := []int{1}

	result := retention.FindDuplicates(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_FindDuplicates_inArrayWithOneElement_findsNone(t *testing.T) {
	input := []int{1}
	expected := []int{}

	result := retention.FindDuplicates(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_Construct2DArray_convertArrayOfFourElements_intoATwoByTwoMatrix(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := [][]int{{1, 2}, {3, 4}}

	result := retention.Construct2DArray(input, 2, 2)

	assert.SliceLengthEqual(t, result, expected)

	if !slices.Equal(result[0], expected[0]) ||
		!slices.Equal(result[1], expected[1]) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_Construct2DArray_convertArrayOfThreeElements_intoAOneByThreeMatrix(t *testing.T) {
	input := []int{1, 2, 3}
	expected := [][]int{{1, 2, 3}}

	result := retention.Construct2DArray(input, 1, 3)

	assert.SliceLengthEqual(t, result, expected)

	assert.SlicesEqual(t, result[0], expected[0], input)
}

func Test_Construct2DArray_doNotConvertArrayOfThree_ifDimensionsAreOneByOne(t *testing.T) {
	input := []int{1, 2}
	expected := [][]int{}

	result := retention.Construct2DArray(input, 1, 1)

	assert.SliceLengthEqual(t, result, expected)
}

func Test_FindDisappearedNumbers_arrayOfEight_disappearedAreFiveAndSix(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{5, 6}

	result := retention.FindDisappearedNumbers(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_FindDisappearedNumbers_arrayOfTwo_disappearedIsTwo(t *testing.T) {
	input := []int{1, 1}
	expected := []int{2}

	result := retention.FindDisappearedNumbers(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_TwoSum_arrayOfFourWithTargetSeven_returnsZeroAndOne(t *testing.T) {
	input := []int{3, 4, 5, 6}
	target := 7
	expected := []int{0, 1}

	result := retention.TwoSum(input, target)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_TwoSum_arrayOfThreeWithTargetTen_returnsZeroAndTwo(t *testing.T) {
	input := []int{4, 5, 6}
	target := 10
	expected := []int{0, 2}

	result := retention.TwoSum(input, target)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_TwoSum_arrayOfTwoWithTargetTen_returnsZeroAndOne(t *testing.T) {
	input := []int{5, 5}
	target := 10
	expected := []int{0, 1}

	result := retention.TwoSum(input, target)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_HasDuplicate_arrayOfFourWithDuplicate_returnsTrue(t *testing.T) {
	input := []int{1, 2, 3, 3}

	result := retention.HasDuplicate(input)

	assert.Duplicates(t, result, input)
}

func Test_HasDuplicate_arrayOfFourWithoutDuplicate_returnsFalse(t *testing.T) {
	input := []int{1, 2, 3, 4}

	result := retention.HasDuplicate(input)

	assert.NoDuplicates(t, result, input)
}

func Test_HasDuplicate_arrayOfOne_returnsFalse(t *testing.T) {
	input := []int{1}

	result := retention.HasDuplicate(input)

	assert.NoDuplicates(t, result, input)
}
