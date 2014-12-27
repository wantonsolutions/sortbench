package stdsort

import (
	"sort"
)

const INSERTFASTER = 16

//qwisort sorts data based on both quicksort and insertion sort, the data is partitioned by quicksort untill a limit is reached, then the data is sorted by insertion sort.
//qwisort runs in O(nlogn) time with O 1 space complexity
//qwisort preforms on average better then quicksort due the the inherit good preformance of insertion sort on small sets
func Qwisort(data sort.Interface) {
	qwisortH(data, 0, data.Len())
}

//qwisortH is the helper for qwisort, using a range of values [a,b] allows the helper function to run recursivly
func qwisortH(data sort.Interface, a int, b int) {
	if a >= b-1 {
		return
	}
	if (b - a) < INSERTFASTER {
		InsertionsortR(data, a, b)
		return
	}
	swapIndex := a
	for i := a; i < b-1; i++ {
		if data.Less(i, b-1) {
			data.Swap(i, swapIndex)
			swapIndex++
		}
	}
	data.Swap(swapIndex, b-1)
	qwisortH(data, a, swapIndex)
	qwisortH(data, swapIndex+1, b)
}
