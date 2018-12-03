package lib_test

import (
	"testing"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week3/lib"
)

func TestCloneSlice(t *testing.T) {
	testCase := []int{1, 2, 3, 4, 5}
	clone := lib.CloneSlice(testCase)
	// compare values, they should match
	if !lib.CompareSlices(clone, testCase) {
		t.Logf("CloneSlice did not produce an exact copy of the original: %v. Produced: %v", testCase, clone)
		t.Fail()
	} else {
		// change something and see if it affects the first
		clone[3] = 100
		if lib.CompareSlices(clone, testCase) {
			t.Logf("CloneSlice did not produce an independent copy of the original")
			t.Fail()
		}
	}
}

func TestCompareSlices(t *testing.T) {
	cases := []struct {
		a   []int
		b   []int
		exp bool
	}{
		{[]int{2, 3, 4, 5}, []int{2, 3, 4}, false},
		{[]int{2, 3, 4, 5}, []int{2, 3, 4, 5}, true},
		{[]int{2, 3, 4, 5}, []int{5, 3, 4, 2}, false},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		result := lib.CompareSlices(testCase.a, testCase.b)
		if result != testCase.exp {
			t.Logf("CompareSlices failed for a: %v, b: %v, expected: %v",
				testCase.a, testCase.b, testCase.exp)
			t.Fail()
		}
	}
}

func TestAllLess(t *testing.T) {
	cases := []struct {
		a   []int
		val int
		exp bool
	}{
		{[]int{1, 2, 3, 4, 5}, 20, true},
		{[]int{1, 2, 3, 4, 5}, 2, false},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		result := lib.AllLess(testCase.a, testCase.val)
		if result != testCase.exp {
			t.Logf("AllLess failed for case: %v all less than: %v, expected: %v", testCase.a, testCase.val, testCase.exp)
			t.Fail()
		}
	}
}

func TestAllGreater(t *testing.T) {
	cases := []struct {
		a   []int
		val int
		exp bool
	}{
		{[]int{2, 3, 4, 5}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 2, false},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		result := lib.AllGreater(testCase.a, testCase.val)
		if result != testCase.exp {
			t.Logf("AllGreater failed for case: %v all greater than: %v, expected: %v", testCase.a, testCase.val, testCase.exp)
			t.Fail()
		}
	}
}

func TestSwap(t *testing.T) {
	cases := []struct {
		a   []int
		i   int
		j   int
		exp []int
	}{
		{[]int{2, 3, 4, 5}, 1, 2, []int{2, 4, 3, 5}},
		{[]int{2, 3, 4, 5}, 0, 3, []int{5, 3, 4, 2}},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		clone := lib.CloneSlice(testCase.a)
		lib.Swap(clone, testCase.i, testCase.j)
		if !lib.CompareSlices(clone, testCase.exp) {
			t.Logf("Swap failed for case: %v swap a[i]: %d, a[j]: %d, expected: %v, produced: %v",
				testCase.a, testCase.a[testCase.i], testCase.a[testCase.j], testCase.exp, clone)
			t.Fail()
		}
	}
}
