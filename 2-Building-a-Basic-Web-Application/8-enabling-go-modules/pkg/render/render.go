package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, temp string) {
	// pToTem, _ := template.ParseFiles("./template/" + temp)
	// err := pToTem.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	//or

	pToTem, err := template.ParseFiles("../../template/" + temp)
	pToTem.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
