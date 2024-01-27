package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(name, email, password string) error {
	// Create a bcrypt hash of the plain-text password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))

	if err != nil {
		var mySQLError *mysql.MySQLError
		// check if the error is of the type mysqlError
		if errors.As(err, &mySQLError) {
			// check the mysqlError number if it is 1062 and if the error message Contains "users_uc_email"
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				// send MyCustom error as the error
				return ErrDuplicateEmail
			}
		}
		// if error is not of type mySQLError then just return the error
		return err
	}
	return nil
}

func (m UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := "SELECT id, hashed_password FROM users WHERE email = ?"

	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)

	// if email doesnt exists then retunr ErrInvalidCredentials
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			// else return the rr
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	// check whether the hashed password and the passwod provided by the user match
	// if they dont match then send ErrInvalidCredentials error
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			// else return the rr
			return 0, err
		}
	}

	return id, nil
}

func (m UserModel) Exists(id int) (bool, error) {
	return false, nil
}
