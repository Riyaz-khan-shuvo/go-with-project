package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, temp string) {

	tc, err := CerateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[temp]

	if !ok {
		log.Fatal(ok)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

	// pToTem, _ := template.ParseFiles("../../templates/" + temp)
	// err := pToTem.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

// CerateTemplateCache creates template cache as a map

func CerateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		// ts = template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("../../templates/*.layouthtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layouthtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
