package stack_test

import (
	stack "dsa/patterns/stack/easy/valid-parentheses"
	"testing"
)

func Test_IsValid(t *testing.T) {
	input := ""
	expected := false

	result := stack.IsValid(input)

	if result != expected {
		t.Fatalf("for input %s expected output %t, but got %t", input, expected, result)
	}
}
