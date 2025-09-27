package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	tils, err := app.models.TILs.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{
		TILs: tils,
	}

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) tilCreate(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "create.html", nil)
}

func (app *application) tilCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "My First TIL"
	content := "This was my very first Today I Learned Post!"

	_, err := app.models.TILs.Insert(title, content)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
