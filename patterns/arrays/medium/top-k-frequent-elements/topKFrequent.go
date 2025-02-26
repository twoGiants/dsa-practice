package arrays

func TopKFrequent(nums []int, k int) []int {
	count := countAppearanceOfEachNumIn(nums)
	grouped := groupNumsByCount(nums, count)
	return grabKMostFrequent(grouped, k)
}

func grabKMostFrequent(grouped [][]int, k int) []int {
	var result []int
	for i := len(grouped) - 1; i > 0; i-- {
		result = append(result, grouped[i]...)
		if len(result) == k {
			return result
		}
	}
	return result
}

func groupNumsByCount(nums []int, count map[int]int) [][]int {
	result := make([][]int, len(nums)+1)
	for num, itsCount := range count {
		result[itsCount] = append(result[itsCount], num)
	}
	return result
}

func countAppearanceOfEachNumIn(nums []int) map[int]int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num]++
	}
	return result
}
