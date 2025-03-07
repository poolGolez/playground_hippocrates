package users

import (
	"time"

	"example.com/gin/loaney/utils"
)

func Login(params LoginParams) *User {
	user, err := findByUsername(params.Username)

	if err != nil {
		panic("Error encountered during login.")
	}

	return user
}

func Register(params RegisterUserParams) (*User, error) {
	user := &User{
		Username:       params.Username,
		Password:       utils.Hash(params.Password),
		DateRegistered: time.Now(),
	}

	user, err := save(user)
	if err != nil {
		panic("Error encountered during user registration")
	}

	return user, nil
}

type LoginParams struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type RegisterUserParams struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
