
package stdsort

import(
	"sort"
)

//bubblesort sorts data by itterating through the data once for each element swapping that element with its neighbour if it is out of order.
//bubblesort runs in O(n^2) time because it strictly compares every element with every other element, it has a space complexity of O(1)
//bubblesort is a bad sorting algorithm and is used as a benchmark for how bad any other comparison sort algorithm should preform
func Bubblesort(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		for j := 1; j < data.Len(); j++ {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
