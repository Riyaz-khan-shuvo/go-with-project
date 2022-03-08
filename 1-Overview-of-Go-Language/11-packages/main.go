package main

import (
	"log"
	"udemy/1-Overview-of-Go-Language/11-packages/helpers"
)

func main() {
	log.Println("Hello I am working")

	var myVar helpers.SomeType
	myVar.TypeName = "Riyaz Khan"
	log.Println(myVar.TypeName)
}
