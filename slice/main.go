package main

import (
	"fmt"
)

func main() {
	/*
		var fruitsList = []string{"Appel", "Banana"}

		fmt.Printf("Type of the fruitsList: %T\n", fruitsList)

		fruitsList = append(fruitsList, "Tomato", "mango")

		fmt.Println(fruitsList)

		fruitsList = append(fruitsList[1:])
		fmt.Println(fruitsList)

		highScores := make([]int, 4)

		highScores[0] = 343
		highScores[1] = 234
		highScores[2] = 324
		highScores[3] = 543
		// highScores[4] = 111

		highScores = append(highScores, 121, 32, 54)

		fmt.Println(highScores)

			fmt.Println(sort.IntsAreSorted(highScores))
			sort.Ints(highScores)

			fmt.Println(highScores)

			fmt.Println(sort.IntsAreSorted(highScores))
	*/

	// how to remove a value from slice based on index

	var courses = []string{"react.js", "next.js", "typescript", "php", "java", "golang"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+2:]...)

	fmt.Println(courses)

}
