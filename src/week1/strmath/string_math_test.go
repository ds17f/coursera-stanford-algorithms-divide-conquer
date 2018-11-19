package strmath_test

import (
	"testing"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week1/strmath"
)

func TestPow10str(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		x string
		n int
		z string
	}{
		{"123", 3, "123000"},
		{"736", 20, "73600000000000000000000"},
	}
	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := strmath.Pow10str(testCase.x, testCase.n)
		if testCase.z != actual {
			t.Logf("Pow10str failed %s*10^%d = %s, expected %s", testCase.x, testCase.n, testCase.z, actual)
			t.Fail()
		}
	}
}

func TestStrMult(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		x string
		y string
		z string
	}{
		{"5", "6", "30"},
		{"2", "8", "16"},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := strmath.StrMult(testCase.x, testCase.y)
		if testCase.z != actual {
			t.Logf("StrMult failed to add %s * %s = %s, expected %s", testCase.x, testCase.y, testCase.z, actual)
			t.Fail()
		}
	}
}

func TestStrAdd(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		x string
		y string
		z string
	}{
		{"100", "100", "200"},
		{"200", "400", "600"},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := strmath.StrAdd(testCase.x, testCase.y)
		if testCase.z != actual {
			t.Logf("StrAdd failed to add %s + %s = %s, expected %s", testCase.x, testCase.y, testCase.z, actual)
			t.Fail()
		}
	}
}

func TestLeftPadString(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		str string
		n   int
		z   string
	}{
		{"123", 6, "000123"},
		{"567", 23, "00000000000000000000567"},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := strmath.LeftPadString(testCase.str, testCase.n)
		if testCase.z != actual {
			t.Logf("LeftPadString failed str: %s, n: %d, expected: %s, result: %s", testCase.str, testCase.n, testCase.z, actual)
			t.Fail()
		}
	}
}

func TestLeftPadPowOfTwo(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		str string
		z   string
	}{
		{"123", "0123"},
		{"12334", "00012334"},
		{"12345678", "12345678"},
		{"1234567890", "0000001234567890"},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := strmath.LeftPadPowOfTwo(testCase.str)
		if testCase.z != actual {
			t.Logf("LeftPadPowOfTwo failed str: %s, expected: %s, result: %s", testCase.str, testCase.z, actual)
			t.Fail()
		}
	}
}
