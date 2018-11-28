package algorithms

import (
	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week3/lib"
)

// fnChoosePivot defines a function prototype for
// the choose pivot routines.  It takes an array
// then chooses a pivot and returns the index of the array
// to pivot around.
type fnChoosePivot func(k []int) int

// QuickSort implements the QuickSort algorithm
// and returns the sorted array and the number of comparisons made
func QuickSort(choosePivot fnChoosePivot, k []int) ([]int, int) {
	return nil, -1
}

// Partition paritions an array k around the element at i.
// It then returns slices of the
// input to the left and right of the pivot element.
// NOTE: We take advantage of the nature of slices in
// go so that we don't have to track indexes on these slices.
// instead we can just work from the start to end of each
// and they will operate on the underlying array elements correctly.
func Partition(k []int, p int) ([]int, []int) {
	// put the pivot element in the first position
	lib.Swap(k, 0, p)

	pivot := k[0]
	i := 1
	j := 1

	for j = 1; j < len(k); j++ {
		if k[j] < pivot {
			lib.Swap(k, i, j)
			i++
		}
	}

	// put the pivot element in the right place
	lib.Swap(k, i-1, 0)

	return k[0 : i-1], k[i:len(k)]
}

// ChooseFirstPivot returns the first index of the array
func ChooseFirstPivot(k []int) int {
	return 0
}

// ChooseLastPivot returns the first index of the array
func ChooseLastPivot(k []int) int {
	return len(k) - 1
}

// ChooseMedianOfThreePivot returns the median of three (fill this in)
func ChooseMedianOfThreePivot(k []int) int {
	//TODO implment this
	return 0
}
