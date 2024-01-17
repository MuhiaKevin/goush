package main

import (
	"fmt"
	"goush/internal/validator"
	"net/http"
)

type shortURLCreateForm struct {
	OriginalURL         string `form:"originalURL"`
	validator.Validator `form:"-"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) shortLinkCreate(w http.ResponseWriter, r *http.Request) {
	var form shortURLCreateForm

	err := app.decodePostForm(r, &form)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.formDecoder.Decode(&form, r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id, err := app.shortLinks.Insert(form.OriginalURL)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/link/%d", id), http.StatusSeeOther)
}

func (app *application) shortLinkView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Viewing the shortUrl"))
}

func (app *application) shortLink(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goush get link"))
}

func (app *application) shortLinkEdit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goush edit short link"))
}

func (app *application) shortLinkDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goush delete short link"))
}
