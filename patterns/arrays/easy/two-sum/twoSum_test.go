package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/two-sum"
	"dsa/utils/assert"
	"testing"
)

func Test_arrayOfTwoElements_returnsTheirIndex(t *testing.T) {
	input := []int{5, 5}
	target := 10
	expected := []int{0, 1}

	result := arrays.TwoSum(input, target)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_arrayOfThreeElements_returnsFirstAndLastIndex(t *testing.T) {
	input := []int{4, 5, 6}
	target := 10
	expected := []int{0, 2}

	result := arrays.TwoSum(input, target)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_arrayOfThreeElements_returnsSecondAndThirdIndex(t *testing.T) {
	input := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}

	result := arrays.TwoSum(input, target)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_arrayOfFourElements_returnsFirstAndSecondIndex(t *testing.T) {
	input := []int{3, 4, 5, 6}
	target := 7
	expected := []int{0, 1}

	result := arrays.TwoSum(input, target)

	assert.SlicesEqual(t, result, expected, input)
}
