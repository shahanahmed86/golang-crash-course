package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://v3.api.deeneeapp.com/queued-jobs"

func main() {
	fmt.Println("Welcome to web requests in golang")

	resp, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", resp)
	defer resp.Body.Close() // callers' responsibility to close the connection

	databyte, err := io.ReadAll(resp.Body)
	checkNilErr(err)

	fmt.Println("Content is:", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Server is about to panic")
		panic(err)
	}
}
