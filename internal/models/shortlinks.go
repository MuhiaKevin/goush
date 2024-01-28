package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type ShortLinks struct {
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

func (m *ShortLinksModel) Insert(originalUrl string, userID int) (string, error) {
	shortUrl := generateShortLink(originalUrl)
	fmt.Println("USER:ID   ", userID)

	stmt := "INSERT INTO short_links (short_code,user_id,original_url, created ) VALUES (?,?,?, UTC_TIMESTAMP())"

	_, err := m.DB.Exec(stmt, shortUrl, userID, originalUrl)

	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (m *ShortLinksModel) Get(shortCode string) (*ShortLinks, error) {
	stmt := `SELECT original_url FROM short_links WHERE short_code = ? `
	row := m.DB.QueryRow(stmt, shortCode)

	s := &ShortLinks{}

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

func (m *ShortLinksModel) Latest() ([]*ShortLinks, error) {
	stmt := `SELECT id, short_code, original_url, created FROM short_links ORDER BY id DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	shortLinks := []*ShortLinks{}

	for rows.Next() {
		s := &ShortLinks{}

		err = rows.Scan(&s.ID, &s.ShortCode, &s.OriginalURL, &s.Created)
		if err != nil {
			return nil, err
		}

		shortLinks = append(shortLinks, s)

	}

	return shortLinks, err
}
