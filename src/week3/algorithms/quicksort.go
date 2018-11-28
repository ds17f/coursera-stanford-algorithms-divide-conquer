package algorithms

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
