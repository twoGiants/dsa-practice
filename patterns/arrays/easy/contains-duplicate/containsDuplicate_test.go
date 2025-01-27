package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/contains-duplicate"
	"dsa/utils/assert"
	"testing"
)

func Test_givenEmptyArray_isFalse(t *testing.T) {
	input := []int{}

	result := arrays.HasDuplicate(input)

	assert.NoDuplicates(t, result, input)
}

func Test_givenArrayOfOneElement_isFalse(t *testing.T) {
	input := []int{0}

	result := arrays.HasDuplicate(input)

	assert.NoDuplicates(t, result, input)
}

func Test_givenArrayOfTwoIdenticalElements_isTrue(t *testing.T) {
	input := []int{0, 0}

	result := arrays.HasDuplicate(input)

	assert.Duplicates(t, result, input)
}

func Test_givenArrayOfTwoUniqueElements_isFalse(t *testing.T) {
	input := []int{0, 1}

	result := arrays.HasDuplicate(input)

	assert.NoDuplicates(t, result, input)
}

func Test_givenArrayOfThreeUniqueElements_isFalse(t *testing.T) {
	input := []int{0, 1, 2}

	result := arrays.HasDuplicate(input)

	assert.NoDuplicates(t, result, input)
}

func Test_givenArrayOfThreeElementsWithTwoIdentical_isTrue(t *testing.T) {
	input := []int{1, 1, 2}

	result := arrays.HasDuplicate(input)

	assert.Duplicates(t, result, input)
}

func Test_givenArrayOfFourUniqueElements_isFalse(t *testing.T) {
	input := []int{1, 2, 3, 4}

	result := arrays.HasDuplicate(input)

	assert.NoDuplicates(t, result, input)
}

func Test_givenArrayOfFourWithOneDuplicate_isTrue(t *testing.T) {
	input := []int{1, 2, 3, 3}

	result := arrays.HasDuplicate(input)

	assert.Duplicates(t, result, input)
}
