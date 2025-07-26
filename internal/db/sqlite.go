package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"rolerocket/internal/logger"
	"strings"

	// ? Is used for opening db file with sqlite. So it looks like this import isn't used but it is.
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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
	// Enable foreign key constraints
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		fmt.Println("Failed to enable foreign keys!", slog.Any("Error", err))
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	DBInstance = &DB{Conn: db}
	DBInstance.Migrate()
	return db
}

func (db *DB) Migrate() {
	_, err := db.Conn.Exec(`CREATE TABLE IF NOT EXISTS logs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
		correlationId TEXT,
		level TEXT,
        log TEXT,
		info TEXT
    )`)
	if err != nil {
		fmt.Println("Error creating logs table", slog.Any("Error", err))
		log.Fatal(err)
	}

	_, err = db.Conn.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
        username TEXT UNIQUE,
		password TEXT
    )`)
	if err != nil {
		fmt.Println("Error creating users table", slog.Any("Error", err))
		log.Fatal(err)
	}

	_, err = db.Conn.Exec(`CREATE TABLE IF NOT EXISTS roles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
		name TEXT UNIQUE
    )`)
	if err != nil {
		fmt.Println("Error creating roles table", slog.Any("Error", err))
		log.Fatal(err)
	}

	_, err = db.Conn.Exec(`CREATE TABLE IF NOT EXISTS permissions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
		name TEXT UNIQUE
    )`)
	if err != nil {
		fmt.Println("Error creating permissions table", slog.Any("Error", err))
		log.Fatal(err)
	}

	_, err = db.Conn.Exec(`CREATE TABLE IF NOT EXISTS user_roles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
		user_id_fk INTEGER,
		role_id_fk INTEGER,
		FOREIGN KEY (user_id_fk) REFERENCES users(id),
		FOREIGN KEY (role_id_fk) REFERENCES roles(id)
    )`)
	if err != nil {
		fmt.Println("Error creating user_roles table", slog.Any("Error", err))
		log.Fatal(err)
	}

	_, err = db.Conn.Exec(`CREATE TABLE IF NOT EXISTS role_permissions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
		role_id_fk INTEGER,
		permission_id_fk INTEGER,
		FOREIGN KEY (role_id_fk) REFERENCES roles(id),
		FOREIGN KEY (permission_id_fk) REFERENCES permissions(id)
    )`)
	if err != nil {
		fmt.Println("Error creating role_permissions table", slog.Any("Error", err))
		log.Fatal(err)
	}
}

func (db *DB) GetUsers(ctx context.Context, search string) ([]string, error) {
	var rows *sql.Rows
	var err error

	if search == "" {
		rows, err = db.Conn.Query("SELECT username FROM users")
	} else if strings.Contains(search, "*") {
		searchTerm := strings.ReplaceAll(search, "*", "%")
		rows, err = db.Conn.Query("SELECT username FROM users WHERE username LIKE $1", searchTerm)
	} else {
		rows, err = db.Conn.Query("SELECT username FROM users WHERE username = $1", search)
	}

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

func (db *DB) InsertUser(ctx context.Context, username string, password string) {
	_, err := db.Conn.Exec("INSERT INTO users ('username', 'password') VALUES ($1, $2)", username, password)
	if err != nil {
		logger.Warn(ctx, "Error returned whilst getting users", slog.Any("error", err))
	}
}

func (db *DB) VerifyLogin(ctx context.Context, username *string, password *string) error {
	rows, err := db.Conn.Query("SELECT password FROM users WHERE username = $1", username)
	if err != nil {
		logger.Error(ctx, "Error returned whilst getting users", slog.Any("error", err))
		return err
	}

	var hashedPassword []byte
	for rows.Next() {
		var password string
		if err := rows.Scan(&password); err != nil {
			logger.Error(ctx, "Error scanning password", slog.Any("error", err))
			return err
		}
		hashedPassword = []byte(password)
	}

	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(*password))
}
