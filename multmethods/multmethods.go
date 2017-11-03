package multmethods

import (
	"math"
	"math/cmplx"
)

// FftConv is a function which takes to vectors representing polynomials and
// returns their convolution.
// x and y mus tbe the same length and their length must be a multiple of 2
// This method is not guarunteed to not mutate its input
func FftConv(x []int, y []int) (res []int) {
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
	res = make([]int, len(resCmplx))
	for i, elem := range resCmplx {
		var roundElem float64
		if real(elem) < 0 {
			roundElem = real(elem) - .5
		} else {
			roundElem = real(elem) + .5
		}
		res[i] = int(roundElem)
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

func DirectMult(x []int, y []int) (res []int) {
	out := make([]int, len(x)+len(y))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			out[i+j] += x[i] * y[j]
		}
	}
	return out
}
