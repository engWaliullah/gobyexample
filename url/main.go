package main

import (
	"fmt"
	"net/url"
)

func main() {
	const myUrl string = "https://portfolio-server-vert-five.vercel.app/skills"

	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawPath)
	// fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

}
