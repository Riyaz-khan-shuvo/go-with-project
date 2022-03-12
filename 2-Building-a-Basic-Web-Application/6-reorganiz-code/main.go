package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8000"

func main() {

	fmt.Println("listing the port http://localhost", portNumber)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	http.ListenAndServe(portNumber, nil)
}
