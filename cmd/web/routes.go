package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/link/redirect/:shortLink", dynamic.ThenFunc(app.shortLink))
	router.Handler(http.MethodPost, "/link/create", dynamic.ThenFunc(app.shortLinkCreate))
	router.Handler(http.MethodPost, "/link/delete/:shortLink", dynamic.ThenFunc(app.shortLinkDelete))
	router.Handler(http.MethodGet, "/link/show/links", dynamic.ThenFunc(app.shortLinkView))

	standard := alice.New(app.requestLogging, secureHeaders)

	return standard.Then(router)
}
