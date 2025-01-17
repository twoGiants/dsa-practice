package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/find-duplicates"
	"slices"
	"testing"
)

func Test_inArrayOfEightWithTwoDuplicates_findsTwoAndThree(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := arrays.FindDuplicates(input)
	expected := []int{2, 3}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}
func Test_inArrayOfThreeWithOneDuplicate_findsOne(t *testing.T) {
	input := []int{1, 1, 2}
	result := arrays.FindDuplicates(input)
	expected := []int{1}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}
func Test_inArrayWithOneElement_findsNone(t *testing.T) {
	input := []int{1}
	result := arrays.FindDuplicates(input)
	expected := []int{}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}
