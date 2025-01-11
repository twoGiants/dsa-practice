package arrays

func HasDuplicate(numbers []int) bool {
	seen := make(map[int]bool)
	for _, number := range numbers {
		if seen[number] {
			return true
		}
		seen[number] = true
	}

	return false
}
