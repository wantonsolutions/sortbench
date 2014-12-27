/*
Package sortbench runs a set of sorting algorithms and times their execution against eachother, eventually the ability to inject your own sorting algorithms will be available

all of the sorting algorithms must sort elements which implement the sort.Interface
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"github.com/wantonsolutions/sortbench/stdsort"
	"github.com/wantonsolutions/sortbench/sortjob"
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

//test defin
	

//man validates the command line arguments and controls the execution of the sorting algorithms
func main() {
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		print("DYING!!!")
	} /*
		stest(size,"merge", mergesort)
	*/
	qwisort := sortjob.New("qwisort",size,stdsort.Qwisort,"random")	
	quicksort := sortjob.New("quicksort",size,stdsort.Quicksort,"random")	
	heapsort := sortjob.New("heapsort",size,stdsort.Heapsort,"random")	
	selectionsort := sortjob.New("selectsort",size,stdsort.Selectionsort,"random")	
	insertionsort := sortjob.New("insertionsort",size,stdsort.Insertionsort,"random")	
	bubblesort := sortjob.New("bubblesort",size,stdsort.Bubblesort,"random")	
	stest(qwisort)
	stest(quicksort)
	stest(heapsort)
	stest(selectionsort)
	stest(insertionsort)
	stest(bubblesort)

}

//stest is the sorting test function, which takes a sorting algorithm as an argument and then times the execution of that sorting algorithm
func stest(job *sortjob.SortJob) {
	//should generalize this to create any type
	array := make([]int, job.Size, job.Size)
	genRand(array)
	data := IntArray(array)
	//genRand may become very large

	c := make(chan int)
	//fmt.Println(data.String())
	start := time.Now()
	go func() {
		job.Sort(data)
		c <- 1
	}()
	<-c
	finnish := time.Since(start)
	fmt.Println("Sort: ", job.Name, "\tn: ", job.Size, "\tTime: ", finnish.String())
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
