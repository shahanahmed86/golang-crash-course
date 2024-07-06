package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in golang")

	myCh := make(chan int)
	myWg := &sync.WaitGroup{}

	myWg.Add(2)

	// receiving channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		value, isChannelOpen := <-ch
		fmt.Println("Channel is open", isChannelOpen)
		fmt.Println("Channel value", value)
		// fmt.Println("1.Channels", <-ch)
		// fmt.Println("2.Channels", <-ch)
		// fmt.Println("3.Channels", <-ch)
	}(myCh, myWg)

	// sending channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 5
		// ch <- 10
		// ch <- 15
		close(ch)
	}(myCh, myWg)

	myWg.Wait()
}
