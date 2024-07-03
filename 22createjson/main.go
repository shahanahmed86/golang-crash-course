package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    float32  `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to creating json in golang")
	EncodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	checkNilErr(err)

	var responseString strings.Builder
	responseString.Write(finalJson)

	fmt.Println("Raw", finalJson)
	fmt.Println("convereted", responseString.String())
}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Server is about to panic")
		panic(err)
	}
}
