package handlers

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/13-optimizing-our-template/pkg/renders"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	renders.RenderPage(w, "testHome.gohtml")
}
