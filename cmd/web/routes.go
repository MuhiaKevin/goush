package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/link/redirect/:shortLink", app.shortLink)
	router.HandlerFunc(http.MethodPost, "/link/create", app.shortLinkCreate)
	router.HandlerFunc(http.MethodDelete, "/link/delete/:shortLink", app.shortLinkDelete)
	router.HandlerFunc(http.MethodGet, "/link/show/links", app.shortLinkView)

	return secureHeaders(app.requestLogging(router))
}
