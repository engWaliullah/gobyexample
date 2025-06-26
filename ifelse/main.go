package main

import "fmt"

func main() {
	loginCount := 15
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 13 {
		result = "watch out"
	} else {
		result = "Irregular User"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}

	// if  err != nil {

	// }

}
