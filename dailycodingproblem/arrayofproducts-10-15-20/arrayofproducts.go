/*
Given an array of integers, return a new array such that each element at index i of the new array
is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/

package main

import (
	"fmt"
)

func main() {
	arrayOne := []int{1, 2, 3, 4, 5}
	fmt.Println(BuildArrayOfProducts(arrayOne)) //expected output [120, 60, 40, 30, 24]

	arrayTwo := []int{3, 2, 1}
	fmt.Println(BuildArrayOfProducts(arrayTwo)) //expected output [2, 3, 6]
}

/*
The idea is to construct two arrays (in the case for 4 elements)

[              1,         a[0],    a[0]*a[1],    a[0]*a[1]*a[2] ]
[ a[1]*a[2]*a[3],    a[2]*a[3],         a[3],                 1 ]

Both of which can be done in O(n) by starting at the left and right edges respectively.

Then multiplying the two arrays element by element gives the required result
*/

//BuildArrayOfProducts ...
func BuildArrayOfProducts(array []int) []int {
	length := len(array)
	products := make([]int, length)

	//get the products below the current index
	p := 1
	for i := 0; i < length; i++ {
		products[i] = p
		p *= array[i]
	}

	//get the products above the curent index
	p = 1
	for i := length - 1; i >= 0; i-- {
		products[i] *= p
		p *= array[i]
	}

	return products
}
