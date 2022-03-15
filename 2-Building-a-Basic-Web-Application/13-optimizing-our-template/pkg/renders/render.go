package renders

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderPage(w http.ResponseWriter, temp string) {

	pToTem, err := template.ParseFiles("../../template/" + temp)
	if err != nil {
		fmt.Println(err)
	}
	pToTem.Execute(w, nil)

}
