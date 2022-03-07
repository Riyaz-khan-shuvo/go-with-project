package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) PrintFirstName() string {
	return m.FirstName
}

func main() {

	var myVar myStruct
	myVar.FirstName = "Riyaz Khan Shuvo"
	log.Println(myVar.FirstName)
	myVar2 := myStruct{
		FirstName: "Riyaz",
	}
	log.Println(myVar2.PrintFirstName())

}
