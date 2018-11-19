package math_test

import (
	"testing"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week1/math"
)

func TestGetPlaces(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		n int
		z int
	}{
		{3, 1},
		{20, 2},
		{1383, 4},
		{9999999999, 10},
	}
	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := math.GetPlaces(testCase.n)
		if testCase.z != actual {
			t.Logf("GetPlaces failed %d digits: %d, expected %d", testCase.n, testCase.z, actual)
			t.Fail()
		}
	}
}

func TestGetHalves(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		n  int
		z  int
		z1 int
	}{
		{1234, 12, 34},
		{1234567, 123, 4567},
		{1234567890, 12345, 67890},
	}
	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual1, actual2 := math.GetHalves(testCase.n)
		if testCase.z != actual1 || testCase.z1 != actual2 {
			t.Logf("GetHalves failed %d split into : %d and %d, expected %d and %d",
				testCase.n, testCase.z, testCase.z1, actual1, actual2)
			t.Fail()
		}
	}
}
