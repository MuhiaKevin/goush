package main

import (
	"log"
	"net/http"
)

type application struct {
}

func main() {
	app := &application{}
	log.Println("Starting server at port 4000")
	log.Fatal(http.ListenAndServe(":4000", app.routes()))
}
