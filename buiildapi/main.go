package main

import (
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
