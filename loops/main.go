package main

import "fmt"

func main() {
	// days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Println("index is and value is %v\n", day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			rougeValue++
			// break
			continue
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++

	}

lco:
	fmt.Println("Jumping at: ")

}
