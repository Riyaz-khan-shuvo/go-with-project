package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Hi , I am working !!!")
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8000", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello , This is Homepage ")
}
