package stdsort

import (
	"sort"
)

//selection sort sorts data by determining the minimum of the unsorted data,and adding it to the end of the sorted data.
//selectionsort runs in O((n^2)) time complexity due to the need to determine the minimum, On, a total of n times.
func Selectionsort(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		min := i
		for j := i; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
