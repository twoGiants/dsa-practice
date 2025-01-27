package retention

import "math"

func ProductExceptSelf(nums []int) []int {
	prefix := 1
	result := make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		prefix *= nums[i-1]
		result[i] = prefix
	}

	suffix := 1
	for i := len(nums) - 1; i >= 1; i-- {
		suffix *= nums[i]
		result[i-1] *= suffix
	}

	return result
}

func FindDuplicates(nums []int) []int {
	var result []int
	for _, n := range nums {
		i := int(math.Abs(float64(n))) - 1

		if nums[i] < 0 {
			result = append(result, i+1)
		}
		nums[i] = -nums[i]
	}
	return result
}

func Construct2DArray(original []int, m, n int) [][]int {
	var result [][]int
	if len(original) != m*n {
		return result
	}

	for i := 0; i < m; i++ {
		start := i * n
		end := start + n
		result = append(result, original[start:end])
	}

	return result
}

func FindDisappearedNumbers(nums []int) []int {
	for _, n := range nums {
		i := int(math.Abs(float64(n))) - 1
		nums[i] = -int(math.Abs(float64(nums[i])))
	}

	var result []int
	for i, n := range nums {
		if n > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func TwoSum(nums []int, target int) []int {
	visited := make(map[int]int)
	for i, n := range nums {
		diff := target - n
		if j, found := visited[diff]; found {
			return []int{j, i}
		}
		visited[n] = i
	}
	return []int{}
}

func HasDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}
