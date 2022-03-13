package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/9-working-with-layout/pkg/render"
)

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.gohtml")
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.gohtml")
}
