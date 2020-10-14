package main

import (
	"fmt"
)

func main() {
	s := []int{1}
	fmt.Println(s)
	changeZeroIndex(s)
	fmt.Println(s)
}

func changeZeroIndex(s []int) {
	s[0] = 2
}
