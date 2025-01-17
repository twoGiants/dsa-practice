package arrays

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
