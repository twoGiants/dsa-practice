package arrays

// Time: O(n), Space: O(n)
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	prefixes := make([]int, n)
	suffixes := make([]int, n)

	prefixes[0], suffixes[n-1] = 1, 1
	for i := 1; i < n; i++ {
		prefixes[i] = prefixes[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		suffixes[i] = suffixes[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		result[i] = prefixes[i] * suffixes[i]
	}

	return result
}

func ProductExceptSelfOptimal(nums []int) []int {
	prefix := 1
	result := []int{}
	for _, n := range nums {
		result = append(result, prefix)
		prefix *= n
	}

	postfix := 1
	for i := len(nums) - 1; i > -1; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
