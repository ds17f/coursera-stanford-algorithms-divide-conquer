package algorithms_test

import (
	"testing"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week2/algorithms"
)

func TestMergeSort(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		n []int
		z int
	}{
		{[]int{2, 3, 4}, 0},
		{[]int{3, 2, 4}, 1},
		{[]int{2, 4, 1, 3, 5}, 3},
		{[]int{3, 4, 1, 2, 5}, 4},
		{[]int{5, 4, 1, 3, 2}, 8},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := algorithms.CountInversions(testCase.n)
		if testCase.z != actual {
			t.Logf("CountInversions failed for case: %v.  Expected: %d but produced: %d", testCase.n, testCase.z, actual)
			t.Fail()
		}
	}
}
