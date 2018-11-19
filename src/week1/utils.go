package main

import (
	"math"
)

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
