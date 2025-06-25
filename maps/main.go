package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["ts"] = "Typescript"
	languages["ps"] = "Python"
	languages["rb"] = "Ruby"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	delete(languages, "rb")

	fmt.Println(languages)

	// for key, value := range languages {
	// 	fmt.Printf("For key %v, value is %v\n", key, value)
	// }

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

}
