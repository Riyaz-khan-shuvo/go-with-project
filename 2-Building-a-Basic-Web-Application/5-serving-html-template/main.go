package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber string = ":8000"

// home page

func homePage(w http.ResponseWriter, r *http.Request) {

	renderPage(w, "home.gohtml")
}

// About Page

func aboutPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "about.gohtml")

}

// render all Page

func renderPage(w http.ResponseWriter, temp string) {
	// parseTemp, _ := template.ParseFiles("./templates/" + temp)

	// err := parseTemp.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// or

	parseTemp, err := template.ParseFiles("./templates/" + temp)
	if err != nil {
		fmt.Println(err.Error())
	}
	parseTemp.Execute(w, nil)
}

func main() {

	fmt.Print("listing the port http://localhost", portNumber)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)

	http.ListenAndServe(portNumber, nil)

}
