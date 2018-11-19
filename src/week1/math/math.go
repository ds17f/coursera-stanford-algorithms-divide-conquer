package math

import (
	"math"
)

// GetPlaces returns the number of places in the integer n
func GetPlaces(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(float64(n))) + 1)
}

// GetHalves splits a number down the middle and returns
// each half as an integer.  If n = 12345678 then GetHalves
// returns 1234 and 5678.
// If there are an odd number of digits the first half will
// be shorter than the second half
func GetHalves(n int) (int, int) {
	nPlaces := GetPlaces(n)
	if int(nPlaces)%2 != 0 {
		nPlaces += 1.0
	}
	nPow := math.Pow10(nPlaces / 2)
	lHalf := math.Floor(float64(n) / nPow)
	rHalf := float64(n) - (lHalf * nPow)
	return int(lHalf), int(rHalf)
}
