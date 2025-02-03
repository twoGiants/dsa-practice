package arrays

func IsAnagram(s, t string) bool {
	notEqualLength := len(s) != len(t)
	if notEqualLength {
		return false
	}

	countS, countT := countCharacters(s), countCharacters(t)

	return equalCharacterCounts(countS, countT)
}

func equalCharacterCounts(countS, countT map[rune]int) bool {
	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}

func countCharacters(str string) map[rune]int {
	result := make(map[rune]int)
	for _, char := range str {
		result[char]++
	}
	return result
}
