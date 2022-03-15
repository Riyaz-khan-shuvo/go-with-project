package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTem(w http.ResponseWriter, tem string) {
	pToTem, err := template.ParseFiles("../../template/" + tem)
	if err != nil {
		fmt.Println(err)
		return
	}
	pToTem.Execute(w, nil)
}
