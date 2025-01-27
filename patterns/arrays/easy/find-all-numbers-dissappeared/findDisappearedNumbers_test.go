package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/find-all-numbers-dissappeared"
	"dsa/utils/assert"
	"testing"
)

func Test_givenArrayOfTwo_returnsTwoAndThree(t *testing.T) {
	input := []int{1, 1}
	expected := []int{2}

	result := arrays.FindDisappearedNumbers(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_givenArrayOfFour_returnsTwoAndThree(t *testing.T) {
	input := []int{1, 4, 4, 4}
	expected := []int{2, 3}

	result := arrays.FindDisappearedNumbers(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_givenArrayOfEight_returnsFiveAndSix(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{5, 6}

	result := arrays.FindDisappearedNumbers(input)

	assert.SlicesEqual(t, result, expected, input)
}
