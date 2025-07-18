package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	// ? Is used for opening db file with sqlite. So it looks like this import isn't used but it is.
	_ "github.com/mattn/go-sqlite3"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Println("Error opening DB!", slog.Any("Error", err))
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS logs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		logTime DATETIME DEFAULT CURRENT_TIMESTAMP,
		correlationId TEXT,
		level TEXT,
        log TEXT,
		info TEXT
    )`)
	if err != nil {
		fmt.Println("Error creating logs", slog.Any("Error", err))
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
        username TEXT,
		password TEXT
    )`)
	if err != nil {
		fmt.Println("Error creating users", slog.Any("Error", err))
		log.Fatal(err)
	}

	return db
}
