package algorithms_test

import (
	"testing"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week3/algorithms"
)

var quickSortCases = []struct {
	n []int
	z int
}{
	{[]int{2, 3, 4}, 0},
	{[]int{3, 2, 4}, 1},
	{[]int{2, 4, 1, 3, 5}, 3},
	{[]int{3, 4, 1, 2, 5}, 4},
	{[]int{5, 4, 1, 3, 2}, 8},
}

// TODO: Need to work out the cases and the array expectations.  Not entirely sure how to mock with go and recursion
func TestQuickSort(t *testing.T) {
	// x, y, expected value
	for i := 0; i < len(quickSortCases); i++ {
		testCase := quickSortCases[i]
		_, actual := algorithms.QuickSort(algorithms.ChooseFirstPivot, testCase.n)
		if testCase.z != actual {
			t.Logf("QuickSort failed for case: %v.  Expected: %d but produced: %d", testCase.n, testCase.z, actual)
			t.Fail()
		}
	}
}

func TestChooseFirstPivot(t *testing.T) {
	cases := []struct {
		n []int
		z int
	}{
		{[]int{3}, 0},
		{[]int{3, 4}, 0},
		{[]int{2, 3, 4}, 0},
		{[]int{2, 1, 3, 5}, 0},
		{[]int{5, 4, 1, 3, 2}, 0},
	}
	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := algorithms.ChooseFirstPivot(testCase.n)
		if testCase.z != actual {
			t.Logf("ChooseFirstPivot failed for case: %v.  Expected: %d but produced: %d", testCase.n, testCase.z, actual)
			t.Fail()
		}
	}
}

func TestChooseLastPivot(t *testing.T) {
	cases := []struct {
		n []int
		z int
	}{
		{[]int{3}, 0},
		{[]int{3, 4}, 1},
		{[]int{2, 3, 4}, 2},
		{[]int{2, 1, 3, 5}, 3},
		{[]int{5, 4, 1, 3, 2}, 4},
	}
	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := algorithms.ChooseLastPivot(testCase.n)
		if testCase.z != actual {
			t.Logf("ChooseLastPivot failed for case: %v.  Expected: %d but produced: %d", testCase.n, testCase.z, actual)
			t.Fail()
		}
	}
}

func TestChooseMedianOfThreePivot(t *testing.T) {
	cases := []struct {
		n []int
		z int
	}{
		{[]int{5, 4, 1, 3, 2}, 4},
		{[]int{4, 1, 5, 3, 2}, 0},
		{[]int{5, 1, 4, 3, 2}, 2},

		{[]int{1, 4, 9, 3, 2, 5}, 5},
		{[]int{5, 4, 8, 10, 2, 3}, 0},
		{[]int{1, 4, 5, 10, 2, 9}, 2},
	}
	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := algorithms.ChooseMedianOfThreePivot(testCase.n)
		if testCase.z != actual {
			t.Logf("ChooseMedianOfThreePivot failed for case: %v.  Expected: %d but produced: %d", testCase.n, testCase.z, actual)
			t.Fail()
		}
	}
}
