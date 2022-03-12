package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderPage(w http.ResponseWriter, temp string) {
	// pToTem, _ := template.ParseFiles("./template/" + temp)
	// err := pToTem.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	//or

	pToTem, err := template.ParseFiles("./template/" + temp)
	pToTem.Execute(w, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
