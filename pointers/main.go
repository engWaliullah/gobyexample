package main

import "fmt"

func main() {

	// var prt *int
	// fmt.Println("Value of the pointer is: ", prt)

	myNumbers := 23

	var prt = &myNumbers

	fmt.Println(*prt)

	*prt = *prt + 2

	fmt.Println("New value is:", myNumbers)

}
