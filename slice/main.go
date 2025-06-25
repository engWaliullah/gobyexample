package main

import (
	"fmt"
	"sort"
)

func main() {

	// create a slice
	var fruitsList = []string{"Appel", "Banana"}

	// check the type of a slice
	fmt.Printf("Type of the fruitsList: %T\n", fruitsList)

	// added a element in a exixting slice
	fruitsList = append(fruitsList, "Tomato", "mango")

	fmt.Println(fruitsList)

	// slicing elements in a exixting slice
	fruitsList = append(fruitsList[1:])
	fmt.Println(fruitsList)

	// make another slice
	highScores := make([]int, 4)

	highScores[0] = 343
	highScores[1] = 234
	highScores[2] = 324
	highScores[3] = 543
	// highScores[4] = 111

	highScores = append(highScores, 121, 32, 54)

	fmt.Println(highScores)

	// check is it sorted or not
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slice based on index
	var courses = []string{"react.js", "next.js", "typescript", "php", "java", "golang"}

	fmt.Println(courses)

	// remove by index
	var index int = 2
	courses = append(courses[:index], courses[index+2:]...)

	fmt.Println(courses)

}

/*

// Golang Slice Tutorial 🚀

package main

import (
	"fmt"
	"sort"
)

func main() {

	// ✅ Creating a slice
	var fruitsList = []string{"Appel", "Banana"}
	fmt.Printf("Type of the fruitsList: %T\n", fruitsList)

	// ✅ Appending elements to a slice
	fruitsList = append(fruitsList, "Tomato", "mango")
	fmt.Println("After append:", fruitsList)

	// 🔪 Slicing to remove the first element
	fruitsList = append(fruitsList[1:])
	fmt.Println("After removing first:", fruitsList)

	// 🧠 Using make() to create a slice with default values
	highScores := make([]int, 4)
	highScores[0] = 343
	highScores[1] = 234
	highScores[2] = 324
	highScores[3] = 543

	// ✅ Appending more values
	highScores = append(highScores, 121, 32, 54)
	fmt.Println("High Scores:", highScores)

	// 📊 Sorting the slice
	fmt.Println("Is sorted?", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted High Scores:", highScores)
	fmt.Println("Is sorted now?", sort.IntsAreSorted(highScores))

	// ❌ Removing a value by index
	var courses = []string{"react.js", "next.js", "typescript", "php", "java", "golang"}
	fmt.Println("Courses:", courses)

	// remove two items starting at index 2
	var index int = 2
	courses = append(courses[:index], courses[index+2:]...)
	fmt.Println("After removal:", courses)
}

/*
🧠 What You Learn:
- Slices are flexible views over arrays in Go.
- Use append() to add elements.
- Use slicing (a[:i], a[i+1:]) to remove elements.
- sort.Ints() for sorting numeric slices.
- make() to preallocate slices.

📌 Slices are memory-efficient, easy to use, and idiomatic in Go.

#golang #goprogramming #100DaysOfCode #codingtips #devlife
*/
