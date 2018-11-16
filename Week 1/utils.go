package main

import (
	"fmt"
	"math"
	"strconv"
)

func leftPadPowOfTwo(str string) string {
	log2 := math.Log2(float64(len(str)))
	floorLog2 := math.Floor(log2)
	// if we're already at a pow of 2
	if log2-floorLog2 <= .0001 {
		// do nothing
		return str
	}
	nextPowOf2 := floorLog2 + 1.0
	newSize := math.Pow(2, nextPowOf2)
	return leftPadString(str, int(newSize))
}

func leftPadString(str string, size int) string {
	formatStr := "%0" + strconv.Itoa(size) + "s"
	return fmt.Sprintf(formatStr, str)
}

func getPlaces(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(float64(n))) + 1)
}

func getHalves(n int) (int, int) {
	nPlaces := getPlaces(n)
	if int(nPlaces)%2 != 0 {
		nPlaces += 1.0
	}
	nPow := math.Pow10(nPlaces / 2)
	lHalf := math.Floor(float64(n) / nPow)
	rHalf := float64(n) - (lHalf * nPow)
	return int(lHalf), int(rHalf)
}
