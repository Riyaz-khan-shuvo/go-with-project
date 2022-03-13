package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/10-settings-config/pkg/config"
	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/9-working-with-layout/pkg/handlers"
	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/9-working-with-layout/pkg/render"
)

const portNumber string = ":9000"

func main() {

	var app config.AppConfig

	tc, err := render.CerateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	fmt.Println("Listing the port http://localhost" + portNumber)
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.ListenAndServe(portNumber, nil)
}
