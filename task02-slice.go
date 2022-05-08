package homework

func reverse(input []int64) (result []int64) {
	cinput := make([]int64, len(input))
	copy(cinput, input)
	for i := len(cinput) - 1; i >= 0; i-- {
		result = append(result, cinput[i])
	}
	return result
}
