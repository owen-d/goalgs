package heapsort

import (
	"math"
)

// siftDown will swap a parent & the first swappable child, optionally recursing into the subtree with the root where the aforementioned child was.
func siftDown(col []int, x int, recur bool) {
	a, b := children(x)

	if inBounds(col, a) && col[x] < col[a] {
		if inBounds(col, b) && col[a] < col[b] {
			swap(col, x, b)
			if recur {
				siftDown(col, b, recur)
			}
		} else {
			swap(col, x, a)
			if recur {
				siftDown(col, a, recur)
			}
		}
	} else if inBounds(col, b) && col[x] < col[b] {
		swap(col, x, b)
		if recur {
			siftDown(col, b, recur)
		}
	}
}

func inBounds(col []int, xs ...int) bool {
	lastIdx := len(col) - 1
	for _, x := range xs {
		if x > lastIdx {
			return false
		}
	}
	return true
}

func swap(col []int, x, y int) {
	col[x], col[y] = col[y], col[x]
}

func children(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func parent(i int) int {
	return int(math.Floor(float64(i-1) / 2))
}

func heapify(col []int) {
	heapEnd := 0
	for heapEnd < len(col) {
		// assume we're adding the next item into the heap from the bottom
		start := parent(heapEnd)
		// iterate up to the beginning of the heap from the parent of the newly added elem,
		// making sure that the subheaps maintain the heap guarantees.
		for start >= 0 {
			siftDown(col[:heapEnd+1], start, false)
			start = parent(start)
		}

		heapEnd++
	}

}

//build a heap in place out of array
//swap first item in heap w/ last array index, & slice out that idx from the heap & recalculating heap.
//this will grow the sorted area behind the shrinking heap
func heapsort(col []int) []int {
	l := len(col)
	if l == 0 {
		return col
	}
	heapify(col)
	for i := l - 1; i > 0; i-- {
		swap(col, 0, i)
		siftDown(col[:i], 0, true)
	}
	return col
}
