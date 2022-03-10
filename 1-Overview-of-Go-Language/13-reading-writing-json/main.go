package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// firstName string `json:"first_name,string"`
	// lastName  string `json:"last_name"`
	// hairColor string `json:"hair_color"`
	// hasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `[
		{
			"first_name":"Riyaz",
			"last_name":"Hossain",
			"hair_color":"Black",
			"has_dog":true
		},
		{
			"first_name":"Shahid",
			"last_name":"Khan",
			"hair_color":"Black",
			"has_dog":false
		}
	]`

	var u []Person

	err := json.Unmarshal([]byte(myJson), &u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("unmarshal : %v", u)

}
