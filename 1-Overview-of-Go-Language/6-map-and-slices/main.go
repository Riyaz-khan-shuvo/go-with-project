package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)
	// the first string is the index and the second string is the type the value.
	myMap["dog"] = "kalu"
	log.Println(myMap)

	myName := make(map[string]User)

	me := User{
		FirstName: "Riyaz",
		LastName:  "Hossain",
	}
	myName["name"] = me
	log.Println(myName["name"])

}
