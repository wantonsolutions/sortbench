package stdsort

//Mergesort sorts an array of integers, the sort is preformed by partitioning the range of values in half recursivly, untill the base case of 0 or 1 elements has been reached.
//partions of size 0 or 1 are considered sorted and can be merged with other sorted partitions, this is done with every aggregated merged partition
//Mergesort runs in O(nlogn) time with a space complexity of O(n)
func Mergesort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	Mergesort(array[0:(len(array) / 2)])
	Mergesort(array[(len(array) / 2):len(array)])
	merge(array)
	return array
}

//Merge takes a slice as an argument, the slice is assumed to contain two sorted lists partitioned by the middle element
func merge(array []int) []int {
	f1 := 0              //front of sorted list 1
	f2 := len(array) / 2 //front of sorted list 2
	var merged = make([]int, len(array), len(array))
	for i := 0; i < len(array); i++ {
		if f1 < len(array)/2 && f2 < len(array) {
			if array[f1] < array[f2] {
				merged[i] = array[f1]
				f1++
			} else {
				merged[i] = array[f2]
				f2++
			}
		} else if f1 < len(array)/2 {
			merged[i] = array[f1]
			f1++
		} else if f2 < len(array) {
			merged[i] = array[f2]
			f2++
		}
	}
	for i := 0; i < len(array); i++ {
		array[i] = merged[i]
	}
	return array
}
