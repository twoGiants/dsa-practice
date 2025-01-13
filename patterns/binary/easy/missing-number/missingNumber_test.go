package binary_test

import (
	binary "dsa/patterns/binary/easy/missing-number"
	"testing"
)

func Test_givenAListOfThreeWithMissingTwo_returnsTwo(t *testing.T) {
	num := binary.MissingNumber([]int{3, 0, 1})
	if num != 2 {
		t.Fatal("expected 2, got", num)
	}
}

func Test_givenAListOfTwoWithMissingTwo_returnsTwo(t *testing.T) {
	num := binary.MissingNumber([]int{0, 1})
	if num != 2 {
		t.Fatal("expected 2, got", num)
	}
}

func Test_givenAListOfNineWithMissingEight_returnsEight(t *testing.T) {
	num := binary.MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	if num != 8 {
		t.Fatal("expected 8, got", num)
	}
}