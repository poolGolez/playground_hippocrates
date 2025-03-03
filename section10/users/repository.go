package users

import "example.com/gin/loaney/db"

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
