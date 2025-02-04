package assert

import (
	"slices"
	"sort"
	"testing"
)

type IntSliceDimensions interface {
	[]int | [][]int
}

func SliceLengthEqual(t *testing.T, result [][]int, expected [][]int) {
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, but got %d", len(expected), len(result))
	}
}

func SlicesEqual[T IntSliceDimensions](t *testing.T, result, expected []int, input T) {
	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}

func NumbersEqual(t *testing.T, result, expected int, input []int) {
	if result != expected {
		t.Fatalf("for input %v expected %d, but got %d", input, expected, result)
	}
}

func NoDuplicates(t *testing.T, result bool, input []int) {
	if result {
		t.Fatalf("expected %v to have no duplicates, but got some", input)
	}
}

func Duplicates(t *testing.T, result bool, input []int) {
	if !result {
		t.Fatalf("expected %v to have duplicates, but got none", input)
	}
}

func MatrixEqual(t *testing.T, input, expected, copy [][]int) {
	equal := true
	for i := 0; i < len(input); i++ {
		if !slices.Equal(input[i], expected[i]) {
			equal = false
		}
	}

	if !equal {
		t.Fatalf("for input %v expected output %v, but got %v", copy, expected, input)
	}
}

func SortGroups(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) != len(groups[j]) {
			return len(groups[i]) < len(groups[j])
		}
		for k := range groups[i] {
			if groups[i][k] != groups[j][k] {
				return groups[i][k] < groups[j][k]
			}
		}
		return false
	})
}
