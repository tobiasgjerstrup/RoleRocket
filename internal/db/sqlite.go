package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"rolerocket/internal/logger"

	// ? Is used for opening db file with sqlite. So it looks like this import isn't used but it is.
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

var DBInstance *DB

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Println("Error opening DB!", slog.Any("Error", err))
		log.Fatal(err)
	}

	// sqlite lite dies a little if multiple are open at the same time
	db.SetMaxOpenConns(1)
	DBInstance = &DB{Conn: db}
	DBInstance.Migrate()
	return db
}

func (db *DB) Migrate() {
	_, err := db.Conn.Exec(`CREATE TABLE IF NOT EXISTS logs (
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
	_, err = db.Conn.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
        username TEXT,
		password TEXT
    )`)
	if err != nil {
		fmt.Println("Error creating users", slog.Any("Error", err))
		log.Fatal(err)
	}
}

func (db *DB) GetUsers(ctx context.Context) ([]string, error) {
	rows, err := db.Conn.Query("SELECT usernamee FROM users")
	if err != nil {
		logger.Error(ctx, "Error returned whilst getting users", slog.Any("error", err))
		return nil, err
	}
	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			logger.Error(ctx, "Error scanning username", slog.Any("error", err))
			return nil, err
		}
		usernames = append(usernames, username)
	}

	if err := rows.Err(); err != nil {
		logger.Error(ctx, "Rows iteration error", slog.Any("error", err))
		return nil, err
	}

	return usernames, nil
}

func (db *DB) InsertUser(ctx context.Context) {
	_, err := db.Conn.Exec("INSERT INTO users ('username', 'password') VALUES ('username123', 'pass123')")
	if err != nil {
		logger.Warn(ctx, "Error returned whilst getting users", slog.Any("error", err))
	}
}
