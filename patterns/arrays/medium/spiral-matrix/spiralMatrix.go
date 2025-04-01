package arrays

func SpiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	left, right := 0, len(matrix[0])
	top, bottom := 0, len(matrix)

	for left < right && top < bottom {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right-1])
		}
		right--

		if left >= right || top >= bottom {
			break
		}

		for i := right - 1; i >= left; i-- {
			res = append(res, matrix[bottom-1][i])
		}
		bottom--

		for i := bottom - 1; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
