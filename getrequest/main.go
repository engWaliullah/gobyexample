package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	getRequse()

}

func getRequse() {
	// const myUrl = "http://localhost:8000/get"
	const myUrl = "https://portfolio-server-vert-five.vercel.app"

	res, err := http.Get(myUrl)
	checkNilErr(err)

	defer res.Body.Close()

	fmt.Println("Response status code: ", res.StatusCode)
	fmt.Println("Response content length", res.ContentLength)

	var resString strings.Builder

	content, _ := ioutil.ReadAll(res.Body)

	byteCount, _ := resString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(resString.String())

	// fmt.Println(string(content))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
