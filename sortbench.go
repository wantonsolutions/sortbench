
/*
Package sortbench runs a set of sorting algorithms and times their execution against eachother, eventually the ability to inject your own sorting algorithms will be available

all of the sorting algorithms must sort elements which implement the sort.Interface
*/

package sortbench

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/wantonsolutions/sortbench/sortjob"
	"github.com/wantonsolutions/sortbench/stdsort"
)
const MAX = 100
const NPL = 25


var jobs []*sortjob.SortJob
var totalJobs int
var printjobs bool

func Init() {
	jobs = make([]*sortjob.SortJob,MAX,MAX)
	printjobs = false
}

func PrintJobs(){
	printjobs = true
}

func AddJob(job *sortjob.SortJob) {
	if totalJobs >= MAX {
		return
	}
	jobs[totalJobs] = job
	totalJobs++
}

func Run() {
	for i:=0 ;i<totalJobs;i++{
		stest(jobs[i])
	}
}

func NLgNS(size int, sortType string){
	AddJob(sortjob.New("qwisort",size,stdsort.Qwisort,sortType))
	AddJob(sortjob.New("quicksort",size,stdsort.Quicksort,sortType))
	AddJob(sortjob.New("heapsort",size,stdsort.Heapsort,sortType))
}

func NSqS(size int, sortType string){
	AddJob(sortjob.New("selectionsort",size,stdsort.Selectionsort,sortType))
	AddJob(sortjob.New("insertionsort",size,stdsort.Insertionsort,sortType))
	AddJob(sortjob.New("bubblesort",size,stdsort.Bubblesort,sortType))
}
		
	

//stest is the sorting test function, which takes a sorting algorithm as an argument and then times the execution of that sorting algorithm
func stest(job *sortjob.SortJob) {
	//should generalize this to create any type
	array := make([]int, job.Size, job.Size)
	genRand(array)
	data := IntArray(array)
	//genRand may become very large

	c := make(chan int)
	if printjobs {
		fmt.Println(data.String())
	}
	start := time.Now()
	go func() {
		job.Sort(data)
		c <- 1
	}()
	<-c
	finnish := time.Since(start)
	fmt.Println("Sort: ", job.Name, "\tn: ", job.Size, "\tTime: ", finnish.String())
	if printjobs {
		fmt.Println(data.String(), "\n")
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

//IntArray is a wrappter for an array of integers, it implements the sort.Interface

//test defin
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
