package auth

import "database/sql"

func loginUser(db *sql.DB, user, pass string) (id int, err error) {
	err = db.QueryRow("SELECT id FROM administrators WHERE username = '" + user + "' AND password = " + pass).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
