package bit_test

import (
	bit "dsa/patterns/bit-manipulation/easy/single-number"
	"testing"
)

func Test_givenAnArrayWithOneDuplicate_returnOne(t *testing.T) {
	input := []int{2, 2, 1}
	result := bit.SingleNumber(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected output %d, but got %d", input, expected, result)
	}
}

func Test_givenAnArrayWithTwoDuplicates_returnFour(t *testing.T) {
	input := []int{4, 1, 2, 1, 2}
	result := bit.SingleNumber(input)
	expected := 4
	if result != expected {
		t.Fatalf("for input %v expected output %d, but got %d", input, expected, result)
	}
}

func Test_givenAnArrayWithOneDigit_returnThatDigit(t *testing.T) {
	input := []int{1}
	result := bit.SingleNumber(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected output %d, but got %d", input, expected, result)
	}
}
