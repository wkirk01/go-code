package main

import (
	"testing"
)

func BenchmarkContainsPairWithSum(b *testing.B) {
	array := []int{10, 15, 3, 7}
	for n := 0; n < b.N; n++ {
		ContainsPairWithSum(array, 17)
	}
}

//go test -run=XXX -bench=.
