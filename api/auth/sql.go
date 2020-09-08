package auth

import (
	"database/sql"
	"log"
)

func loginUser(db *sql.DB, user, pass string) (id int, err error) {
	err = db.QueryRow("SELECT id FROM administrators WHERE username = $1 AND password = $2", user, pass).Scan(&id)
	log.Print(err)
	if err != nil {
		return 0, err
	}
	return id, nil
}
