package retention

func MissingNumber(nums []int) int {
	result := len(nums)
	for i := 0; i < len(nums); i++ {
		result += i - nums[i]
	}
	return result
}

func MissingNumberXor(nums []int) int {
	xor := len(nums)
	for i := 0; i < len(nums); i++ {
		xor ^= i ^ nums[i]
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

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}
