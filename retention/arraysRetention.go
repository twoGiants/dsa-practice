package retention

import "math"

func SetZeroes(matrix [][]int) {
	ROWS, COLS := len(matrix), len(matrix[0])
	rowZero := false

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				if r > 0 {
					matrix[r][0] = 0
				} else {
					rowZero = true
				}
			}
		}
	}

	for r := 1; r < ROWS; r++ {
		for c := 1; c < COLS; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for r := 0; r < ROWS; r++ {
			matrix[r][0] = 0
		}
	}

	if rowZero {
		for c := 0; c < COLS; c++ {
			matrix[0][c] = 0
		}
	}
}

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
