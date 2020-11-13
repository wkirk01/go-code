package main

import (
	"github.com/wkirk01/tasks"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go worker(&wg, tasks.TaskOne)
	go worker(&wg, tasks.TaskTwo)

	wg.Wait()
}

func worker(wg *sync.WaitGroup, f func()) {
	defer wg.Done()
	f()
}
