package algorithms

import "math"

// CountInversions returns a count of the number of inversions
// (out of sort order elements) in an array of integers specified
// as the input variable
func CountInversions(n []int) int {
	_, inversionCount := mergeSort(n)
	return inversionCount
}

// mergeSort performs a merge sort on the input array n
// and returns the sorted input array and the number of inversions
// found during the sort
func mergeSort(n []int) ([]int, int) {
	if len(n) == 1 {
		return n, 0
	}
	left, right := splitArray(n)
	leftSorted, leftInv := mergeSort(left)
	rightSorted, rightInv := mergeSort(right)
	mergeSorted, inversions := mergeAndCount(leftSorted, rightSorted)
	return mergeSorted, (leftInv + rightInv + inversions)
}

// mergeAndCount merges the left and right arrays while counting the
// number of inversions.  It returns the sorted array and the count
// of inversions
func mergeAndCount(left []int, right []int) ([]int, int) {
	inversions := 0
	i := 0
	j := 0
	var merged []int

	for i < len(left) && j < len(right) {
		if right[j] >= left[i] {
			// if right is larger we're good
			// copy the left[i] to the output
			// it is in proper order
			merged = append(merged, left[i])
			i++
		} else if right[j] < left[i] {
			// if right is smaller then it is out of order
			// since both arrays (left, right) are sorted
			// then left[i] is the smallest element remaining in left
			// therefore all remaining left elements are also
			// larger than right[j] and they are all inversions.
			// we can count the remaining elements of left
			// and add them to our inversion total
			inversions += len(left) - i
			merged = append(merged, right[j])
			j++
		}
	}

	// take care of stragglers
	// NOTE: check out the variadic function and the spread operator
	//       this reminds me of javascript's spread operator
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged, inversions
}

// splitArray splits the input array in half
// with a right side bias when len
func splitArray(n []int) ([]int, []int) {
	mid := int(math.Floor(float64(len(n) / 2)))
	left := n[:mid]
	right := n[mid:]
	return left, right
}
