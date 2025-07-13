package sqlite

import (
	"database/sql"
	"log"
	"log/slog"
	"rolerocket/logger"

	_ "github.com/mattn/go-sqlite3"
)

func Main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		logger.Logger.Error("Error opening DB!", slog.Any("Error", err))
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS logs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        log TEXT
    )`)
	if err != nil {
		logger.Logger.Error("Error creating table!", slog.Any("Error", err))
		log.Fatal(err)
	}

}
