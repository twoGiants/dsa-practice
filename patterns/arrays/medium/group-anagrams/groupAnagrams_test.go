package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/group-anagrams"
	"slices"
	"sort"
	"testing"
)

func sortGroups(groups [][]string) {
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

			sortGroups(result)
			sortGroups(at.expected)

			for i := 0; i < len(at.expected); i++ {
				if !slices.Equal(result[i], at.expected[i]) {
					t.Errorf("for input %v expected output %v, but got %v", at.input, at.expected[i], result[i])
				}
			}
		})
	}
}
