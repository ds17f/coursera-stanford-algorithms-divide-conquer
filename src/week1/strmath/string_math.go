package strmath

import (
	"fmt"
	"math"
	"strconv"
)

// Pow10str raises x * 10^n where x is some integer as a string, an n is an integer
// This function is designed to allow you to raise a number to a power of 10 that
//  is very large.  For example, x = "12345678" n = 128.  This would overflow an int
//  but this method will pad x with 128 0's on the right and return the new string
func Pow10str(x string, n int) string {
	formatStr := "%0" + strconv.Itoa(n) + "s"
	return x + fmt.Sprintf(formatStr, "")
}

// LeftPadString left pads str with 0's so len(result) = n
func LeftPadString(str string, n int) string {
	formatStr := "%0" + strconv.Itoa(n) + "s"
	return fmt.Sprintf(formatStr, str)
}

// LeftPadPowOfTwo adds leading 0's to str until len(result) is a power of 2
func LeftPadPowOfTwo(str string) string {
	log2 := math.Log2(float64(len(str)))
	floorLog2 := math.Floor(log2)
	// if we're already at a pow of 2
	if log2-floorLog2 <= .0001 {
		// do nothing
		return str
	}
	nextPowOf2 := floorLog2 + 1.0
	newSize := math.Pow(2, nextPowOf2)
	return LeftPadString(str, int(newSize))
}

// StrMult multiplies x * y and returns the result
// x and y are both string representations of integers
// but should be small in magnitude (length) to prevent
// integer overflow during type conversion
func StrMult(x string, y string) string {
	xi := strToInt64(x)
	yi := strToInt64(y)
	//xi, _ := strconv.Atoi(x)
	//yi, _ := strconv.Atoi(y)
	return strconv.FormatInt(xi*yi, 10)
}
func strToInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("Integer overflow for: %s, value: %d\n", str, n)
	}
	return n
}

// StrAdd adds xa and ya together and returns the result.
// xa and xy are both string representations of integers of arbitrary length.
// The returned string is the string representation of the sum of xa, ya.
// This function is used to add integers greater than 64 bits.
func StrAdd(x string, y string) string {
	// isolate the signs of each number
	// from the base of the number
	xSign, absX := StrSignAndAbs(x)
	ySign, absY := StrSignAndAbs(y)

	// TODO: can probably just pad to pow of 2
	// Pad strings if necessary
	if len(absX) != len(absY) {
		z := int(math.Max(float64(len(absX)), float64(len(absY))))
		absX = LeftPadString(absX, z)
		absY = LeftPadString(absY, z)
	}

	// if the signs are different then we have to do subtraction
	if xSign != ySign {
		return StrSub(x, y)
	}

	result := ""
	carry := 0
	// grade school addition
	for i := len(absX) - 1; i > -1; i-- {
		// get chars at both positiongs
		xi, _ := strconv.Atoi(string(absX[i]))
		yi, _ := strconv.Atoi(string(absY[i]))

		sum := xi + yi + carry

		carry = 0
		if sum > 9 {
			carry = 1
			sum -= 10
		}

		result = strconv.Itoa(sum) + result
	}
	// if we've exhausted all places
	// but we still have a carried 1 we need to
	// prepend it
	// ex: 6 + 7, we ended the loop with result = 3
	// but we've still got a carried 1 so that we can make 13
	if carry == 1 {
		result = "1" + result
	}

	// and finally, we know that the signs were the same
	// for both x and y, so we can just use x's sign in the result
	return xSign + result
}

// StrSignAndAbs takes x and returns a tuple of strings
// the first string returned is "" if x is positive, "-" if x is negative
// the second string is the absolute value of x
func StrSignAndAbs(x string) (string, string) {
	if string(x[0]) == "-" {
		return "-", x[1:len(x)]
	}
	return "", x
}

// StrSub subtracts x - y and returns the result as a string.
// x is some int as a string.  y is some int as a string.
// this is intended for large numbers
func StrSub(x string, y string) string {
	// isolate the signs of each number
	// from the base of the number
	xSign, absX := StrSignAndAbs(x)
	ySign, absY := StrSignAndAbs(y)

	// Pad strings if necessary
	if len(absX) != len(absY) {
		z := int(math.Max(float64(len(absX)), float64(len(absY))))
		absX = LeftPadString(absX, z)
		absY = LeftPadString(absY, z)
	}

	// we can group the "-" with the y term
	// which effectively means we can flip
	// y's sign
	if ySign == "-" {
		ySign = ""
	} else {
		ySign = "-"
	}

	// now if the signs match we're just
	// doing addition and maintaining the common sign
	if xSign == ySign {
		return xSign + StrAdd(absX, absY)
	}

	// if the signs are different and the numbers are the same
	if absX == absY {
		// the result will always be 0 (3 - 3) or (-3 - (-3))
		return "0"
	}

	// we need to find out which is bigger, abs(x) or abs(y)
	// because that will tell us the sign (xSign or ySign)
	// of the resulting number as well as when we subtract,
	// which term we subtract from the other.
	// to do this we loop through the digits of x and y
	// comparing each digit to find the larger left most digit
	var larger string
	var smaller string
	var largeSign string
	for i := 0; i < len(absX); i++ {
		xi, _ := strconv.Atoi(string(absX[i]))
		yi, _ := strconv.Atoi(string(absY[i]))
		if xi > yi {
			larger = absX
			largeSign = xSign
			smaller = absY
			break
		} else if xi < yi {
			larger = absY
			largeSign = ySign
			smaller = absX
			break
		}
	}

	// finally use grade school subtraction
	// on a special set (larger > smaller, both are > 0)
	// and use the sign of the larger on the result
	return largeSign + gradeSchoolSubtraction(larger, smaller)

}

// Assumptions:
//  x > y
//  x > 0 && y > 0
//  len(x) == len(y)
func gradeSchoolSubtraction(x string, y string) string {
	result := ""
	carry := 0
	// grade school subtraction
	for i := len(x) - 1; i > -1; i-- {
		// get chars at both positiongs
		xi, _ := strconv.Atoi(string(x[i]))
		yi, _ := strconv.Atoi(string(y[i]))

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
	}
	return result
}
