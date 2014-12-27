package stdsort

import (
	"sort"
)

//insertionsort sorts data by taking an element from the unsorted set of data, and swaping it into its correct position within the sorted set of data.
//insertionsort runs in O(n^2) time but has a best case complexity of On on sorted data.
func Insertionsort(data sort.Interface) {
	InsertionsortR(data, 0, data.Len())
}

//insertionsortH is the helper function for insertion sort, which allows it to be run over an interval of data.
func InsertionsortR(data sort.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
			if j == 1 {
				break
			}
		}
	}
}
