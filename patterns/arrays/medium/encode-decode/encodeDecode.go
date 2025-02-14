package arrays

import "strconv"

type Transfer struct{}

func (t *Transfer) Encode(strs []string) string {
	var result string
	for _, str := range strs {
		result += strconv.Itoa(len(str)) + "#" + str
	}
	return result
}

func (t *Transfer) Decode(encoded string) []string {
	var result []string
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		wordLength, _ := strconv.Atoi(encoded[i:j])
		wordStart := j + 1
		wordEnd := wordStart + wordLength
		result = append(result, encoded[wordStart:wordEnd])
		i = wordEnd
	}

	return result
}
