package main

import (
	"fmt"
	"math"
)

func karatsuba(x int, y int) int {
	n := getPlaces(x)
	m := getPlaces(y)

	// error
	//if n != m {
	//	return -1
	//}

	// Base case
	if n == 1 {
		return x * y
	}

	a, b := getHalves(x)
	c, d := getHalves(y)
	p := a + b
	q := c + d
	// fmt.Println(a, b, c, d)
	// fmt.Println(p, q)

	ac := karatsuba(a, c)
	bd := karatsuba(b, d)
	pq := karatsuba(p, q)

	abcd := pq - ac - bd

	tenToN := int(math.Pow10(n))
	tenToHalfN := int(math.Pow10(n / 2))

	fmt.Println(n, m, tenToN, tenToHalfN)
	fmt.Println(ac, abcd, bd)
	return (tenToN * ac) + (tenToHalfN * abcd) + bd
}
