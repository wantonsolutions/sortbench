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
	"github.com/wantonsolutions/sortbench/stdsort"
)

const MAX = 100
const NPL = 25

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
	stest(size, "quwins", stdsort.Qwisort)
	stest(size, "quick", stdsort.Quicksort)
	stest(size, "heap", stdsort.Heapsort)
	stest(size, "select", stdsort.Selectionsort)
	stest(size, "insert", stdsort.Insertionsort)
	stest(size, "bubble", stdsort.Bubblesort)

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
