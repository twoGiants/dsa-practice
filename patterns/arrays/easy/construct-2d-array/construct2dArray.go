package arrays

func Construct2DArray(original []int, m, n int) [][]int {
	var result [][]int
	if len(original) != m*n {
		return result
	}

	for i := 0; i < m; i++ {
		start := i * n
		end := start + n
		// sub slice does not include the number at the second index
		result = append(result, original[start:end])
	}

	return result
}
