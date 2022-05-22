package homework

func average(input [15]float32) (result float32) {
	var sum float32
	for _, v := range input {
		sum += v
	}
	return sum / float32(len(input))
}
