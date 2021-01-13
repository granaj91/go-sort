package main
import (
	"fmt"
	"time"
)

func main(){
	var a = []int{4,9,5,3,2}
	bubble_sort(a)
	insertion_sort(a)
	selection_sort(a)
	merge_sort(a)
	heap_sort(a)
}

func bubble_sort(arr []int){
	start := time.Now()
	for i := len(arr) - 1; i >= 1; i--{
		for j := 0; j <= i - 1; j++{
			if arr[j] > arr[j+1]{
				var temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Bubble Sort: ", arr)
	fmt.Println("Execution time: ", elapsed)
}

func insertion_sort(arr []int){
	start := time.Now()
	for i := 1; i < len(arr); i++{
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp{
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	elapsed := time.Since(start)
	fmt.Println("Insertion Sort: ", arr)
	fmt.Println("Execution time: ", elapsed)
}

func selection_sort(arr []int){
	start := time.Now()
	for i := len(arr) - 1; i >= 1; i--{
		t := 0
		for j := 1; j <= i; j++{
			if arr[j] > arr[t]{
				t = j
			}
		}
		temp := arr[t]
		arr[t] = arr[i]
		arr[i] = temp
	}
	elapsed := time.Since(start)
	fmt.Println("Selection Sort: ", arr)
	fmt.Println("Execution time: ", elapsed)
}

func merge_sort(arr []int){
	start := time.Now()
	newArr := merge_sort_aux(arr)
	elapsed := time.Since(start)
	fmt.Println("Merge Sort: ", newArr)
	fmt.Println("Execution time: ", elapsed)
}

func merge_sort_aux(arr []int) []int{
	n := len(arr)
	if n <= 1 {
		return arr
	}
	leftArr := merge_sort_aux(arr[:n/2])
	rightArr := merge_sort_aux(arr[n/2:])
	return merge(leftArr, rightArr)
}

func merge(left, right []int) []int{
	ln := len(left)
	rn := len(right)
	n := ln + rn
	sorted :=make([]int, n)
	i := 0
	j := 0
	k := 0

	for i < ln && j < rn {
		if left[i] < right[j]{
			sorted[k] = left[i]
			i++
		}else {
			sorted[k] = right[j]
			j++
		}
		k++
	}

	for i < ln {
		sorted[k] = left[i]
		i++
		k++
	}
	for j < rn {
		sorted[k] = right[j]
		j++
		k++
	}
	return sorted
}

func heap_sort(arr []int){
	start := time.Now()
	size := len(arr)
	build_max_heap(arr)
	for i := len(arr) - 1; i > 1; i-- {
		temp := arr[i]
		arr[i] = arr[1]
		arr[1] = temp
		size--
		max_heapify(arr[:size], 1)
	}
	elapsed := time.Since(start)

	fmt.Println("Heap Sort: ", arr)
	fmt.Println("Execution time: ", elapsed)
}

func build_max_heap(arr []int){
	for i := len(arr)/2; i > 0; i-- {
		max_heapify(arr, i)
	}
}

func parent(index int) int {
	return index/2
}

func left(index int) int {
	return 2 * index
}

func right(index int) int {
	return 2*index + 1
}

func max_heapify(arr []int, i int) []int {
	n := len(arr) - 1
	l := left(i)
	r := right(i)
	largest := 0

	if l <= n && arr[l] > arr[i]{
		largest = l
	}else{
		largest = i
	}

	if r <= n && arr[r] > arr[largest]{
		largest = r
	}

	if largest != i {
		temp := arr[i]
		arr[i] = arr[largest]
		arr[largest] = temp
		max_heapify(arr, largest)
	}

	return arr
}
