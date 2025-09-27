package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/til/create", app.tilCreate)
	mux.HandleFunc("/til/view/", app.tilView)

	return mux
}
