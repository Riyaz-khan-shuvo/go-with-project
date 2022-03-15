package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/11-wide-config/pkg/handlers"
)

const portNumber string = ":8000"

func main() {

	fmt.Println("Listing the server http://localhost" + portNumber)
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.ListenAndServe(portNumber, nil)
}
