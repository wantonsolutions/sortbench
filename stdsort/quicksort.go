package stdsort

import (
	"sort"
)

//quicksort sorts data by determining a pivot, then partitioning data based on its value relative to the pivot.
//each quicksorted partition is then respective quicksorted, with a recursive base case occuring when the partition is of size [0,1]
//quicksort run is O(nlogn) time and has a space complexity of O 1
func Quicksort(data sort.Interface) {
	quicksortH(data, 0, data.Len())
}

//quicksortH is the helper fucnction for qucksort, the range [a,b] allows the function to be called in a recursive manner.
func quicksortH(data sort.Interface, a int, b int) {
	if a >= b-1 {
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
	quicksortH(data, a, swapIndex)
	quicksortH(data, swapIndex+1, b)
}
