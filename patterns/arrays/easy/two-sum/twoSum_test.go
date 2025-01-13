package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/two-sum"
	"testing"
)

func Test_arrayOfTwoElements_returnsTheirIndex(t *testing.T) {
	indexes := arrays.TwoSum([]int{5, 5}, 10)
	if indexes[0] != 0 || indexes[1] != 1 {
		t.Fatal("expected (0, 1), but got", indexes)
	}
}

func Test_arrayOfThreeElements_returnsFirstAndLastIndex(t *testing.T) {
	indexes := arrays.TwoSum([]int{4, 5, 6}, 10)
	if indexes[0] != 0 || indexes[1] != 2 {
		t.Fatal("expected (0, 2), but got", indexes)
	}
}

func Test_arrayOfThreeElements_returnsSecondAndThirdIndex(t *testing.T) {
	indexes := arrays.TwoSum([]int{3, 2, 4}, 6)
	if indexes[0] != 1 || indexes[1] != 2 {
		t.Fatal("expected (1, 2), but got", indexes)
	}
}

func Test_arrayOfFourElements_returnsFirstAndSecondIndex(t *testing.T) {
	indexes := arrays.TwoSum([]int{3, 4, 5, 6}, 7)
	if indexes[0] != 0 || indexes[1] != 1 {
		t.Fatal("expected (0, 1), but got", indexes)
	}
}
