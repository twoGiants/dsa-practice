package arrays

func TwoSum(numbers []int, target int) []int {
	visited := make(map[int]int)

	for i, num := range numbers {
		diff := target - num
		if j, found := visited[diff]; found {
			return []int{j, i}
		}
		visited[num] = i
	}

	return nil
}
