package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/contains-duplicate"
	"testing"
)

func Test_givenEmptyArray_isFalse(t *testing.T) {
	input := []int{}
	if arrays.HasDuplicate(input) {
		t.Fatalf("expected %v to have no duplicates, but got some", input)
	}
}

func Test_givenArrayOfOneElement_isFalse(t *testing.T) {
	input := []int{0}
	if arrays.HasDuplicate(input) {
		t.Fatalf("expected %v to have no duplicates, but got some", input)
	}
}

func Test_givenArrayOfTwoIdenticalElements_isTrue(t *testing.T) {
	input := []int{0, 0}
	if !arrays.HasDuplicate(input) {
		t.Fatalf("expected %v to have no duplicates, but got some", input)
	}
}

func Test_givenArrayOfTwoUniqueElements_isFalse(t *testing.T) {
	input := []int{0, 1}
	if arrays.HasDuplicate(input) {
		t.Fatalf("expected %v to have no duplicates, but got some", input)
	}
}

func Test_givenArrayOfThreeUniqueElements_isFalse(t *testing.T) {
	input := []int{0, 1, 2}
	if arrays.HasDuplicate(input) {
		t.Fatalf("expected %v to have no duplicates, but got some", input)
	}
}

func Test_givenArrayOfThreeElementsWithTwoIdentical_isTrue(t *testing.T) {
	input := []int{1, 1, 2}
	if !arrays.HasDuplicate(input) {
		t.Fatalf("expected %v to have duplicates, but got none", input)
	}
}

func Test_givenArrayOfFourUniqueElements_isFalse(t *testing.T) {
	input := []int{1, 2, 3, 4}
	if arrays.HasDuplicate(input) {
		t.Fatalf("expected %v to have no duplicates, but got some", input)
	}
}

func Test_givenArrayOfFourWithOneDuplicate_isTrue(t *testing.T) {
	input := []int{1, 2, 3, 3}
	if !arrays.HasDuplicate(input) {
		t.Fatalf("expected %v to have no duplicates, but got some", input)
	}
}
