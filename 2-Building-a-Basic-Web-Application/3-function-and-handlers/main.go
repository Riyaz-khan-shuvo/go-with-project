package main

import (
	"fmt"
	"net/http"
)

const portNum string = ":8000"

func main() {
	fmt.Printf("Listining ths port http://localhost%s", portNum)
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.ListenAndServe(portNum, nil)

}

func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello , This is Homepage ")
}

func addValues(x, y int) int {
	return x + y
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)

	fmt.Fprintf(w, (fmt.Sprintf("This is About Page and 2 + 2 is %d", sum)))
}
