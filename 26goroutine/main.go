package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

var signals = []string{"test"}

func main() {
	fmt.Println("Welcome to goroutines in golang")
	// go greeter("Hello")
	// greeter("World")

	websites := []string{
		"https://test.api.deeneeapp.com/healthcheck",
		"https://v3.api.deeneeapp.com/healthcheck",
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
		"https://youtube.com",
	}

	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("Here is the signals", signals)
}

func getStatusCode(endpoint string) int8 {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	mut.Lock()
	defer mut.Unlock()
	signals = append(signals, endpoint)

	fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
	return int8(res.StatusCode)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
