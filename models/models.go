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
	//query := "CREATE TABLE IF NOT EXISTS url ( id int, url string )"
	query := "CREATE TABLE IF NOT EXISTS url ( id integer primary key autoincrement, url string )"
	//query := `CREATE TABLE IF NOT EXISTS url(
	//	id INT
	//	url STRING
	//)`
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
	//stmt, err := DB.Query("SELECT id, url from url WHERE id = 1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err : getDB :%v\n", err)
		return Url{}, err
	}

	defer stmt.Close()
	url := Url{}

	//sqlErr := stmt.QueryRow(id).Scan(&url.Id, &url.Url)
	for stmt.Next() {
		fmt.Println("f")
		sqlErr := stmt.Scan(&url.Id, &url.Url)
		//var id int
		//var url string
		//stmt.Scan(&id, &url)
		if sqlErr != nil {
			fmt.Fprintf(os.Stderr, "err2: %v\n", err)
		}
		//fmt.Println(id)
		//fmt.Println(url)
	}
	//println(url.Id)
	//println(url.Url)
	//sqlErr := stmt.Scan(&url.Id, &url.Url)
	//if sqlErr != nil {
	//	if sqlErr == sql.ErrNoRows {
	//		return Url{}, nil
	//	}
	//	return Url{}, sqlErr
	//}
	return url, nil
	//return Url{}, nil
}

func AddUrl(newUrl Url) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO url ( url ) VALUES ( ? )")
	//query := "INSERT INTO url ( url ) VALUES ( ? )"
	//stmt, err := DB.Exec(query, newUrl.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: addUrl :%v\n", err)
		return false, err
	}

	//defer stmt.Close()

	_, err = stmt.Exec(newUrl.Url)
	if err != nil {
		return false, err
	}

	tx.Commit()
	return true, nil
}
