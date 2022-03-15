package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/13-optimizing-our-template/pkg/handlers"
)

const portNumber string = ":5000"

func main() {
	fmt.Println("Listing the Port http://localhost" + portNumber)

	http.HandleFunc("/", handlers.HomePage)

	http.ListenAndServe(portNumber, nil)

}
