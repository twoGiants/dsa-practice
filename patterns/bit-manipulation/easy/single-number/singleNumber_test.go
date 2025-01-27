package bit_test

import (
	bit "dsa/patterns/bit-manipulation/easy/single-number"
	"dsa/utils/assert"
	"testing"
)

func Test_givenAnArrayWithOneDuplicate_returnOne(t *testing.T) {
	input := []int{2, 2, 1}
	expected := 1

	result := bit.SingleNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_givenAnArrayWithTwoDuplicates_returnFour(t *testing.T) {
	input := []int{4, 1, 2, 1, 2}
	expected := 4

	result := bit.SingleNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_givenAnArrayWithOneDigit_returnThatDigit(t *testing.T) {
	input := []int{1}
	expected := 1

	result := bit.SingleNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}
