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

type Permission struct {
	ID         int
	Name       string
	CreateTime string
}

type Role struct {
	ID         int
	Name       string
	CreateTime string
}

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

	_, err = db.Exec("PRAGMA journal_mode = WAL")
	if err != nil {
		fmt.Println("Failed to enable WAL mode!", slog.Any("Error", err))
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	DBInstance = &DB{Conn: db}
	DBInstance.Migrate()
	return db
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

func (db *DB) InsertUser(ctx context.Context, username string, password string) error {
	_, err := db.Conn.Exec("INSERT INTO users ('username', 'password') VALUES ($1, $2)", username, password)
	if err != nil {
		logger.Warn(ctx, "Error returned whilst inserting user", slog.Any("error", err))
		return err
	}
	return nil
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

func (db *DB) GetPermissions(ctx context.Context, search string) ([]Permission, error) {
	var rows *sql.Rows
	var err error

	if search == "" {
		rows, err = db.Conn.Query("SELECT * FROM permissions")
	} else if strings.Contains(search, "*") {
		searchTerm := strings.ReplaceAll(search, "*", "%")
		rows, err = db.Conn.Query("SELECT * FROM permissions WHERE name LIKE $1", searchTerm)
	} else {
		rows, err = db.Conn.Query("SELECT * FROM permissions WHERE name = $1", search)
	}

	if err != nil {
		logger.Error(ctx, "Error returned whilst getting permissions", slog.Any("error", err))
		return nil, err
	}
	defer rows.Close()

	var permissions []Permission
	for rows.Next() {
		var p Permission
		if err := rows.Scan(&p.ID, &p.Name, &p.CreateTime); err != nil {
			logger.Error(ctx, "Error scanning permission", slog.Any("error", err))
			return nil, err
		}
		permissions = append(permissions, p)
	}

	if err := rows.Err(); err != nil {
		logger.Error(ctx, "Rows iteration error", slog.Any("error", err))
		return nil, err
	}

	return permissions, nil
}

func (db *DB) InsertPermission(ctx context.Context, name string) error {
	_, err := db.Conn.Exec("INSERT INTO permissions ('name') VALUES ($1)", name)
	if err != nil {
		logger.Warn(ctx, "Error returned whilst inserting permission", slog.Any("error", err))
		return err
	}

	return nil
}

func (db *DB) GetRoles(ctx context.Context, search string) ([]Role, error) {
	var rows *sql.Rows
	var err error

	if search == "" {
		rows, err = db.Conn.Query("SELECT * FROM roles")
	} else if strings.Contains(search, "*") {
		searchTerm := strings.ReplaceAll(search, "*", "%")
		rows, err = db.Conn.Query("SELECT * FROM roles WHERE name LIKE $1", searchTerm)
	} else {
		rows, err = db.Conn.Query("SELECT * FROM roles WHERE name = $1", search)
	}

	if err != nil {
		logger.Error(ctx, "Error returned whilst getting roles", slog.Any("error", err))
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var r Role
		if err := rows.Scan(&r.ID, &r.Name, &r.CreateTime); err != nil {
			logger.Error(ctx, "Error scanning role", slog.Any("error", err))
			return nil, err
		}
		roles = append(roles, r)
	}

	if err := rows.Err(); err != nil {
		logger.Error(ctx, "Rows iteration error", slog.Any("error", err))
		return nil, err
	}

	return roles, nil
}

func (db *DB) InsertRole(ctx context.Context, name string) error {
	_, err := db.Conn.Exec("INSERT INTO roles ('name') VALUES ($1)", name)
	if err != nil {
		logger.Warn(ctx, "Error returned whilst inserting role", slog.Any("error", err))
		return err
	}

	return nil
}
