package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodPost, "/link/create", app.shortLinkCreate)
	router.HandlerFunc(http.MethodGet, "/link/:shortlink", app.shortLink)
	// router.HandlerFunc(http.MethodGet, "/link/view/:shortlink", app.shortLinkView)
	router.HandlerFunc(http.MethodPut, "/link/edit", app.shortLinkEdit)
	router.HandlerFunc(http.MethodDelete, "/link/delete/:shortlink", app.shortLinkDelete)

	return secureHeaders(app.requestLogging(router))
}
