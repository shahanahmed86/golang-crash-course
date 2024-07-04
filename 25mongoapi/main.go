package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/shahanahmed86/mongo-api/routers"
)

func main() {
	fmt.Println("Welcome to mongo-api section")
	fmt.Println("Server is getting started...")

	r := router.Router()
	fmt.Println("Listening at port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
