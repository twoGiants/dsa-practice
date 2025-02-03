package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/group-anagrams"
	"slices"
	"testing"
)

func Test_GroupAnagrams(t *testing.T) {
	var anagramTests = []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{"six words", []string{"act", "pots", "tops", "cat", "stop", "hat"}, [][]string{{"act", "cat"}, {"pots", "tops", "stop"}, {"hat"}}},
		{"x", []string{"x"}, [][]string{{"x"}}},
		{"empty string", []string{""}, [][]string{{""}}},
	}

	for _, at := range anagramTests {
		t.Run(at.name, func(t *testing.T) {

			result := arrays.GroupAnagrams(at.input)

			for i := 0; i < len(at.expected); i++ {
				if !slices.Equal(result[i], at.expected[i]) {
					t.Errorf("for input %v expected output %v, but got %v", at.input, at.expected[i], result[i])
				}
			}
		})
	}
}
