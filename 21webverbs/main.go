package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verbs in golang")
	PerformGetRequest("/get")
}

const baseUrl string = "http://localhost:3000"

func PerformGetRequest(route string) {
	myUrl := baseUrl + route

	resp, err := http.Get(myUrl)
	checkNilErr(err)

	defer resp.Body.Close()

	fmt.Println("Status code is:", resp.StatusCode)
	fmt.Println("Content length is:", resp.ContentLength)

	var responseString strings.Builder
	content, contentErr := io.ReadAll(resp.Body)
	checkNilErr(contentErr)

	responseString.Write(content)

	fmt.Println("Data inside the response is:", responseString.String())
}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Server is about to panic")
		panic(err)
	}
}
