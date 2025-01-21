package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/product-except-self"
	"slices"
	"testing"
)

func Test_givenArrayOfFourIntegers_return_48_24_12_8(t *testing.T) {
	input := []int{1, 2, 4, 6}
	result := arrays.ProductExceptSelf(input)
	expected := []int{48, 24, 12, 8}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_givenArrayOfFiveIntegers_returnZerosAndOneSix(t *testing.T) {
	input := []int{-1, 0, 1, 2, 3}
	result := arrays.ProductExceptSelf(input)
	expected := []int{0, -6, 0, 0, 0}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}
