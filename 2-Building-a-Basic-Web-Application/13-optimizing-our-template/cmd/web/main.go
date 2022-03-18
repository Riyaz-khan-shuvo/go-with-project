package main

import (
	"fmt"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/13-optimizing-our-template/pkg/config"
	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/13-optimizing-our-template/pkg/handlers"
	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/13-optimizing-our-template/pkg/renders"
)

const portNumber string = ":5000"

func main() {
	var app config.AppConfig
	tc, err := renders.CreateTemplateRender()
	if err != nil {
		fmt.Println(err)
	}
	app.TemplateCache = tc

	renders.NewTemplates(&app)

	fmt.Println("Listing the Port http://localhost" + portNumber)

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.ListenAndServe(portNumber, nil)

}
