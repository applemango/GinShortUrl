package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Url struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func ConnectDB() error {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		return err
	}
	query := `CREATE TABLE IF NOT EXISTS url(
		id INT
		url STRING
	)`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func GetUrl(id int) (Url, error) {
	stmt, err := DB.Prepare("SELECT id, url from url WHERE id = ?")
	if err != nil {
		return Url{}, err
	}
	url := Url{}
	sqlErr := stmt.QueryRow(id).Scan(&url.Id, &url.Url)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Url{}, nil
		}
		return Url{}, sqlErr
	}
	return url, nil
}

func AddUrl(newUrl Url) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO url (id, url) VALUES (?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newUrl.Id, newUrl.Url)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}
