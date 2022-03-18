package handler

import (
	"net/http"

	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/13-optimizing-template/pkg/config"
	"github.com/Riyaz-khan-shuvo/go-with-project/2-Building-a-Basic-Web-Application/13-optimizing-template/pkg/render"
)

// Repo the repositroy used by the handlers

var Repo *Repository

// Repository is repository Type

type Repository struct {
	App *config.AppConfig
}

// NewRepo Creates the new repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New handlers set the repository for the handlers

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderPages(w, "home.gohtml")

}

func (m *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {

	render.RenderPages(w, "about.gohtml")

}
