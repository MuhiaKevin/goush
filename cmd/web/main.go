package main

import (
	"database/sql"
	"goush/internal/models"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	shortLinks *models.ShortLinksModel
}

const dsn = "web:pass@/goush?parseTime=true"

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:   errorLog,
		infoLog:    infoLog,
		shortLinks: &models.ShortLinksModel{DB: db},
	}

	app.infoLog.Println("Starting server at port 4000")
	app.errorLog.Fatal(http.ListenAndServe(":4000", app.routes()))
}

// connect to dabase
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
