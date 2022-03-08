package main

import (
	"log"
	"sort"
)

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

	// slice area start

	var mySlice []string

	mySlice = append(mySlice, "Riyaz", "Khan", "Shuvo")
	mySlice = append(mySlice, "Riyaz", "Khan", "Shuvo")

	log.Println(mySlice)

	//  to sort the slice element

	mySliceElement := []int{1, 5, 3, 7, 15, 9, 11, 13}

	log.Println("The Slice element is : ", mySliceElement)
	sort.Ints(mySliceElement)
	log.Println("The Slice element after sort: ", mySliceElement)
	log.Println(mySliceElement[5:7]) //this is the range

}
