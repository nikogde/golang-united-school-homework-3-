package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var newMap = make([]int, 0)
	for i, _ := range input {
		newMap = append(newMap, i)
	}
	sort.Ints(newMap)

	var sorted_vals = make([]string, 0)
	for i, _ := range newMap {
		sorted_vals = append(sorted_vals, input[newMap[i]])
	}

	return sorted_vals
}
