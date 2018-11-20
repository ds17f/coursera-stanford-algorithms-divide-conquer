package algorithms

import (
	"math"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week1/strmath"
)

func splitString(str string) (string, string) {
	midPoint := int(math.Floor(float64(len(str) / 2)))
	return str[:midPoint], str[midPoint:]
}

//KaratsubaMult multiplies x * y using the karatsuba algorithm.
// x and y are both string representations of integers of
// an arbitrary length.
func KaratsubaMult(x string, y string) string {
	// we want to work with the abs(x) and abs(y)
	// when we evaluate their need for padding
	// as well as their length
	// We'll add back in the signs when we do
	// work on the numbers
	xSign, x1 := strmath.StrSignAndAbs(x)
	ySign, y1 := strmath.StrSignAndAbs(y)

	// pad both strings to be powers of 2
	x1 = strmath.LeftPadPowOfTwo(x1)
	y1 = strmath.LeftPadPowOfTwo(y1)

	m := float64(len(x1))
	o := float64(len(y1))
	n := int(math.Max(m, o))

	x1 = strmath.LeftPadString(x1, n)
	y1 = strmath.LeftPadString(y1, n)

	if n == 1 {
		return strmath.StrMult(xSign+x1, ySign+y1)
	}

	a, b := splitString(x1)
	c, d := splitString(y1)
	p := strmath.StrAdd(a, b)
	q := strmath.StrAdd(c, d)

	ac := KaratsubaMult(a, c)
	bd := KaratsubaMult(b, d)
	pq := KaratsubaMult(p, q)

	adbc := strmath.StrSub(strmath.StrSub(pq, ac), bd)

	term1 := strmath.Pow10str(ac, n)
	term2 := strmath.Pow10str(adbc, n/2)

	result := strmath.StrAdd(strmath.StrAdd(term1, term2), bd)
	// strip out any leading 0's
	for i := 0; i < len(result); i++ {
		if string(result[i]) != "0" {
			result = result[i:len(result)]
			break
		}
	}
	// if the signs didn't match then it's a negative
	if xSign != ySign {
		return "-" + result
	}
	return result
}
