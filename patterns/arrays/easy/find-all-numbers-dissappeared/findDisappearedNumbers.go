package arrays

import "math"

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
