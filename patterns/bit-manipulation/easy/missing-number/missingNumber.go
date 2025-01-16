package bit

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
