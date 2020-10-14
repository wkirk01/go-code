package main

import (
	"bkirk/utilities"
	"fmt"
)

func main() {
	files := utilities.GetDirContents("/Users/billykirk/Downloads/UPHPRD/Final Files")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
