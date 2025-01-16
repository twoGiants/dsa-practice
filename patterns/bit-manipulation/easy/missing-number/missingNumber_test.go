package bit_test

import (
	bit "dsa/patterns/bit-manipulation/easy/missing-number"
	"testing"
)

func Test_givenAListOfThreeWithMissingTwo_returnsTwo(t *testing.T) {
	num := bit.MissingNumber([]int{3, 0, 1})
	if num != 2 {
		t.Fatal("expected 2, got", num)
	}
}

func Test_givenAListOfTwoWithMissingTwo_returnsTwo(t *testing.T) {
	num := bit.MissingNumber([]int{0, 1})
	if num != 2 {
		t.Fatal("expected 2, got", num)
	}
}

func Test_givenAListOfNineWithMissingEight_returnsEight(t *testing.T) {
	num := bit.MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	if num != 8 {
		t.Fatal("expected 8, got", num)
	}
}

func TestXor_givenAListOfThreeWithMissingTwo_returnsTwo(t *testing.T) {
	num := bit.MissingNumberXor([]int{3, 0, 1})
	if num != 2 {
		t.Fatal("expected 2, got", num)
	}
}

func TestXor_givenAListOfTwoWithMissingTwo_returnsTwo(t *testing.T) {
	num := bit.MissingNumberXor([]int{0, 1})
	if num != 2 {
		t.Fatal("expected 2, got", num)
	}
}

func TestXor_givenAListOfNineWithMissingEight_returnsEight(t *testing.T) {
	num := bit.MissingNumberXor([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	if num != 8 {
		t.Fatal("expected 8, got", num)
	}
}
