package main

import (
	"bkirk/scratch/tasks/tasks"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	worker(&wg, tasks.TaskOne)
	worker(&wg, tasks.TaskTwo)

	wg.Wait()
}

func worker(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	defer wg.Done()
	f()
}
