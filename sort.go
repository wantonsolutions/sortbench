package main

import(
	"os"
	"fmt"
	"time"
	"math/rand"
	"sort"
	"strconv"
)
const MAX = 100
const NPL = 25
const INSERTFASTER = 16

type IntArray []int

func (ia IntArray) Len() int {
	return len(ia)
}

func (ia IntArray) Less(i, j int) bool {
	return ia[i] < ia[j]
}

func (ia IntArray) Swap (i,j int) {
	ia[i], ia[j] = ia[j], ia[i]
}

func (ia IntArray) String() string{
	str := "["
	for i, e := range ia {
		if i > 0{
			str+= " "
		}
		str += fmt.Sprint(e)
	}
	return str + "]"
}

func main(){
	size,err := strconv.Atoi(os.Args[1])
	if(err != nil){
		print("DYING!!!")
	}/*
	stest(size,"quick", quicksort)
	stest(size,"quwins", qwisort)
	stest(size,"merge", mergesort)
	stest(size,"heap", heapsort)
	stest(size,"insert",insertionsort)
	stest(size,"select",selectionsort)
	*/
	stest(size,"bubble",bubblesort)
	
}

func stest(size int,name string, sort func(sort.Interface)){
	//should generalize this to create any type
	array := make([]int, size, size)
	genRand(array)
	data := IntArray(array)
	//genRand may become very large
	
	c := make(chan int)
	fmt.Println(data.String())
	start := time.Now()
	go func() {
		sort(data)
		c<-1
	}()
	<-c
	finnish := time.Since(start)
	fmt.Println("Sort: ",name,"\tn: ",size,"\tTime: ",finnish.String())
	fmt.Println(data.String(),"\n")
}


func heapsort(array []int) []int{
	children:=2
	maxHeapify(array, children)
	for i:=0;i<len(array)-1;i++{
		tmp:=array[len(array)-1-i]
		array[len(array)-1-i]=array[0]
		array[0] = tmp
		shiftdown(array[0:len(array)-i-1],children)
	}
	return array
}

func shiftdown(array []int,children int) []int{
	node:=0
	swapped := false
	swapIndex := node
	for child:=1;child<=children;child++{
		childnode := node*children + child
		if childnode >= len(array){
			break
		}
		if array[node] < array[childnode]{
			tmp:=array[node]
			array[node]=array[childnode]
			array[childnode]=tmp
			if !swapped {
				swapIndex = childnode
				swapped = true
			}
		}
		if child == children && swapped{
			//reset for the next level
			child = 0
			swapped = false
			node = swapIndex
		}
	}
	return array
}	
	

func maxHeapify(array []int, children int) []int{
	//if the heap has been modifed, it may have broken, reheapify	
	for heaped:=heapify(array, children, 0);heaped;{
		heaped=heapify(array, children, 0)
	}
	return array
}

func heapify(array[] int, children int, index int) bool{
	modified := false
	for child:=1; child <= children; child++{
		if children*index + child < len(array){
			//if any part of the heap is modified retain that
			if(heapify(array,children,children*index+child)){
				modified = true
			}
		}
		if (index -1)/children >= 0 && array[index] > array[(index-1)/children]{
			tmp:= array[(index-1)/children]
			array[(index-1)/children] = array[index]
			array[index] = tmp
			modified = true
		}
	}
	return modified
}



func mergesort(array []int) []int{
	if len(array) <= 1{
		return array
	}
	mergesort(array[0:(len(array)/2)])
	mergesort(array[(len(array)/2):len(array)])
	merge(array)
	return array
}

func merge(array []int) []int{
	f1 := 0			//front of sorted list 1
	f2 := len(array)/2	//front of sorted list 2
	var merged = make([]int, len(array), len(array))
	for i := 0;i<len(array); i++{
		if f1 < len(array)/2 && f2 < len(array){
			if array[f1] < array[f2]{
				merged[i] = array[f1]
				f1++
			} else {
				merged[i] = array[f2]
				f2++
			}
		} else if f1 < len(array)/2{
			merged[i] = array[f1]
			f1++
		} else if f2 < len(array){
			merged[i] = array[f2]
			f2++
		}
	}
	for i := 0;i<len(array); i++{
		array[i] = merged[i]
	}
	return array
}

func selectionsort(array[] int) []int{
	for i:=0;i<len(array);i++{
		min:=i
		for j:=i;j<len(array);j++{
			if array[j] < array [min] {
				min = j
			}
		}
		tmp:= array[i]
		array[i] = array[min]
		array[min] = tmp
	}
	return array
}

func insertionsort(array []int) []int{
	for i:=1;i<len(array);i++{
		for j:=i;array[j]<=array[j-1];j--{
			tmp := array[j]
			array[j] = array[j-1]
			array[j-1] = tmp
			if j == 1{
				break
			}
		}
	}
	return array
}
				

//recursive with call to insertion after min point found
func quicksort(array []int) []int{
	front := 0;
	end := len(array)-1
	if front >= end{
		return array
	}
	pivot := array[end]
	swapIndex := front
	for i:=0;i<end;i++{
		if array[i]<pivot {
			tmp := array[i]
			array[i] = array[swapIndex]
			array[swapIndex] = tmp
			swapIndex++
		}
	}	
	tmp := array[end]	//put the pivot back where it belongs
	array[end] = array[swapIndex]
	array[swapIndex] = tmp
	quicksort(array[front:(swapIndex)])
	quicksort(array[(swapIndex+1):end+1])
	return array
}

func qwisort(array []int) []int{
	front := 0;
	end := len(array)-1
	if front >= end{
		return array
	}
	if (end - front) < INSERTFASTER{
		insertionsort(array)
		return array
	}
	pivot := array[end]
	swapIndex := front
	for i:=0;i<end;i++{
		if array[i]<pivot {
			tmp := array[i]
			array[i] = array[swapIndex]
			array[swapIndex] = tmp
			swapIndex++
		}
	}	
	tmp := array[end]	//put the pivot back where it belongs
	array[end] = array[swapIndex]
	array[swapIndex] = tmp
	quicksort(array[front:(swapIndex)])
	quicksort(array[(swapIndex+1):end+1])
	return array
}

func bubblesort(data sort.Interface){
	for i:=0; i<data.Len(); i++{
		for j :=1 ;j<data.Len(); j++{
			if data.Less(j,j-1){
				data.Swap(j,j-1)
			}
		}
	}
}
			

func genRand(array IntArray){
	r := rand.New(rand.NewSource(69))
	for i:=0; i<len(array); i++{
		array[i] = r.Int()%MAX
	}
}

func printArray(array []int) {
	for i:=0; i<len(array); i++{
		fmt.Print(array[i],",")
		if i%NPL == 0 && i!= 0{
			fmt.Println("")
		}
	}
	print("\n\n")
}

	
	
