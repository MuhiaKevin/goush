package main

import (
	"encoding/json"
	"net/http"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Goush Server"}`))
}

func (app *application) shortLinkCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Goush Server"}`))
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
