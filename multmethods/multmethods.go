package multmethods

import (
	"math"
	"math/cmplx"
)

// FftConv is a function which takes to vectors representing polynomials and
// returns their convolution.
// x and y mus tbe the same length and their length must be a multiple of 2
// This method is not guarunteed to not mutate its input
func FftConv(x []float64, y []float64) (res []float64) {
	xFft := make([]complex128, len(x)*2)
	yFft := make([]complex128, len(x)*2)
	for i, elem := range x {
		xFft[i] = complex(float64(elem), float64(0))
	}
	for i, elem := range y {
		yFft[i] = complex(float64(elem), float64(0))
	}

	xFft = Fft(xFft)
	yFft = Fft(yFft)

	convFft := make([]complex128, len(xFft))
	for i := range yFft {
		convFft[i] = xFft[i] * yFft[i]
	}

	resCmplx := InvFft(convFft)
	res = make([]float64, len(resCmplx))
	for i, elem := range resCmplx {
		res[i] = float64(roundFloat(real(elem), 4))
	}

	return
}

// Fft is a function which produces the Discreet Fourier Transform of input data
// It will mutate any input passed to it, so if you need to preserve input
// make sure to deref the slice you pass in.
func Fft(data []complex128) []complex128 {
	if len(data) == 1 {
		return data
	}
	hl := len(data) / 2
	even := make([]complex128, hl)
	odd := make([]complex128, hl)
	for i := 0; i < hl; i++ {
		even[i] = data[i] + data[i+hl]
		odd[i] = (data[i] - data[i+hl]) *
			cmplx.Exp(complex(0, 2*float64(i)*math.Pi/float64(len(data))))
	}

	Fft(even)
	Fft(odd)
	for i := 0; i < hl; i++ {
		data[2*i] = even[i]
		data[2*i+1] = odd[i]
	}

	return data
}

func InvFft(data []complex128) []complex128 {
	revStr := data[1:]
	last := len(revStr) - 1
	for i := 0; i < len(revStr)/2; i++ {
		revStr[i], revStr[last-i] = revStr[last-i], revStr[i]
	}

	Fft(data)
	for i := range data {
		data[i] = data[i] / complex(float64(len(data)), float64(0))
	}

	return data
}

func DirectMult(x []float64, y []float64) (res []float64) {
	out := make([]float64, len(x)+len(y))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			out[i+j] += x[i] * y[j]
		}
	}
	return out
}

func RemoveTrailingZeros(in []int) []int {
	i := 0
	for i = len(in) - 1; i >= 0 && in[i] == 0; i-- {
	}
	return in[:i]
}

// return rounded version of x with prec precision.
func roundFloat(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	intermed += .5
	x = .5
	if frac < 0.0 {
		x = -.5
		intermed -= 1
	}
	if frac >= x {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
