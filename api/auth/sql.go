package auth

import (
	"customer/api/utils"
	"database/sql"
	"errors"
	"log"
)

func loginUser(db *sql.DB, user, pass string) (id int, err error) {
	var passwordHash string
	var salt string
	err = db.QueryRow("SELECT id, password, salt FROM administrators WHERE username = $1", user).Scan(&id, &passwordHash, &salt)
	log.Print(err)

	match, err := utils.CheckPassword(pass, passwordHash, salt)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	if !match {
		return id, errors.New("Password not match")
	}
	return id, nil
}
