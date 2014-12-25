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
	stest(size,"merge", mergesort)
	*/
	stest(size,"quwins", qwisort)
	stest(size,"quick", quicksort)
	stest(size,"heap", heapsort)
	stest(size,"select",selectionsort)
	stest(size,"insert",insertionsort)
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

func heapsort(data sort.Interface){
	heapsortH(data,0,data.Len(),2)
}

func heapsortH(data sort.Interface, a, b, children int){
	//heapify
	for i:= (b-1)/2;i>=a;i--{
		shiftdown(data,i,b,children)
	}
	for i:=b-1;i>=a;i-- {
		data.Swap(i,a)
		shiftdown(data,a,i,children)
	}
}

func shiftdown(data sort.Interface, a, b, children int){
	node:=a
	max:=node
	for child:=1;child<=children;child++{
		childnode:=(node-a)*children +child +a
		//fmt.Println(child, ",")
		if childnode>=b{
			break
		}
		if data.Less(max,childnode){
			max = childnode
		}
		if child == children && max > node{
			child =0
			data.Swap(node,max)
			node = max
		}
	}
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

				
func qwisort(data sort.Interface){
	qwisortH(data, 0, data.Len())
}

func qwisortH(data sort.Interface, a int, b int){
	if a >= b-1{
		return
	}
	if (b - a) < INSERTFASTER{
		insertionsortH(data,a,b)
		return
	}
	swapIndex := a
	for i:=a;i<b-1;i++{
		if data.Less(i,b-1){
			data.Swap(i,swapIndex)
			swapIndex++
		}
	}
	data.Swap(swapIndex,b-1)
	qwisortH(data,a,swapIndex)
	qwisortH(data,swapIndex+1,b)	
}

//recursive with call to insertion after min point found
func quicksort(data sort.Interface){
	quicksortH(data, 0, data.Len())
}

func quicksortH(data sort.Interface, a int, b int){
	if a >= b-1{
		return
	}
	swapIndex := a
	for i:=a;i<b-1;i++{
		if data.Less(i,b-1){
			data.Swap(i,swapIndex)
			swapIndex++
		}
	}
	data.Swap(swapIndex,b-1)
	quicksortH(data,a,swapIndex)
	quicksortH(data,swapIndex+1,b)	
}

func selectionsort(data sort.Interface){
	for i:=0;i<data.Len();i++{
		min:=i
		for j:=i;j<data.Len();j++{
			if data.Less(j,min){
				min = j
			}
		}
		data.Swap(i,min)
	}
}

func insertionsort(data sort.Interface){
	insertionsortH(data,0,data.Len())
}

func insertionsortH(data sort.Interface, a, b int){
	for i:=a+1;i<b;i++{
		for j:=i;data.Less(j,j-1);j--{
			data.Swap(j,j-1)
			if j == 1{
				break
			}
		}
	}
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

	
	
