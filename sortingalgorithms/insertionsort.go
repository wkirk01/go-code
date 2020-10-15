package main

import (
	"fmt"
)

func main() {
	array := []int{5, 4, 3, 2, 1}
	fmt.Println("Unsorted array", array)
	sortedArray := BubbleSort(array)
	fmt.Println("Sorted array", sortedArray)
}

//BubbleSort ...
func BubbleSort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := 0; j < (length - 1 - i); j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

//InsertionSort ...
func InsertionSort(array []int) []int {
	length := len(array)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}
