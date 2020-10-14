package main

import (
	"bkirk/scratch/tasks/tasks"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go worker(&wg, tasks.TaskOne)

	wg.Add(1)
	go worker(&wg, tasks.TaskTwo)

	wg.Wait()
}

func worker(wg *sync.WaitGroup, f func()) {
	defer wg.Done()
	f()
}
