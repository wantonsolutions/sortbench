package stdsort

import (
	"sort"
)
//heapsort in the calling function for the heapsort algorithm, which takes as an argument a set of data.
func Heapsort(data sort.Interface) {
	heapsortH(data, 0, data.Len(), 2)
}

//heapsortH is the helper function for heapsort, it takes a variable number of children, and length of the dataset.
//heapsortH sorts a dataset by first constructing a heap out of the data, then continuously removing the largest element and reheapifying the heap
//heapsortH sorts in OnLogn time complexit and O 1 space complexity
func heapsortH(data sort.Interface, a, b, children int) {
	//heapify
	for i := (b - 1) / 2; i >= a; i-- {
		shiftdown(data, i, b, children)
	}
	for i := b - 1; i >= a; i-- {
		data.Swap(i, a)
		shiftdown(data, a, i, children)
	}
}

//shiftdown assumes that the data within the range [(a+1),b] is a valid heapand shifts the contents of position a into a valid position, reconstructing the heap
//shiftdown runs in Ologn time and requires O 1 space complexity 
func shiftdown(data sort.Interface, a, b, children int) {
	node := a
	max := node
	//traverse the heap, considering each child at each level
	for child := 1; child <= children; child++ {
		childnode := (node-a)*children + child + a
		//break if out of range
		if childnode >= b {
			break
		}
		//determine the largest child
		if data.Less(max, childnode) {
			max = childnode
		}
		//swap the head node with the largest child
		if child == children && max > node {
			child = 0
			data.Swap(node, max)
			node = max
		}
	}
}
