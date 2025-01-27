package assert

import (
	"slices"
	"testing"
)

func SliceLengthEqual(t *testing.T, result [][]int, expected [][]int) {
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, but got %d", len(expected), len(result))
	}
}

func SlicesEqual(t *testing.T, result, expected, input []int) {
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func NumbersEqual(t *testing.T, result, expected int, input []int) {
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}
