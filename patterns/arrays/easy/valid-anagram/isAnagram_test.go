package arrays_test

import (
	arrays "dsa/patterns/arrays/easy/valid-anagram"
	"testing"
)

func Test_IsAnagram(t *testing.T) {
	var anagramTests = []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"racecar", "carrace", true},
		{"jar", "jam", false},
	}

	for _, at := range anagramTests {
		t.Run(at.s1+"."+at.s2, func(t *testing.T) {

			result := arrays.IsAnagram(at.s1, at.s2)

			if result != at.expected {
				t.Errorf("for input '%s', '%s' expected output %t, but got %t", at.s1, at.s2, at.expected, result)
			}
		})
	}
}
