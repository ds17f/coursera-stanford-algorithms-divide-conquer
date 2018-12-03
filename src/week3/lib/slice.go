package lib

// CloneSlice creates a clone of the slice
// in a new array and then returns a slice
// of that new array
func CloneSlice(slice []int) []int {
	clone := make([]int, len(slice))
	copy(clone, slice)
	return clone
}

// CompareSlices returns true if all elements of the slices are the same
func CompareSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// AllLess returns true if all values in a are less than val
func AllLess(a []int, val int) bool {
	for _, v := range a {
		if v > val {
			return false
		}
	}
	return true
}

// AllGreater returns true if all values in a are greater than val
func AllGreater(a []int, val int) bool {
	for _, v := range a {
		if v < val {
			return false
		}
	}
	return true

}

// Swap does an in place swap of 2 array elements, i and j in the array a
func Swap(a []int, i int, j int) {
	// swap k[0] and k[p]
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
