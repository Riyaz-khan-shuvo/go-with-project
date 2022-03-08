package main

import (
	"log"
	"time"
)

var s = "Seven"

// declaring the struct
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var s2 = "six"
	log.Println("s is : ", s)
	log.Println("s2 is : ", s2)

	saySomething("Riyaz") // the value of s will override because we assign the value of s

	user := User{
		FirstName: "Riyaz",
		LastName:  " Hossain",
		Age:       22,
	}
	log.Println(user.FirstName, user.LastName, "Birth Day : ", user.BirthDate)
}

func saySomething(s string) (string, string) {
	log.Println("S from saySomeThing function is : ", s)

	return s, " world"

}
