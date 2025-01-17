package retention_test

import (
	"dsa/retention"
	"slices"
	"testing"
)

func Test_arrayWithOneDoubleAndOneSingle_returnsOne(t *testing.T) {
	input := []int{2, 2, 1}
	result := retention.SingleNumber(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_arrayWithTwoDoublesAndOneSingle_returnsFour(t *testing.T) {
	input := []int{4, 1, 2, 1, 2}
	result := retention.SingleNumber(input)
	expected := 4
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_arrayWithOneEntry_returnsThatEntry(t *testing.T) {
	input := []int{1}
	result := retention.SingleNumber(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_arrayOfEight_disappearedAreFiveAndSix(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := retention.FindDisappearedNumbers(input)
	expected := []int{5, 6}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_arrayOfTwo_disappearedIsTwo(t *testing.T) {
	input := []int{1, 1}
	result := retention.FindDisappearedNumbers(input)
	expected := []int{2}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_arrayOfTree_returnsMissingZero(t *testing.T) {
	input := []int{1, 2, 3}
	result := retention.MissingNumber(input)
	expected := 0
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_arrayOfTwo_returnsMissingOne(t *testing.T) {
	input := []int{0, 2}
	result := retention.MissingNumber(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func TestXor_arrayOfTree_returnsMissingZero(t *testing.T) {
	input := []int{1, 2, 3}
	result := retention.MissingNumberXor(input)
	expected := 0
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func TestXor_arrayOfTwo_returnsMissingOne(t *testing.T) {
	input := []int{0, 2}
	result := retention.MissingNumberXor(input)
	expected := 1
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func Test_arrayOfFourWithTargetSeven_returnsZeroAndOne(t *testing.T) {
	input := []int{3, 4, 5, 6}
	target := 7
	result := retention.TwoSum(input, target)
	expected := []int{0, 1}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_arrayOfThreeWithTargetTen_returnsZeroAndTwo(t *testing.T) {
	input := []int{4, 5, 6}
	target := 10
	result := retention.TwoSum(input, target)
	expected := []int{0, 2}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_arrayOfTwoWithTargetTen_returnsZeroAndOne(t *testing.T) {
	input := []int{5, 5}
	target := 10
	result := retention.TwoSum(input, target)
	expected := []int{0, 1}
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected %v, but got %v", input, expected, result)
	}
}

func Test_arrayOfFourWithDuplicate_returnsTrue(t *testing.T) {
	input := []int{1, 2, 3, 3}
	result := retention.HasDuplicate(input)
	if !result {
		t.Fatalf("expected %v to have duplicate, but didn't", input)
	}
}

func Test_arrayOfFourWithoutDuplicate_returnsFalse(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := retention.HasDuplicate(input)
	if result {
		t.Fatalf("expected %v to have no duplicate, but did", input)
	}
}

func Test_arrayOfOne_returnsFalse(t *testing.T) {
	input := []int{1}
	result := retention.HasDuplicate(input)
	if result {
		t.Fatalf("expected %v to have no duplicate, but did", input)
	}
}
