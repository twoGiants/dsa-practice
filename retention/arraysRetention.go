package retention

import (
	"math"
)

func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	grouped := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num]++
	}

	for num, itsCount := range count {
		grouped[itsCount] = append(grouped[itsCount], num)
	}

	var result []int
	for i := len(grouped) - 1; i > 0; i-- {
		for _, num := range grouped[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

func GroupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, ch := range str {
			count[ch-'a']++
		}
		anagramMap[count] = append(anagramMap[count], str)
	}

	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)
	for i, ch := range s {
		countS[ch]++
		countT[rune(t[i])]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}

	return true
}

func LongestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	longest := 0
	for num := range numSet {
		if _, found := numSet[num-1]; !found {
			length := 1
			for {
				if _, exists := numSet[num+length]; exists {
					length++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

func SpiralOrder(matrix [][]int) []int {
	var result []int
	left, right := 0, len(matrix[0])
	top, bottom := 0, len(matrix)

	for left < right && top < bottom {
		// collect from top left to top right
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// collect from top right to bottom right
		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right-1])
		}
		right--

		if !(left < right && top < bottom) {
			break
		}

		// collect from bottom right to bottom left
		for i := right - 1; i >= left; i-- {
			result = append(result, matrix[bottom-1][i])
		}
		bottom--

		// collect from bottom left to top left
		for i := bottom - 1; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	return result
}

func RotateImage(matrix [][]int) {
	left, right := 0, len(matrix)-1

	for left < right {
		top, bottom := left, right

		for i := 0; i < right-left; i++ {
			// save top left
			topLeft := matrix[top][left+i]

			// move bottom left to top left
			matrix[top][left+i] = matrix[bottom-i][left]

			// move bottom right to bottom left
			matrix[bottom-i][left] = matrix[bottom][right-i]

			// move top right to bottom right
			matrix[bottom][right-i] = matrix[top+i][right]

			// move top left to top right
			matrix[top+i][right] = topLeft
		}

		left++
		right--
	}
}

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
