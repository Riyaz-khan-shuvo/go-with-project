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

func RenderPages(w http.ResponseWriter, temp string) {

	tc, err := CreateTemplatePage()
	if err != nil {
		fmt.Println(err)
	}

	t, ok := tc[temp]
	if !ok {
		fmt.Println(err)
	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

}

func CreateTemplatePage() (map[string]*template.Template, error) {

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
		matches, err := filepath.Glob("../../template/*.temhtml")
		if err != nil {
			fmt.Println(err)
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../template/*.temhtml")
			if err != nil {
				fmt.Println(err)
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
