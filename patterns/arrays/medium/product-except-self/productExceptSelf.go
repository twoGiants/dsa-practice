package arrays

// Time: O(n), Space: O(n)
func ProductExceptSelfV2(nums []int) []int {
	prefix := 1
	prefixList := []int{}
	for _, n := range nums {
		prefixList = append(prefixList, prefix)
		prefix *= n
	}

	postfix := 1
	postfixList := make([]int, len(nums))
	for i := len(nums) - 1; i > -1; i-- {
		postfix *= nums[i]
		postfixList[i] = postfix
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			res = append(res, prefixList[i]*1)
		} else {
			res = append(res, prefixList[i]*postfixList[i+1])
		}
	}

	return res
}

func ProductExceptSelf(nums []int) []int {
	prefix := 1
	res := []int{}
	for _, n := range nums {
		res = append(res, prefix)
		prefix *= n
	}

	postfix := 1
	for i := len(nums) - 1; i > -1; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
