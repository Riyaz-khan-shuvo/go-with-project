package main

import (
	"fmt"

	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/11-wide-config/pkg/config"
	"github.com/Riyaz-khan-shuvo/go-with-project/3-improved-routing/1-using-pat-for-route/pkg/handlers"
	"github.com/Riyaz-khan-shuvo/go-with-project/3-improved-routing/1-using-pat-for-route/pkg/render"
)

const portNumber string = ":8888"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplatePage()
	if err != nil {
		fmt.Println(err)
	}
	app.TemplateCache = tc

	fmt.Println("listing the port http://localhost" + portNumber)

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	http.ListenAndServe(portNumber, nil)
}
