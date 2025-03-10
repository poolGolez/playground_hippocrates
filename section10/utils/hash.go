package utils

import "golang.org/x/crypto/bcrypt"

func Hash(plainPassword string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	if err != nil {
		panic("Error encountered while hashing password")
	}

	return string(hashedPassword)
}

func CompareHash(hashedPassword string, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}
