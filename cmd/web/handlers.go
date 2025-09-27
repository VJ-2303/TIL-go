package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	if r.Method == http.MethodGet {
		app.render(w, r, http.StatusOK, "create.html", nil)
		return
	}
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	if strings.TrimSpace(title) == "" {
		app.render(w, r, http.StatusUnprocessableEntity, "create.html", nil)
		return
	}
	id, err := app.models.TILs.Insert(title, content)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/til/view/%d", id), http.StatusSeeOther)
}

func (app *application) tilView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/til/view/"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	til, err := app.models.TILs.Get(id)
	if err != nil {
		app.serverError(w, err)
	}
	data := &templateData{
		TIL: til,
	}
	app.render(w, r, http.StatusOK, "view.html", data)
}
