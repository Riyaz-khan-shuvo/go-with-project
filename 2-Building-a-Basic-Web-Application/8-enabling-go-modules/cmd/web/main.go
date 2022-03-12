package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/8-enabling-go-modules/pkg/handlers"
)

const portNumber string = ":9000"

func main() {

	fmt.Println("listing the port http://localhost" + portNumber)
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.ListenAndServe(portNumber, nil)
}
