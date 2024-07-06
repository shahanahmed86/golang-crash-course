package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to race condition in golang")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var scores = []uint8{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("One R")
		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Two R")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Three R")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Four R")
		mut.RLock()
		fmt.Println("Scores is", scores)
		mut.RUnlock()
	}(wg, mut)

	wg.Wait()
	fmt.Println("Scores are: ", scores)
}
