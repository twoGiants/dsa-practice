package retention

import (
	"math"
	"strconv"
)

type Transfer struct{}

func (*Transfer) Encode(strs []string) string {
	var encoded string
	for _, word := range strs {
		encoded += strconv.Itoa(len(word)) + "#" + word
	}
	return encoded
}

func (*Transfer) Decode(encoded string) []string {
	var decoded []string
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		end := start + length
		decoded = append(decoded, encoded[start:end])
		i = end
	}
	return decoded
}

func TopKFrequent(nums []int, k int) []int {
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		counts[num]++
	}

	grouped := make([][]int, len(nums)+1)
	for num, itsCount := range counts {
		grouped[itsCount] = append(grouped[itsCount], num)
	}

	var result []int
	for i := len(grouped) - 1; i > 0; i-- {
		for j := 0; j < len(grouped[i]); j++ {
			result = append(result, grouped[i][j])
			if len(result) == k {
				return result
			}
		}
	}

	return result
}

func GroupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[[26]int][]string)

	for _, word := range strs {
		var letterCount [26]int
		for _, letter := range word {
			letterCount[letter-'a']++
		}
		anagramGroups[letterCount] = append(anagramGroups[letterCount], word)
	}

	var result [][]string
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)
	for i, letter := range s {
		countS[letter]++
		countT[rune(t[i])]++
	}

	for letter, itsCount := range countS {
		if countT[letter] != itsCount {
			return false
		}
	}

	return true
}

func LongestConsecutive(nums []int) int {
	uniqueNums := make(map[int]bool)
	for _, num := range nums {
		uniqueNums[num] = true
	}

	longest := 1
	for num := range uniqueNums {
		if !uniqueNums[num-1] {
			length := 1
			for {
				if uniqueNums[num+length] {
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

		if left >= right || top >= bottom {
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
				if r == 0 {
					rowZero = true
				} else {
					matrix[r][0] = 0
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
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
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
	abs := func(n int) int { return int(math.Abs(float64(n))) }
	for _, n := range nums {
		i := abs(n) - 1
		nums[i] = -abs(nums[i])
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
	visited := make(map[int]int, len(nums))
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
