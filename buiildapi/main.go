package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course -- file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int8    `json:"price"`
	Author      *Author `json:"author" `
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"name"`
}

// fake DB

var courses []Course

// middleware, helper - file
func (c *Course) IsEmpthy() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome)

}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to new era </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get single course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with provided ID")
	return
}

func createSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Create single course")
	w.Header().Set("Content-Type", "application/json")

	// what if: bosy is empthy
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please fullfill the required data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpthy() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	// append this new course into courses

	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get single course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with myId

	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Remove the old course
			courses = append(courses[:index], courses[index+1:]...)
			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId == params["id"]

			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
}
