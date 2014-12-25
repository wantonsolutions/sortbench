/*
Package sortbench runs a set of sorting algorithms and times their execution against eachother, eventually the ability to inject your own sorting algorithms will be available

all of the sorting algorithms must sort elements which implement the sort.Interface
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

const MAX = 100
const NPL = 25
const INSERTFASTER = 16

//IntArray is a wrappter for an array of integers, it implements the sort.Interface
type IntArray []int

func (ia IntArray) Len() int {
	return len(ia)
}

func (ia IntArray) Less(i, j int) bool {
	return ia[i] < ia[j]
}

func (ia IntArray) Swap(i, j int) {
	ia[i], ia[j] = ia[j], ia[i]
}

func (ia IntArray) String() string {
	str := "["
	for i, e := range ia {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(e)
	}
	return str + "]"
}

//man validates the command line arguments and controls the execution of the sorting algorithms
func main() {
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		print("DYING!!!")
	} /*
		stest(size,"merge", mergesort)
	*/
	stest(size, "quwins", qwisort)
	stest(size, "quick", quicksort)
	stest(size, "heap", heapsort)
	stest(size, "select", selectionsort)
	stest(size, "insert", insertionsort)
	stest(size, "bubble", bubblesort)

}

//stest is the sorting test function, which takes a sorting algorithm as an argument and then times the execution of that sorting algorithm
func stest(size int, name string, sort func(sort.Interface)) {
	//should generalize this to create any type
	array := make([]int, size, size)
	genRand(array)
	data := IntArray(array)
	//genRand may become very large

	c := make(chan int)
	//fmt.Println(data.String())
	start := time.Now()
	go func() {
		sort(data)
		c <- 1
	}()
	<-c
	finnish := time.Since(start)
	fmt.Println("Sort: ", name, "\tn: ", size, "\tTime: ", finnish.String())
	//fmt.Println(data.String(), "\n")
}

//heapsort in the calling function for the heapsort algorithm, which takes as an argument a set of data.
func heapsort(data sort.Interface) {
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

func mergesort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	mergesort(array[0:(len(array) / 2)])
	mergesort(array[(len(array) / 2):len(array)])
	merge(array)
	return array
}

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
//qwisort sorts data based on both quicksort and insertion sort, the data is partitioned by quicksort untill a limit is reached, then the data is sorted by insertion sort.
//qwisort runs in O(nlogn) time with O 1 space complexity
//qwisort preforms on average better then quicksort due the the inherit good preformance of insertion sort on small sets
func qwisort(data sort.Interface) {
	qwisortH(data, 0, data.Len())
}

//qwisortH is the helper for qwisort, using a range of values [a,b] allows the helper function to run recursivly
func qwisortH(data sort.Interface, a int, b int) {
	if a >= b-1 {
		return
	}
	if (b - a) < INSERTFASTER {
		insertionsortH(data, a, b)
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

//quicksort sorts data by determining a pivot, then partitioning data based on its value relative to the pivot.
//each quicksorted partition is then respective quicksorted, with a recursive base case occuring when the partition is of size [0,1]
//quicksort run is O(nlogn) time and has a space complexity of O 1
func quicksort(data sort.Interface) {
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


//selection sort sorts data by determining the minimum of the unsorted data,and adding it to the end of the sorted data.
//selectionsort runs in O((n^2)) time complexity due to the need to determine the minimum, On, a total of n times.
func selectionsort(data sort.Interface) {
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

//insertionsort sorts data by taking an element from the unsorted set of data, and swaping it into its correct position within the sorted set of data.
//insertionsort runs in O(n^2) time but has a best case complexity of On on sorted data.
func insertionsort(data sort.Interface) {
	insertionsortH(data, 0, data.Len())
}

//insertionsortH is the helper function for insertion sort, which allows it to be run over an interval of data.
func insertionsortH(data sort.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
			if j == 1 {
				break
			}
		}
	}
}

//bubblesort sorts data by itterating through the data once for each element swapping that element with its neighbour if it is out of order.
//bubblesort runs in O(n^2) time because it strictly compares every element with every other element, it has a space complexity of O(1)
//bubblesort is a bad sorting algorithm and is used as a benchmark for how bad any other comparison sort algorithm should preform
func bubblesort(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		for j := 1; j < data.Len(); j++ {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
//genRand creates a collection of random valued integers within an intArrray
func genRand(array IntArray) {
	r := rand.New(rand.NewSource(69))
	for i := 0; i < len(array); i++ {
		array[i] = r.Int() % MAX
	}
}

//printArray prints each element in an integer array
func printArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], ",")
		if i%NPL == 0 && i != 0 {
			fmt.Println("")
		}
	}
	print("\n\n")
}
