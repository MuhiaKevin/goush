package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goush shortenr service Home"))
}

func (app *application) shortLinkCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goush create short link"))
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
