package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTem(w http.ResponseWriter, tem string) {

	tc, err := CreateRenderPage()
	if err != nil {
		fmt.Println(err)
	}

	t, ok := tc[tem]
	if !ok {
		fmt.Println(ok)
	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

	// pToTem, err := template.ParseFiles("../../template/" + tem)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// pToTem.Execute(w, nil)
	// CreateRenderPage()
}

func CreateRenderPage() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../template/*.gohtml")

	if err != nil {
		fmt.Println(err)
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(err)
		}
		matches, err := filepath.Glob("../../template/*.layouthtml")
		if err != nil {
			fmt.Println(err)
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../template/*layouthtml")
			if err != nil {
				fmt.Println(err)
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
