package main

import (
	"math"

	mainmath "github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week1/math"
)

func recursiveIntegerMutliplication(x int, y int) int {
	n := mainmath.GetPlaces(x)
	m := mainmath.GetPlaces(y)

	// error
	if n != m {
		return -1
	}

	// Base case
	if n == 1 {
		return x * y
	}

	a, b := mainmath.GetHalves(x)
	c, d := mainmath.GetHalves(y)

	ac := recursiveIntegerMutliplication(a, c)
	ad := recursiveIntegerMutliplication(a, d)
	bc := recursiveIntegerMutliplication(b, c)
	bd := recursiveIntegerMutliplication(b, d)

	tenToN := int(math.Pow10(n))
	tenToHalfN := int(math.Pow10(n / 2))

	return (tenToN * ac) + (tenToHalfN * (ad + bc)) + bd
}
