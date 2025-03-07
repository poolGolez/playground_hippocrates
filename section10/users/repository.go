package users

import (
	"fmt"

	"example.com/gin/loaney/db"
)

func findByUsername(username string) (*User, error) {
	sql := `
	SELECT id, username, password, date_registered
	FROM users
	WHERE username = ?;
	`
	fmt.Println("Username: (before)", username)

	row := db.DB.QueryRow(sql, username)
	fmt.Println("Username:", username)

	var user User
	var dateString string
	err := row.Scan(&user.Id, &user.Username, &user.Password, &dateString)
	if err != nil {
		return nil, nil
	}

	user.DateRegistered = db.ParseTime(dateString)
	return &user, nil
}

func save(user *User) (*User, error) {
	sql := `
	INSERT INTO users(username, password, date_registered)
	VALUES(?, ?, ?);
	`
	result, err := db.DB.Exec(sql, user.Username, user.Password, user.DateRegistered)
	if err != nil {
		return nil, err
	}

	user.Id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return user, nil
}
