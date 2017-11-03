package multmethods

func Fft(x []int, y []int) (res []int) {
	return make([]int, 0)
}

func DirectMult(x []int, y []int) (res []int) {
	out := make([]int, len(x) + len(y))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			out[i+j] += x[i] * y[j]
		}
	}
	return out
}
