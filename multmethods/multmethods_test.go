package multmethods

import (
	"reflect"
	"testing"
)

func roundHelper(x complex128) complex128 {
	real := real(x)
	cmplx := imag(x)
	if real < 0 {
		real = float64(int(real - .5))
	} else {
		real = float64(int(real + .5))
	}
	if cmplx < 0 {
		cmplx = float64(int(cmplx - .5))
	} else {
		cmplx = float64(int(cmplx + .5))
	}

	return complex(real, cmplx)
}

var fftTest = []struct {
	in  []complex128
	out []complex128
}{
	{[]complex128{complex(5, 0), complex(0, 0)},
		[]complex128{complex(5, 0), complex(5, 0)}},
	{[]complex128{complex(1, 0), complex(-1, 0), complex(0, 0), complex(3, 0)},
		[]complex128{complex(3, 0), complex(1, -4), complex(-1, 0), complex(1, 4)}},
}

func TestFft(t *testing.T) {
	for _, tt := range fftTest {
		actual := Fft(tt.in)
		for i, elem := range actual {
			actual[i] = roundHelper(elem)
		}

		if !(reflect.DeepEqual(actual, tt.out)) {
			t.Errorf("Fft(%v) => %v, want %v", tt.in, actual, tt.out)
		}
	}
}

var multTests = []struct {
	inX []int
	inY []int
	out []int
}{
	{[]int{0}, []int{0}, []int{0}},
	{[]int{1}, []int{1}, []int{1}},
	{[]int{0, 1}, []int{0, 1}, []int{0, 0, 1}},
	{[]int{1, -1, 1, -2}, []int{1, -1, 2, -4}, []int{1, -2, 4, -9, 8, -8, 8}},
}

func TestFftConv(t *testing.T) {
	for _, tt := range multTests {
		actual := FftConv(tt.inX, tt.inY)
		actual = actual[:len(actual)-1]

		if !(reflect.DeepEqual(actual, tt.out)) {
			t.Errorf("FftConv(%v, %v) => %v, want %v", tt.inX, tt.inY, actual, tt.out)
		}
	}
}

func TestDirectMult(t *testing.T) {
	for _, tt := range multTests {
		actual := DirectMult(tt.inX, tt.inY)
		actual = actual[:len(actual)-1]

		if !(reflect.DeepEqual(actual, tt.out)) {
			t.Errorf("FftConv(%v, %v) => %v, want %v", tt.inX, tt.inY, actual, tt.out)
		}
	}
}
