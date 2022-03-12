package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, temp string) {
	pToTem, _ := template.ParseFiles("../../templates/" + temp)
	err := pToTem.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
