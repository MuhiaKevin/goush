package main

import (
	"errors"
	"fmt"
	"goush/internal/models"
	"goush/internal/validator"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type shortURLCreateForm struct {
	OriginalURL         string `form:"originalURL"`
	validator.Validator `form:"-"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = shortURLCreateForm{}

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) shortLinkCreate(w http.ResponseWriter, r *http.Request) {
	// will hold originalURL here
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

	form.CheckField(validator.NotBlank(form.OriginalURL), "originalURL", "This field cannot be blank")
	form.CheckField(validator.IsURL(form.OriginalURL), "originalURL", "This value is not a valid url")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "home.tmpl", data)
		return
	}

	// get shortcode from the database after saving
	_, err = app.shortLinks.Insert(form.OriginalURL)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Succsessfully created a url!")

	http.Redirect(w, r, "/link/show/links", http.StatusSeeOther)
}

func (app *application) shortLink(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	shortCode := params.ByName("shortLink")

	shortLinks, err := app.shortLinks.Get(shortCode)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	http.Redirect(w, r, shortLinks.OriginalURL, http.StatusSeeOther)
}

func (app *application) shortLinkDelete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	shortCode := params.ByName("shortLink")

	err := app.shortLinks.Delete(shortCode)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Succsessfully deleted  short url!")

	http.Redirect(w, r, "/link/show/links", http.StatusSeeOther)
}

func (app *application) shortLinkView(w http.ResponseWriter, r *http.Request) {
	shortLinks, err := app.shortLinks.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.ShortLinks = shortLinks

	app.render(w, http.StatusOK, "view.tmpl", data)
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User signUp webpage")
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User signUp url endpoint")
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User login webpage")
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User login url endpoint")
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User logout url endpoint")
}
