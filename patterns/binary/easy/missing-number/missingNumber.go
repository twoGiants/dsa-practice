package binary

func MissingNumber(nums []int) int {
	existing := make(map[int]int)

	for i, num := range nums {
		existing[num] = i
	}

	for i := 0; i <= len(nums); i++ {
		if _, found := existing[i]; !found {
			return i
		}
	}

	return -1
}
