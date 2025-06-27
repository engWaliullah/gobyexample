package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is important for learning files system"

	file, err := os.Create("./myGolangDoc.txt")

	// if err != nil {
	// 	// return nil, err
	// 	panic(err)
	// }

	chechNilErr(err)

	length, err := io.WriteString(file, content)
	chechNilErr(err)

	fmt.Println("Length is:", length)
	file.Close()
	readFile("./myGolangDoc.txt")
}

func readFile(fileName string) {
	// _, err := ioutil.ReadFile(fileName)
	dataBytes, err := ioutil.ReadFile(fileName)
	chechNilErr(err)

	fmt.Println("text data inside the file is: \n", string(dataBytes))
}

func chechNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
