package main

import (
	"math"
)

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
