package main

import "net/http"

func homePage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "homePage.gohtml")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "aboutPage.gohtml")
}
