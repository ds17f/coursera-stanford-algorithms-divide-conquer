package main

import (
	"fmt"
	"math"
)

func getPlaces(n int) int {
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

func recursiveIntegerMutliplication(x int, y int) int {
	n := getPlaces(x)
	m := getPlaces(y)

	// error
	if n != m {
		return -1
	}

	// Base case
	if n == 1 {
		return x * y
	}

	a, b := getHalves(x)
	c, d := getHalves(y)

	ac := recursiveIntegerMutliplication(a, c)
	ad := recursiveIntegerMutliplication(a, d)
	bc := recursiveIntegerMutliplication(b, c)
	bd := recursiveIntegerMutliplication(b, d)

	tenToN := int(math.Pow10(n))
	tenToHalfN := int(math.Pow10(n / 2))

	return (tenToN * ac) + (tenToHalfN * (ad + bc)) + bd
}

func main() {
	fmt.Println(recursiveIntegerMutliplication(1234, 5678))
}
