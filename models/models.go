package models

import (
	"database/sql"
	"fmt"
	"os"

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

	query := "CREATE TABLE IF NOT EXISTS url ( id integer primary key autoincrement, url string )"

	_, err = db.Exec(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : createDB :%v\n", err)
		return err
	}
	DB = db
	return nil
}

func GetUrl(id int) (Url, error) {
	stmt, err := DB.Prepare("SELECT id, url from url WHERE id = ?")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : getDB :%v\n", err)
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
func GetLastUrl() (Url, error) {
	stmt, err := DB.Query("SELECT id, url from url WHERE id = last_insert_rowid()")

	if err != nil {
		fmt.Fprintf(os.Stderr, "err : getDB :%v\n", err)
		return Url{}, err
	}

	defer stmt.Close()
	url := Url{}

	for stmt.Next() {
		sqlErr := stmt.Scan(&url.Id, &url.Url)
		if sqlErr != nil {
			return Url{}, sqlErr
		}
	}
	return url, nil
}

func AddUrl(newUrl Url) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO url ( url ) VALUES ( ? )")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: addUrl :%v\n", err)
		return false, err
	}

	_, err = stmt.Exec(newUrl.Url)
	if err != nil {
		return false, err
	}

	tx.Commit()
	return true, nil
}
