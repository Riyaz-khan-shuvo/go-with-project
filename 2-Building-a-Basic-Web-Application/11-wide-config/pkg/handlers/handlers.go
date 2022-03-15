package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/11-wide-config/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTem(w, "homePage.gohtml")

}
func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTem(w, "about.gohtml")
}
