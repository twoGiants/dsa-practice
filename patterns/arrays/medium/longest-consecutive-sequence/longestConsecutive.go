package arrays

func LongestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	longest := 0
	for num := range numSet {
		length := 1
		if _, found := numSet[num-1]; !found {
			for {
				if _, exists := numSet[num+length]; exists {
					length++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
