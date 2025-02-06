package arrays_test

import (
	arrays "dsa/patterns/arrays/medium/encode-decode"
	"slices"
	"testing"
)

func Test_EncodeDecode(t *testing.T) {
	input := []string{"neet", "code", "love", "you"}
	expected := []string{"neet", "code", "love", "you"}

	transfer := arrays.Transfer{}

	encoded := transfer.Encode(input)
	result := transfer.Decode(encoded)

	if !slices.Equal(result, expected) {
		t.Fatalf("for input %v expected output %v, but got %v", input, expected, result)
	}
}
