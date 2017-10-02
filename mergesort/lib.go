package mergesort

import (
	"math"
)

/*
Ok, we want to take an array of int32s
*/
func Compare(a, b int) int {
	return a - b
}

func Merge(list []int, comparator func(int, int) int) []int {
	/*
		if ln = 1, return itself
		if ln > 1, calc left side & right side recursively
		this will yield two sorted slices.
		we iterate over the slices, keeping a ptr on an element in both slices, starting from idx 0.
		we Merge the results into a result slice, finally returning that slice upwards.
	*/
	l := len(list)
	if l <= 1 {
		return list
	}

	midpoint := int(math.Floor(float64(l) / 2))

	left := Merge(list[:midpoint], comparator)
	right := Merge(list[midpoint:], comparator)

	// allocate a result slice w/ capacity for all elements
	result := make([]int, l)

	j := 0
	k := 0
	for i := 0; i < l; i++ {
		//check out of bounds on ptrs.
		// if we've consumed left list, can safely append the remaining right list to result
		if j == len(left) {
			copy(result[i:], right[k:])
			break
		}
		// conversely, if we've consumed right list, can do the opposite
		if k == len(right) {
			copy(result[i:], left[j:])
			break
		}

		//compare
		if comparator(left[j], right[k]) <= 0 {
			result[i] = left[j]
			j++
		} else {
			result[i] = right[k]
			k++
		}
	}

	return result
}
