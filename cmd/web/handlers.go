package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/vj-2303/til-go/internal/data"
	"github.com/vj-2303/til-go/internal/validator"
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
	til := &data.TIL{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
	}

	v := validator.New()

	if data.ValidateTIL(v, til); !v.Valid() {
		data := &templateData{
			TIL:    til,
			Errors: v.Errors,
		}
		app.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	id, err := app.models.TILs.Insert(til.Title, til.Content)
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
		if errors.Is(err, data.ErrTilNotExists) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := &templateData{
		TIL: til,
	}
	app.render(w, r, http.StatusOK, "view.html", data)
}

func (app *application) tilEdit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/til/edit/"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	til, err := app.models.TILs.Get(id)
	if err != nil {
		if errors.Is(err, data.ErrTilNotExists) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, err)
		}
		return
	}
	if r.Method == http.MethodGet {
		data := &templateData{
			TIL: til,
		}
		app.render(w, r, http.StatusOK, "edit.html", data)
		return
	}
	err = r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}
	til.Title = r.PostForm.Get("title")
	til.Content = r.PostForm.Get("content")

	v := validator.New()

	if data.ValidateTIL(v, til); !v.Valid() {
		data := &templateData{
			TIL:    til,
			Errors: v.Errors,
		}
		app.render(w, r, http.StatusUnprocessableEntity, "edit.html", data)
		return
	}
	err = app.models.TILs.Update(id, til.Title, til.Content)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/til/view/%d", id), http.StatusSeeOther)
}
