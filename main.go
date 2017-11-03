package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	bm "github.com/VivaLaPanda/binary-math/bitmath"
)

func main() {
	stopCode := false

	fmt.Printf("Algorithms Homework 2\n")
	fmt.Printf("Usage:\n  Check Primality of N with confidence K: '1'\n  Generate N bit prime with K confidence: '2'\n  N bit RSA keypair with confidence K: '3'\n  Kill: '4'\n")
	for stopCode == false {
		fmt.Printf("Enter your desired operation: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inText := scanner.Text()
		if len(inText) == 0 {
			continue
		}
		opChar := inText[0]

		switch opChar {
		case '1':
			fmt.Println("Please enter integers N,K in a nonspaced comma seperated list:")
			scanner.Scan()
			argText := scanner.Text()
			argArr := parseArgs(argText)
			fmt.Printf("Primality: %v\n", problem1(argArr[0], argArr[1]))
		case '2':
			fmt.Println("Please enter integers N,K in a nonspaced comma seperated list:")
			scanner.Scan()
			argText := scanner.Text()
			argArr := parseArgs(argText)
			fmt.Printf("New Prime: %v\n", problem2(argArr[0], argArr[1]))
		case '3':
			fmt.Println("Please enter integers N,K in a nonspaced comma seperated list:")
			scanner.Scan()
			argText := scanner.Text()
			argArr := parseArgs(argText)
			N, E, D := problem3a(argArr[0], argArr[1])
			fmt.Printf("N: %v, E: %v, D: %v\n", bm.BigBin2dec(N), bm.BigBin2dec(E), bm.BigBin2dec(D))
			fmt.Println("Please enter integers N,E,D,M in a nonspaced comma seperated list:")
			scanner.Scan()
			argText = scanner.Text()
			argArr = parseArgs(argText)
			encMsg, decMsg := problem3b(argArr[0], argArr[1], argArr[2], argArr[3])
			fmt.Printf("Encrypted Msg: %v, Decrypted Msg: %v\n", encMsg, decMsg)
		case '4':
			fmt.Println("Shutting down")
			stopCode = true
		default:
			fmt.Println("I'm sorry, I didn't understand that...")
			fmt.Printf("Usage:\n  Check Primality of N with confidence K: '1'\n  Generate N bit prime with K confidence: '2'\n  N bit RSA keypair with confidence K: '3'\n  Kill: '4'\n")
		}
	}
}

func problem1(n big.Int, confidence big.Int) (isPrime bool) {
	return bm.PrimalityThree(bm.BigDec2bin(n), bm.BigDec2bin(confidence))
}

func problem2(n big.Int, confidence big.Int) (newPrime big.Int) {
	return bm.BigBin2dec(bm.NBitPrime(bm.BigDec2bin(n), bm.BigDec2bin(confidence)))
}

func problem3a(nBits, confidence big.Int) (N, E, D []bool) {
	return bm.RSAKeygen(bm.BigDec2bin(nBits), bm.BigDec2bin(confidence))
}

func problem3b(N, E, D, M big.Int) (big.Int, big.Int) {
	encMsg := bm.RSAEnc(bm.BigDec2bin(M), bm.BigDec2bin(E), bm.BigDec2bin(N))
	decMsg := bm.RSADec(encMsg, bm.BigDec2bin(D), bm.BigDec2bin(N))

	return bm.BigBin2dec(encMsg), bm.BigBin2dec(decMsg)
}

func parseArgs(argText string) []big.Int {
	argArr := strings.Split(argText, ",")
	intArr := make([]big.Int, len(argArr))
	for i, elem := range argArr {
		intArr[i].SetString(elem, 10)
	}

	return intArr
}
