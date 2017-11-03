package multmethods

import (
	"math"
	"math/cmplx"
)

func FftConv(x []int, y []int) (res []int) {
	xFft := make([]complex128, len(x)*2)
	yFft := make([]complex128, len(x)*2)
	for i, elem := range x {
		xFft[i] = complex(float64(elem), float64(0))
	}
	for i, elem := range x {
		yFft[i] = complex(float64(elem), float64(0))
	}

	xFft = Fft(xFft)
	yFft = Fft(yFft)

	convFft := make([]complex128, len(xFft))
	for i := range yFft {
		convFft[i] = xFft[i] * yFft[i]
	}

	resCmplx := InvFft(convFft)
	for i, elem := range resCmplx {
		res[i] = int(real(elem))
	}

	return
}

func Fft(data []complex128) []complex128 {
	return fft(data)
}

func fft(data []complex128) []complex128 {
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
	for i := range data {
		data[i] = complex(imag(data[i]), real(data[i]))
	}
	fft(data)
	scale := 1.0 / float64(len(data))
	for i := range data {
		data[i] = complex(imag(data[i])*scale, real(data[i])*scale)
	}

	return data
}

func DirectMult(x []int, y []int) (res []int) {
	return make([]int, 0)
}
