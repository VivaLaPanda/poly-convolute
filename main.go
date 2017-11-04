package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/VivaLaPanda/poly-convolute/multmethods"
)

func main() {
	fmt.Printf("Enter the length of the polynomials: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inText := scanner.Text()
	temp, _ := strconv.ParseInt(inText, 10, 64)
	polyLength := int(temp)
	resLength := NextPowerOfTwo(polyLength)

	// Pad to have a degree of form 2^k-1
	// We use two identical slices to avoid mutation problems - non-optimal
	fftX := make([]float64, resLength)
	dirX := make([]float64, resLength)
	for i := 0; i < polyLength; i++ {
		rand := rand.Float64()
		fftX[i] = rand
		dirX[i] = rand
	}
	fftY := make([]float64, resLength)
	dirY := make([]float64, resLength)
	for i := 0; i < polyLength; i++ {
		rand := rand.Float64()
		fftY[i] = rand
		dirY[i] = rand
	}

	fftStart := time.Now()
	fftConv := multmethods.FftConv(fftX, fftY)
	fftElapsed := time.Since(fftStart)

	dirStart := time.Now()
	dirConv := multmethods.DirectMult(dirX, dirY)
	dirElapsed := time.Since(dirStart)

	if polyLength < 100 {
		fmt.Printf("Convolution of %v * %v\n", fftX, fftY)
		fmt.Printf("FFT resulted in\n%v\nin %v\n", fftConv, fftElapsed)
		fmt.Printf("Direct resulted in\n%v\nin %v\n", dirConv, dirElapsed)
	} else {
		fmt.Printf("Convolution:\n")
		fmt.Printf("FFT took %v\n", fftElapsed)
		fmt.Printf("Direct took %v\n", dirElapsed)

		data := fmt.Sprintf("%v * %v => %v", fftX, fftY, fftConv)

		_ = ioutil.WriteFile("out.txt", []byte(data), 0644)
	}
}

func NextPowerOfTwo(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
