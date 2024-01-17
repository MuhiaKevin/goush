package models

import (
	"database/sql"
	"errors"
	"time"
)

type shortLinks struct {
	ID          int
	ShortCode   string
	OriginalURL string
	Created     time.Time
}

type Storage interface {
}

type ShortLinksModel struct {
	DB *sql.DB
}

func (m *ShortLinksModel) Insert(originalUrl string) (string, error) {
	shortUrl := generateShortLink(originalUrl)

	stmt := "INSERT INTO short_links (short_code, original_url, created ) VALUES (?,?, UTC_TIMESTAMP())"

	_, err := m.DB.Exec(stmt, shortUrl, originalUrl)

	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (m *ShortLinksModel) Get(shortCode string) (*shortLinks, error) {
	stmt := `SELECT original_url FROM short_links WHERE short_code = ? `
	row := m.DB.QueryRow(stmt, shortCode)

	s := &shortLinks{}

	err := row.Scan(&s.OriginalURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *ShortLinksModel) Delete(shortCode string) error {
	stmt := `DELETE FROM short_links WHERE short_code = ? `
	result, err := m.DB.Exec(stmt, shortCode)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count <= 0 {
		return ErrNoRecord
	}

	return nil
}
