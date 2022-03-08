package main

import "log"

func main() {
	// for loop

	for i := 0; i < 6; i++ {
		log.Println(i)
	}

	// for range loop

	mySlice := []string{"Riyaz", "Khan", "Shuvo", "Fahim", "Asraful"}

	for i, value := range mySlice {
		log.Println(i, value)
	}

	myMap := make(map[string]string)

	myMap["dog"] = "Dog"
	myMap["fish"] = "fish"
	myMap["flower"] = "Flower"
	myMap["cat"] = "Cat"

	// if we don't want to use the index then we use _(underscore)

	for _, value := range myMap {
		log.Println(value)
	}

}
