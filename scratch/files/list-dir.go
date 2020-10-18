package main

import (
	"fmt"
	"github.com/wkirk01/utilities"
)

func main() {
	files := utilities.GetDirContents("/Users/billykirk/Downloads/UPHPRD/Final Files")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
