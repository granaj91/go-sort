package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	eventNames, err := getHackathonEvents("https://mlh.io/seasons/2021/events")

	if err != nil {
		log.Println(err)
	}

	bubbleSort(eventNames)
	insertionSort(eventNames)
	selectionSort(eventNames)
	mergeSort(eventNames)
	sorted := heapSort(eventNames)
	fmt.Println("\nSorted List: ")
	printList(sorted)
}

func deepCopy(original, copy []string) {
	for i, v := range original {
		copy[i] = v
	}
}

func bubbleSort(events []string) []string {
	arr := make([]string, len(events))
	deepCopy(events, arr)
	start := time.Now()
	for i := len(arr) - 1; i >= 1; i-- {
		for j := 0; j <= i-1; j++ {
			if strings.ToLower(arr[j]) > strings.ToLower(arr[j+1]) {
				var temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Bubble sort execution time: ", elapsed)
	return arr
}

func insertionSort(events []string) []string {
	arr := make([]string, len(events))
	deepCopy(events, arr)
	start := time.Now()
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && strings.ToLower(arr[j]) > strings.ToLower(temp) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	elapsed := time.Since(start)
	fmt.Println("Insertion sort execution time: ", elapsed)
	return arr
}

func selectionSort(events []string) []string {
	arr := make([]string, len(events))
	deepCopy(events, arr)
	start := time.Now()
	for i := len(arr) - 1; i >= 1; i-- {
		t := 0
		for j := 1; j <= i; j++ {
			if strings.ToLower(arr[j]) > strings.ToLower(arr[t]) {
				t = j
			}
		}
		temp := arr[t]
		arr[t] = arr[i]
		arr[i] = temp
	}
	elapsed := time.Since(start)
	fmt.Println("Selection sort execution time: ", elapsed)
	return arr
}

func mergeSort(events []string) []string {
	arr := make([]string, len(events))
	deepCopy(events, arr)
	start := time.Now()
	newArr := mergeSortAux(arr)
	elapsed := time.Since(start)
	fmt.Println("Merge sort execution time: ", elapsed)
	return newArr
}

func mergeSortAux(arr []string) []string {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	leftArr := mergeSortAux(arr[:n/2])
	rightArr := mergeSortAux(arr[n/2:])
	return merge(leftArr, rightArr)
}

func merge(left, right []string) []string {
	ln := len(left)
	rn := len(right)
	n := ln + rn
	sorted := make([]string, n)
	i := 0
	j := 0
	k := 0

	for i < ln && j < rn {
		if strings.ToLower(left[i]) < strings.ToLower(right[j]) {
			sorted[k] = left[i]
			i++
		} else {
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

func heapSort(events []string) []string {
	arr := make([]string, len(events))
	deepCopy(events, arr)
	start := time.Now()
	size := len(arr)
	buildMaxHeap(arr)
	for i := len(arr) - 1; i >= 1; i-- {
		temp := arr[i]
		arr[i] = arr[0]
		arr[0] = temp
		size--
		maxHeapify(arr[:size], 0)
	}
	elapsed := time.Since(start)
	fmt.Println("Heap sort execution time: ", elapsed)
	return arr
}

func buildMaxHeap(arr []string) {
	for i := len(arr) / 2; i > 0; i-- {
		maxHeapify(arr, i)
	}
}

func parent(index int) int {
	return index / 2
}

func left(index int) int {
	return 2 * index
}

func right(index int) int {
	return 2*index + 1
}

func maxHeapify(arr []string, i int) []string {
	n := len(arr) - 1
	l := left(i)
	r := right(i)
	largest := 0

	if l <= n && strings.ToLower(arr[l]) > strings.ToLower(arr[i]) {
		largest = l
	} else {
		largest = i
	}

	if r <= n && strings.ToLower(arr[r]) > strings.ToLower(arr[largest]) {
		largest = r
	}

	if largest != i {
		temp := arr[i]
		arr[i] = arr[largest]
		arr[largest] = temp
		maxHeapify(arr, largest)
	}

	return arr
}

func printList(list []string) {
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}
