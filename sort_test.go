package main

import "testing"

func TestSmallList(t *testing.T) {
	testInput := []string{"red", "gold", "black", "maroon"}
	correctOutput := []string{"black", "gold", "maroon", "red"}

	output := bubbleSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red] , but bubbleSort returned: ", output)
	}

	output = insertionSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red], but insertionSort returned: ", output)
	}

	output = selectionSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red] , but selectionSort returned: ", output)
	}

	output = mergeSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red] , but mergeSort returned: ", output)
	}

	output = heapSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red], but heapSort returned: ", output)
	}

}

func TestReverseOrderedList(t *testing.T) {
	testInput := []string{"red", "maroon", "gold", "black"}
	correctOutput := []string{"black", "gold", "maroon", "red"}

	output := bubbleSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red] , but bubbleSort returned: ", output)
	}

	output = insertionSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red], but insertionSort returned: ", output)
	}

	output = selectionSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red] , but selectionSort returned: ", output)
	}

	output = mergeSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red] , but mergeSort returned: ", output)
	}

	output = heapSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [black gold maroon red], but heapSort returned: ", output)
	}
}

func TestMixedCaseList(t *testing.T) {
	testInput := []string{"red", "gold", "Black", "maroon", "Orange"}
	correctOutput := []string{"Black", "gold", "maroon", "Orange", "red"}

	output := bubbleSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [Black gold maroon Orange red] , but bubbleSort returned: ", output)
	}

	output = insertionSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [Black gold maroon Orange red], but insertionSort returned: ", output)
	}

	output = selectionSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [Black gold maroon Orange red] , but selectionSort returned: ", output)
	}

	output = mergeSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [Black gold maroon Orange red] , but mergeSort returned: ", output)
	}

	output = heapSort(testInput)
	if !isEqual(output, correctOutput) {
		t.Error("Expected [Black gold maroon Orange red], but heapSort returned: ", output)
	}
}

func isEqual(a, b []string) bool {
	for i, value := range a {
		if value != b[i] {
			return false
		}
	}
	return true
}
