package retention

import "math"

func FindDuplicates(nums []int) []int {
	res := []int{}
	for _, n := range nums {
		n := int(math.Abs(float64(n)))

		if nums[n-1] < 0 {
			res = append(res, n)
		}

		nums[n-1] = -1 * nums[n-1]
	}
	return res
}

func Construct2DArray(original []int, m, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	res := [][]int{}
	for i := 0; i < m; i++ {
		start := i * n
		end := start + n
		res = append(res, original[start:end])
	}

	return res
}

func SingleNumber(nums []int) int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	return xor
}

func FindDisappearedNumbers(nums []int) []int {
	for _, n := range nums {
		i := int(math.Abs(float64(n))) - 1
		nums[i] = -1 * int(math.Abs(float64(nums[i])))
	}

	result := []int{}
	for i, n := range nums {
		if n > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func MissingNumber(nums []int) int {
	result := len(nums)
	for i, n := range nums {
		result += i - n
	}
	return result
}

func MissingNumberXor(nums []int) int {
	xor := len(nums)
	for i, n := range nums {
		xor ^= i ^ n
	}
	return xor
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
