package bit_test

import (
	bit "dsa/patterns/bit-manipulation/easy/missing-number"
	"dsa/utils/assert"
	"testing"
)

func Test_givenAListOfThreeWithMissingTwo_returnsTwo(t *testing.T) {
	input := []int{3, 0, 1}
	expected := 2

	result := bit.MissingNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_givenAListOfTwoWithMissingTwo_returnsTwo(t *testing.T) {
	input := []int{0, 1}
	expected := 2

	result := bit.MissingNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_givenAListOfNineWithMissingEight_returnsEight(t *testing.T) {
	input := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expected := 8

	result := bit.MissingNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func TestXor_givenAListOfThreeWithMissingTwo_returnsTwo(t *testing.T) {
	input := []int{3, 0, 1}
	expected := 2

	result := bit.MissingNumberXor(input)

	assert.NumbersEqual(t, result, expected, input)
}

func TestXor_givenAListOfTwoWithMissingTwo_returnsTwo(t *testing.T) {
	input := []int{0, 1}
	expected := 2

	result := bit.MissingNumberXor(input)

	assert.NumbersEqual(t, result, expected, input)
}

func TestXor_givenAListOfNineWithMissingEight_returnsEight(t *testing.T) {
	input := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expected := 8

	result := bit.MissingNumberXor(input)

	assert.NumbersEqual(t, result, expected, input)
}
