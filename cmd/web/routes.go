package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/til/create", app.tilCreatePost)
	mux.HandleFunc("/til/create/form", app.tilCreate)

	return mux
}
