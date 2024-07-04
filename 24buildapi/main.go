package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	Id     string  `json:"Id"`
	Name   string  `json:"Name"`
	Price  string  `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// create a Fake DB
var courses []Course

// middlewares / helpers - file
func (c *Course) isEmpty() bool {
	// return c.Id == "" && c.Name == ""
	return c.Name == ""
}

func main() {
	fmt.Println("Welcome to build API in golang")

	r := mux.NewRouter()

	// seeding
	courses = append(
		courses,
		Course{Id: "1", Name: "ReactJs", Price: "299",
			Author: &Author{Fullname: "Shahan Ahmed Khan", Website: "lco.dev"},
		},
		Course{Id: "2", Name: "MERN Stack", Price: "199",
			Author: &Author{Fullname: "Shahan Ahmed Khan", Website: "lco.dev"},
		},
	)

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by learncodeonline.in</h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return the response
	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	var newCourse Course
	body := json.NewDecoder(r.Body).Decode(&newCourse)
	fmt.Println("Here is the body", body)

	if newCourse.isEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// checking duplication
	for _, course := range courses {
		if course.Name == newCourse.Name {
			json.NewEncoder(w).Encode("the Course is already present with this name!")
			return
		}
	}

	// generate a unique ID, and convert that into a string
	rand.New(rand.NewSource(time.Now().UnixNano()))
	newCourse.Id = strconv.Itoa(rand.Intn(100))

	// append new course into courses
	courses = append(courses, newCourse)

	json.NewEncoder(w).Encode(newCourse)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// send a response when ID is not found

	// loop through the IDs, remove, add with my ID
	for idx, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)

			var newCourseValue Course
			json.NewDecoder(r.Body).Decode(&newCourseValue)
			newCourseValue.Id = params["id"]
			courses = append(courses, newCourseValue)

			json.NewEncoder(w).Encode(newCourseValue)
			return
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// send a response when ID is not found

	// loop through the IDs, remove
	for idx, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)

			json.NewEncoder(w).Encode("The course has been deleted")
			break
		}
	}
}
