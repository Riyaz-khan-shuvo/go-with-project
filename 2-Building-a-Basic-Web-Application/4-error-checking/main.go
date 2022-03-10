package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNum string = ":8000"

func main() {

	fmt.Printf("Listining ths port http://localhost%s", portNum)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/divide", DividePage)

	http.ListenAndServe(portNum, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am from HomePage")
}

func addValues(x, y int) int {
	return x + y
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, (fmt.Sprintf("I am From About Page and 2 + 2 is : %d", sum)))
}

func DividePage(w http.ResponseWriter, r *http.Request) {
	f, err := divide(100.00, 0.00)
	if err != nil {
		fmt.Fprintf(w, "can't divide by 0")
		return
		// here if we don't use return it will show the both output
	}

	fmt.Fprintf(w, (fmt.Sprintf("%f divide by %f is %f", 100.00, 0.00, f)))

}

func divide(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("can't divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
