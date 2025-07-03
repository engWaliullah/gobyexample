package main

import (
	"encoding/json"
	"fmt"
)

type skill struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func main() {

	encodeJson()

}

func encodeJson() {
	skills := []skill{
		{"Rezwan", "Sotto "},
		{"Habiba", "Tmr sobi rihoye"},
	}

	// pakage this data as JSON data

	// finalJson, error := json.Marshal(skills)
	// finalJson, error := json.MarshalIndent(skills, "", "\t")
	finalJson, error := json.MarshalIndent(skills, "~", "\t")
	checkNilErr(error)

	fmt.Printf("%s\n", finalJson)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
