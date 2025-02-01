package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/longest-consecutive-sequence"
	"dsa/utils/assert"
	"testing"
)

func Test_arrayOfSeven_hasASequenceOfFour(t *testing.T) {
	input := []int{2, 20, 4, 10, 3, 4, 5}
	expected := 4

	result := arrays.LongestConsecutive(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_arrayOfEight_hasASequenceOfSeven(t *testing.T) {
	input := []int{0, 3, 2, 5, 4, 6, 1, 1}
	expected := 7

	result := arrays.LongestConsecutive(input)

	assert.NumbersEqual(t, result, expected, input)
}
