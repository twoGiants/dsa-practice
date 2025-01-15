package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/find-all-numbers-dissappeared"
	"slices"
	"testing"
)

func Test_givenArrayOfTwo_returnsTwoAndThree(t *testing.T) {
	input := []int{1, 1}
	result := arrays.FindDisappearedNumbers(input)
	expected := []int{2}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_givenArrayOfFour_returnsTwoAndThree(t *testing.T) {
	input := []int{1, 4, 4, 4}
	result := arrays.FindDisappearedNumbers(input)
	expected := []int{2, 3}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func Test_givenArrayOfEight_returnsFiveAndSix(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := arrays.FindDisappearedNumbers(input)
	expected := []int{5, 6}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}
