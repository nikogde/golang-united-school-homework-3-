package homework

func reverse(input []int64) (result []int64) {
	i := 0
	j := len(input) - 1
	for i < j {
		input[i], input[j] = input[j], input[i]
		i++
		j--
	}
	return input
}
