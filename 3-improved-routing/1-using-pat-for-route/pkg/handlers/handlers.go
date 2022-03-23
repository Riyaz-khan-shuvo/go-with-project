package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/3-improved-routing/1-using-pat-for-route/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderPages(w, "home.gohtml")

}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPages(w, "about.gohtml")
}
