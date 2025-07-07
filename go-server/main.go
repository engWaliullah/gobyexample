package main

import (
	"fmt"
	"io"
	"net/http"
)

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("https://portfolio-server-vert-five.vercel.app")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Hello from BD", string(body))
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
