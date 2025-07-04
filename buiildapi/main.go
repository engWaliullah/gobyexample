package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	return c.CourseId == "" && c.CourseName == ""
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

}
