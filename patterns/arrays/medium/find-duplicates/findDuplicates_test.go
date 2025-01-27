package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/find-duplicates"
	"dsa/utils/assert"
	"testing"
)

func Test_inArrayOfEightWithTwoDuplicates_findsTwoAndThree(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{2, 3}

	result := arrays.FindDuplicates(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_inArrayOfThreeWithOneDuplicate_findsOne(t *testing.T) {
	input := []int{1, 1, 2}
	expected := []int{1}

	result := arrays.FindDuplicates(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_inArrayWithOneElement_findsNone(t *testing.T) {
	input := []int{1}
	expected := []int{}

	result := arrays.FindDuplicates(input)

	assert.SlicesEqual(t, result, expected, input)
}
