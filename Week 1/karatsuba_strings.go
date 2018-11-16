package main

import (
	"fmt"
	"math"
	"strconv"
)

func strToInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("Integer overflow for: %s, value: %d\n", str, n)
	}
	return n
}

func splitString(str string) (string, string) {
	midPoint := int(math.Floor(float64(len(str) / 2)))
	return str[:midPoint], str[midPoint:]
}

func strMult(x string, y string) string {
	xi := strToInt64(x)
	yi := strToInt64(y)
	//xi, _ := strconv.Atoi(x)
	//yi, _ := strconv.Atoi(y)
	return strconv.FormatInt(xi*yi, 10)
}

func stringNegative(x string) (int, string) {
	if string(x[0]) == "-" {
		return -1, x[1:len(x)]
	}
	return 1, x
}

func strAdd(xa string, ya string) string {
	// TODO: check leading sign and update x or y
	//  1 if pos, -1 if neg
	xNeg, x := stringNegative(xa)
	yNeg, y := stringNegative(ya)

	// Pad strings if necessary
	if len(x) != len(y) {
		z := int(math.Max(float64(len(x)), float64(len(y))))
		x = leftPadString(x, z)
		y = leftPadString(y, z)
	}

	result := ""
	carry := 0
	// grade school addition
	// TODO: Consider adding negatives
	for i := len(x) - 1; i > -1; i-- {
		// get chars at both positiongs
		xi, _ := strconv.Atoi(string(x[i]))
		yi, _ := strconv.Atoi(string(y[i]))
		//fmt.Printf("xi: %d, yi: %d\n", xi, yi)

		sum := (xi * xNeg) + (yi * yNeg) + (carry * yNeg * xNeg)

		carry = 0
		if sum < 0 {
			carry = 1
			sum = int(math.Abs(float64(sum)))
		} else if sum > 9 {
			carry = 1
			sum -= 10
		}

		result = strconv.Itoa(sum) + result
		//println(result)
	}
	if xNeg*yNeg < 0 {
		result = "-" + result
	}
	return result
}

func strSub(x string, y string) string {
	xi := strToInt64(x)
	yi := strToInt64(y)
	return strconv.FormatInt(xi-yi, 10)

}

func karatsubaString(x string, y string) string {
	// pad both strings to be powers of 2
	x1 := leftPadPowOfTwo(x)
	y1 := leftPadPowOfTwo(y)

	m := float64(len(x1))
	o := float64(len(y1))
	n := int(math.Max(m, o))

	x1 = leftPadString(x1, n)
	y1 = leftPadString(y1, n)

	if n == 1 {
		return strMult(x1, y1)
	}

	a, b := splitString(x1)
	c, d := splitString(y1)
	p := strAdd(a, b)
	q := strAdd(c, d)

	ac := karatsubaString(a, c)
	bd := karatsubaString(b, d)
	pq := karatsubaString(p, q)

	adbc := strSub(strSub(pq, ac), bd)

	_10toN := strconv.Itoa(int(math.Pow10(n)))
	_10toHalfN := strconv.Itoa(int(math.Pow10(n / 2)))

	term1 := strMult(_10toN, ac)
	term2 := strMult(_10toHalfN, adbc)
	if n > 16 {
		fmt.Println("-------")
		fmt.Printf("n: %d\n", n)
		fmt.Printf("a: %s, b: %s, c: %s, d: %s\n", a, b, c, d)
		fmt.Printf("p: %s, q: %s\n", p, q)
		fmt.Printf("adbc: %s\n", adbc)
		fmt.Printf("ac: %s, bd: %s, pq: %s\n", ac, bd, pq)
		fmt.Printf("term1: %s, term2: %s, term3: %s\n", term1, term2, bd)
	}
	return strAdd(strAdd(term1, term2), bd)
}
