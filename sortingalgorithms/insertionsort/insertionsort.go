package main

import (
	"fmt"
)

func main() {
	array := []int{5, 4, 3, 2, 1}
	sortedArray := InsertionSort(array)
	fmt.Println(sortedArray)
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
