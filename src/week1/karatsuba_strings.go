package main

import (
	"math"
	"strconv"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week1/strmath"
)

func splitString(str string) (string, string) {
	midPoint := int(math.Floor(float64(len(str) / 2)))
	return str[:midPoint], str[midPoint:]
}

func stringNegative(x string) (int, string) {
	if string(x[0]) == "-" {
		return -1, x[1:len(x)]
	}
	return 1, x
}

// Assumptions:
//  x > y
//  x > 0 && y > 0
//  len(x) == len(y)
func _strSub(x string, y string) string {
	result := ""
	carry := 0
	// grade school subtraction
	for i := len(x) - 1; i > -1; i-- {
		// get chars at both positiongs
		xi, _ := strconv.Atoi(string(x[i]))
		yi, _ := strconv.Atoi(string(y[i]))
		//fmt.Printf("xi: %d, yi: %d\n", xi, yi)
		offset := 0
		// if x is bigger we need to add 10
		if xi+carry < yi {
			offset = 10
		}

		diff := offset + xi - yi + carry

		carry = 0
		// if we had an offset we've got to carry
		if offset == 10 {
			carry = -1
		} else {
			carry = 0
		}

		result = strconv.Itoa(diff) + result
		//println(result)
	}
	return result
}

func strAdd(xa string, ya string) string {
	// TODO: check leading sign and update x or y
	//  1 if pos, -1 if neg
	xNeg, x := stringNegative(xa)
	yNeg, y := stringNegative(ya)
	sign := ""
	if xNeg < 0 {
		sign = "-"
	}

	// Pad strings if necessary
	if len(x) != len(y) {
		z := int(math.Max(float64(len(x)), float64(len(y))))
		x = strmath.LeftPadString(x, z)
		y = strmath.LeftPadString(y, z)
	}

	// if the signs are different then we have to do subtraction
	if xNeg != yNeg {
		sign = ""
		//TODO: Do subtraction
		for i := 0; i < len(x); i++ {
			xi, _ := strconv.Atoi(string(x[i]))
			yi, _ := strconv.Atoi(string(y[i]))
			if xi > yi {
				if xNeg < 0 {
					sign = "-"
				}
				// x is the larger number
				return sign + _strSub(x, y)
			} else if xi < yi {
				if yNeg < 0 {
					sign = "-"
				}
				// y is the larger number
				return sign + _strSub(y, x)
			}
		}
		// the strings are the same, we return 0
		return "0"
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

		sum := xi + yi + carry

		carry = 0
		if sum < 0 {
			carry = 0
			sum = int(math.Abs(float64(sum)))
		} else if sum > 9 {
			carry = 1
			sum -= 10
		}

		result = strconv.Itoa(sum) + result
		//println(result)
	}
	if carry == 1 {
		result = "1" + result

	}
	return sign + result
}

func strSub(x string, y string) string {
	// we're going to negate y and then use
	// the add routine

	// if y's leading char is a "-"
	if string(y[0]) == "-" {
		// remove it
		return strAdd(x, y[1:len(y)])
	}

	// insert it
	return strAdd(x, "-"+y)

	//xi := strToInt64(x)
	//yi := strToInt64(y)
	//return strconv.FormatInt(xi-yi, 10)
}

func karatsubaString(x string, y string) string {
	// pad both strings to be powers of 2
	x1 := leftPadPowOfTwo(x)
	y1 := leftPadPowOfTwo(y)

	m := float64(len(x1))
	o := float64(len(y1))
	n := int(math.Max(m, o))

	x1 = strmath.LeftPadString(x1, n)
	y1 = strmath.LeftPadString(y1, n)

	if n == 1 {
		return strmath.StrMult(x1, y1)
	}

	a, b := splitString(x1)
	c, d := splitString(y1)
	p := strAdd(a, b)
	q := strAdd(c, d)

	ac := karatsubaString(a, c)
	bd := karatsubaString(b, d)
	pq := karatsubaString(p, q)

	adbc := strSub(strSub(pq, ac), bd)

	term1 := strmath.Pow10str(ac, n)
	term2 := strmath.Pow10str(adbc, n/2)

	return strAdd(strAdd(term1, term2), bd)
}
