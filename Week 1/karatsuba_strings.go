package main

import "strconv"
import "math"

func splitString(str string) (string, string) {
	midPoint := int(math.Floor(float64(len(str) / 2)))
	return str[:midPoint], str[midPoint:]
}

func strMult(x string, y string) string {
	xi, _ := strconv.Atoi(x)
	yi, _ := strconv.Atoi(y)
	return strconv.Itoa(xi * yi)
}

func strAdd(x string, y string) string {
	xi, _ := strconv.Atoi(x)
	yi, _ := strconv.Atoi(y)
	return strconv.Itoa(xi + yi)
}

func strSub(x string, y string) string {
	xi, _ := strconv.Atoi(x)
	yi, _ := strconv.Atoi(y)
	return strconv.Itoa(xi - yi)
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
	return strAdd(strAdd(term1, term2), bd)
}
