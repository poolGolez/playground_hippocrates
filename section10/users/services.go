package users

import "time"

func Register(params RegisterUserParams) (*User, error) {
	user := &User{
		Username:       params.Username,
		Password:       params.Password,
		DateRegistered: time.Now(),
	}

	user, err := save(user)
	if err != nil {
		panic("Error encountered during user registration")
	}

	return user, nil
}

type RegisterUserParams struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
