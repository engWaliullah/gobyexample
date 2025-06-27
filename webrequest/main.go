package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://portfolio-server-vert-five.vercel.app/skills"

func main() {

	res, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", res)

	defer res.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)

	content := string(dataBytes)

	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
