package retention

func SingleNumber(nums []int) int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	return xor
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
