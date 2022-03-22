package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8888"

func main() {

	fmt.Println("listing the port http://localhost" + portNumber)

	http.ListenAndServe(portNumber, nil)
}
