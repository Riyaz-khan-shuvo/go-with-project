package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/8-enabling-go-modules/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "homePage.gohtml")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "aboutPage.gohtml")
}
