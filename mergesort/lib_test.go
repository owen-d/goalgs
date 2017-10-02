package mergesort

import (
	"reflect"
	"testing"
)

type testpair struct {
	values []int
	result []int
}

var tests = []testpair{
	{[]int{}, []int{}},
	{[]int{1, 2, 3}, []int{1, 2, 3}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
	{[]int{1, 5, 8, 2, 6, 8, 3}, []int{1, 2, 3, 5, 6, 8, 8}},
	{[]int{1, 3, 2, 5, 4}, []int{1, 2, 3, 4, 5}},
}

func TestMergeSort(t *testing.T) {
	for _, pair := range tests {
		computed := Merge(pair.values, Compare)
		match := reflect.DeepEqual(computed, pair.result)
		if !match {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", computed,
			)
		}
	}
}
