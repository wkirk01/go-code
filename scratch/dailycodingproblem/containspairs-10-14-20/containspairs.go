/*

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

*/

package main

import (
	"fmt"
)

func main() {
	array := []int{10, 15, 3, 7}
	fmt.Println(ContainsPairWithSum(array, 17))
	fmt.Println(ContainsPairWithSum(array, 16))
}

//ContainsPairWithSum ...
func ContainsPairWithSum(array []int, k int) bool {
	hashmap := map[int]int{}
	for _, value := range array {
		hashmap[value] = 0
		if _, present := hashmap[k-value]; present {
			return true
		}
	}
	return false
}
