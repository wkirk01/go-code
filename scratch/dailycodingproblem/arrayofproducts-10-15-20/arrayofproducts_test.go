package main

import (
	"testing"
)

func BenchmarkBuildArrayOfProducts(b *testing.B) {
	array := []int{1, 2, 3, 4, 5}
	for n := 0; n < b.N; n++ {
		BuildArrayOfProducts(array)
	}
}

//go test -run=XXX -bench=.
