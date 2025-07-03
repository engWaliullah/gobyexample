package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	// getRequse()
	postJsonRequest()

}

func postJsonRequest() {
	const myUrl = "https://portfolio-server-vert-five.vercel.app/skills"

	// fake json fayload

	requestBody := strings.NewReader(`

			{
				    "name": "Firebase",
    				"image": "https://waliullah99.netlify.app/assets/firebase-2zL_EyTg.png"

			}

		`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	checkNilErr(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println(string(content))

}

func getRequse() {
	// const myUrl = "http://localhost:8000/get"
	const myUrl = "https://portfolio-server-vert-five.vercel.app/skills"

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
