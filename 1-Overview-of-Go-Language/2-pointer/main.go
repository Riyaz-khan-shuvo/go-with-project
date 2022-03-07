package main

import "log"

func main() {

	var myString string

	myString = "green"

	log.Println("My String is set to : ", myString)

	changeUsingPointer(&myString)
	log.Println("After func call myString is set : ", myString)

}

func changeUsingPointer(s *string) {

	log.Println("s is set to : ", s)
	newValue := "red"
	*s = newValue
}
