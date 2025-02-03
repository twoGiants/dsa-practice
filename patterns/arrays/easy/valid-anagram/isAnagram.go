package arrays

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)
	for i, ch := range s {
		countS[ch]++
		countT[rune(t[i])]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}
