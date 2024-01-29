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
	router.Handler(http.MethodPost, "/link/delete/:shortLink", dynamic.ThenFunc(app.shortLinkDelete))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	// will make sure all the bellow endpoints require user to be authenticated
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/link/show/links", protected.ThenFunc(app.shortLinkView))
	router.Handler(http.MethodPost, "/link/delete/:shortLink", protected.ThenFunc(app.shortLinkDelete))
	router.Handler(http.MethodGet, "/link/create", protected.ThenFunc(app.shortLinkCreate))
	router.Handler(http.MethodPost, "/link/create", protected.ThenFunc(app.shortLinkCreatePost))
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.requestLogging, secureHeaders)

	return standard.Then(router)
}
