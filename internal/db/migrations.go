package sqlite

import (
	"fmt"
	"log"
	"log/slog"
)

func (db *DB) Migrate() {
	// TODO: Use `PRAGMA foreign_key_list('user_roles');` & `SELECT * FROM pragma_table_info('user_roles')` to get the current sqlite schema and see if there are any updates that needs to be applied.
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
