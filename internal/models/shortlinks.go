package models

import (
	"database/sql"
	"time"
)

type shortLinks struct {
	ID          int
	shortCode   string
	originalURL string
	Created     time.Time
}

type Storage interface {
}

type ShortLinksModel struct {
	DB *sql.DB
}

func (m *ShortLinksModel) Insert(originalUrl string) (int, error) {
	shortUrl := generateShortLink(originalUrl)

	stmt := "INSERT INTO short_links (short_code, original_url, created ) VALUES (?,?, UTC_TIMESTAMP())"

	result, err := m.DB.Exec(stmt, shortUrl, originalUrl)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
