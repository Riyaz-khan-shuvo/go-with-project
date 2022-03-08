package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	color         string
	numberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "Henry",
	}
	printInfo(dog)
	gorilla := Gorilla{
		Name:          "Samson",
		color:         "Black",
		numberOfTeeth: 6,
	}
	printInfo(gorilla)

}

func (d Dog) Says() string {
	return "Woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func printInfo(a Animal) {
	log.Println("This Animal says ", a.Says(), "and has", a.NumberOfLegs(), "Legs")
}

func (d Gorilla) Says() string {
	return "Grunt"
}

func (d Gorilla) NumberOfLegs() int {
	return 2
}
