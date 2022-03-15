package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/11-wide-config/pkg/config"
	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/11-wide-config/pkg/handlers"
	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/11-wide-config/pkg/render"
)

const portNumber string = ":8000"

func main() {

	var app config.AppConfig
	tc, err := render.CreateRenderPage()
	if err != nil {
		fmt.Println(err)
	}
	app.TemplateCache = tc
	fmt.Println("Listing the server http://localhost" + portNumber)
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.HandleFunc("/project", handlers.ProjectPage)
	http.ListenAndServe(portNumber, nil)
}
