package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", sesrveHome).Methods("GET")

	// http.ListenAndServe(":8000", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func greeter() {
	fmt.Println("Hello")
}

func sesrveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to </h1>"))
}

/*

 go mod tidy --> remove all unusages pakages
 go mod verify
 go mod vendor
 go list
 go list all
 go list -m all
 go list -m -versions github.com/gorilla/mux


*/
