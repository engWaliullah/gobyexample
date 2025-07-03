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

	// encodeJson()
	decodeJson()

}

func decodeJson() {

	jsonDataFromWeb := []byte(`
	 		{        
					"name": "Rezwan",     
			       "image": "Sotto "  
			}
	`)

	var skills skill

	checkValidate := json.Valid(jsonDataFromWeb)

	if checkValidate {
		fmt.Println("JSON was validated")
		json.Unmarshal(jsonDataFromWeb, &skills)

		fmt.Printf("%#v\n", skills)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some case where you just want to add data ti key value pair
	var newSkill map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &newSkill)
	fmt.Printf("%#v\n", newSkill)

	for k, v := range newSkill {
		fmt.Printf("key is %v and the value is %v and type is: %T\n", k, v, v)
	}

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
