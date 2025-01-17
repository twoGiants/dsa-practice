package boilerplate_test

import (
	"dsa/boilerplate"
	"testing"
)

func Test_XXX(t *testing.T) {
	input := []int{}
	result := boilerplate.BoilerplateFunc(input)
	expected := 0
	if result != expected {
		t.Fatalf("for input %v expected output %d, but got %d", input, expected, result)
	}
	// expected := []int{}
	// if !slices.Equal(result, expected) {
	// 	t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	// }
}
