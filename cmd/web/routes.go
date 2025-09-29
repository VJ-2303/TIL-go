package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/til/create", app.tilCreate)
	mux.HandleFunc("/til/view/", app.tilView)
	mux.HandleFunc("/til/edit/", app.tilEdit)

	return mux
}
