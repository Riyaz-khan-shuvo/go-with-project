package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/13-optimizing-our-template/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderPage(w http.ResponseWriter, temp string) {
	tc := app.TemplateCache

	// tc, err := CreateTemplateRender()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	t, ok := tc[temp]
	if !ok {
		fmt.Println(ok)
	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

	// pToTem, err := template.ParseFiles("../../template/" + temp)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// pToTem.Execute(w, nil)

}

func CreateTemplateRender() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../template/*.gohtml")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(err)
			return myCache, err
		}
		matches, err := filepath.Glob("../../template/*.layouthtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../template/*.layouthtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}

	return myCache, nil
}
