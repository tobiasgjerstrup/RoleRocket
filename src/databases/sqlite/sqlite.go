package sqlite

import (
	"database/sql"
)

func main() {
	_, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {

	}
}
