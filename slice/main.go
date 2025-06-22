package main

import (
	"fmt"
	"sort"
)

func main() {
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

}
