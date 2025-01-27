package retention_test

import (
	"dsa/retention"
	"dsa/utils/assert"
	"testing"
)

func Test_SingleNumber_arrayWithOneDoubleAndOneSingle_returnsOne(t *testing.T) {
	input := []int{2, 2, 1}
	expected := 1

	result := retention.SingleNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_SingleNumber_arrayWithTwoDoublesAndOneSingle_returnsFour(t *testing.T) {
	input := []int{4, 1, 2, 1, 2}
	expected := 4

	result := retention.SingleNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_SingleNumber_arrayWithOneEntry_returnsThatEntry(t *testing.T) {
	input := []int{1}
	expected := 1

	result := retention.SingleNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_MissingNumber_arrayOfTree_returnsMissingZero(t *testing.T) {
	input := []int{1, 2, 3}
	expected := 0

	result := retention.MissingNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_MissingNumber_arrayOfTwo_returnsMissingOne(t *testing.T) {
	input := []int{0, 2}
	expected := 1

	result := retention.MissingNumber(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_MissingNumberXor_arrayOfTree_returnsMissingZero(t *testing.T) {
	input := []int{1, 2, 3}
	expected := 0

	result := retention.MissingNumberXor(input)

	assert.NumbersEqual(t, result, expected, input)
}

func Test_MissingNumberXor_arrayOfTwo_returnsMissingOne(t *testing.T) {
	input := []int{0, 2}
	expected := 1

	result := retention.MissingNumberXor(input)

	assert.NumbersEqual(t, result, expected, input)
}
