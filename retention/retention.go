package retention

import "math"

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
	indexes := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if j, found := indexes[diff]; found {
			return []int{j, i}
		}
		indexes[n] = i
	}

	return nil
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
