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
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJs Bootcamp",
			"price": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev", "js"]
		}
	`)

	var myCourse course

	isValid := json.Valid(jsonDataFromWeb)
	if !isValid {
		panic("Invalid Json")
	}

	fmt.Println("JSON data is valid")

	err := json.Unmarshal(jsonDataFromWeb, &myCourse)
	checkNilErr(err)

	fmt.Printf("%#v\n", myCourse)

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n", k, v, v)
	}
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
