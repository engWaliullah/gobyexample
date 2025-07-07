package main

import (
	"fmt"
	"net/http"
)

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from BD")
}

func aboutHabibaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hab... ")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/me", aboutMeHandler)
	router.HandleFunc("/hab", aboutHabibaHandler)

	// port := 3000

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		fmt.Println("Err:", err)
	}
}
