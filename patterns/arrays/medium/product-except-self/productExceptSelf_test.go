package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/product-except-self"
	"dsa/utils/assert"
	"testing"
)

func Test_givenArrayOfFourIntegers_return_24_12_8_6(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	result := arrays.ProductExceptSelf(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_givenArrayOfFourIntegers_return_48_24_12_8(t *testing.T) {
	input := []int{1, 2, 4, 6}
	expected := []int{48, 24, 12, 8}

	result := arrays.ProductExceptSelf(input)

	assert.SlicesEqual(t, result, expected, input)
}

func Test_givenArrayOfFiveIntegers_returnZerosAndOneSix(t *testing.T) {
	input := []int{-1, 0, 1, 2, 3}
	expected := []int{0, -6, 0, 0, 0}

	result := arrays.ProductExceptSelf(input)

	assert.SlicesEqual(t, result, expected, input)
}
