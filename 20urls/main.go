package main

import (
	"fmt"
	"net/url"
)

const website string = "https://api.github.com/users?since=100&per_page=5"

func main() {
	fmt.Println("Welcome to URLs in golang")

	result, err := url.Parse(website)
	checkNilErr(err)

	fmt.Println("Scheme", result.Scheme)
	fmt.Println("Host", result.Host)
	fmt.Println("Path", result.Path)
	fmt.Println("RawQuery", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Printf("The type of per_page param is: %T\n", qparams["per_page"])
	fmt.Println(qparams["per_page"])

	for _, val := range qparams {
		fmt.Println("Params is:", val)
	}

	parseOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "api.github.com",
		Path:     "/users",
		RawQuery: "since=50&per_page=10",
	}

	anotherUrl := parseOfUrl.String()
	fmt.Println("Another url is:", anotherUrl)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
