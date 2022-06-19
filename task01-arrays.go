package homework

func average(input [15]float32) (result float32) {
	var array_sum float32 = 0
	for _, i := range input {
		array_sum += i
	}
	result = array_sum / float32(len(input))
	return result
}
