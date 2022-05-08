package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//value := make([]string, 0, len(input))
	for _, v := range keys {
		result = append(result, input[v])
	}
	return result

}
