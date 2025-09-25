package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.errorLog.Printf("%s\n", err.Error())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data *templateData) {

	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exists", page)
		app.serverError(w, err)
		return
	}
	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", app.addDefaultData(data, r))
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(status)
	buf.WriteTo(w)
}
