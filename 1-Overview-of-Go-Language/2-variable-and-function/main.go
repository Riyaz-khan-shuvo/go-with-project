package main

import "log"

func main() {

	var say string
	var saySomethingElse string

	say = saySomething("Hello")

	log.Println(say)
	saySomethingElse = saySomething("Good Bye")
	log.Println(saySomethingElse)

	var yourName string

	yourName, _ = Name("Riyaz")

	// if we want to ignore the parameter of the function we have to use  _(under score)

	log.Println(yourName)

	log.Println(Name("Shuvo"))

}

func saySomething(s string) string {
	return s
}

// return more than one value

func Name(name string) (string, string) {
	return "hello ", name
}
