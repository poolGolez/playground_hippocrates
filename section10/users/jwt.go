package users

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

const PRIVATE_KEY = "17ba460bbc4dba4028a8f8e07d9cb73b3a131c1b74957a6a6a2ccb4c62c295713bf8062449d22194fc2485cdc5613f7a4cafea7ee8091cfeb451cf9310008d24c31781341abd3ab66da25973d2b1ec8a3c5876bbac611fcbbc9b9ff939d56a67caffdc92bbdc62fcbce43c1569944606312876c5b3a28bff124272efd13a8873b0f0f42bb201fcb02ece47b6540bad4c9fbfe53261986984ad21232ade84a6ccc8a144a304bd90f92b45357a60a868a4cb5faee4cda7a3fcf7d38a245c271b1899b9e5981ce265e9cf15d30780846ccf3b67fb917c2422164c5c829bdcb74632a2d3e4ef071b9bfb451096e5faeca38dc00ce3bdff8ea9fdf615511b4d9ae4e7"

func GenerateAuthToken(user *User) string {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"x-userId": user.Id,
		},
	)
	jwtString, err := token.SignedString([]byte(PRIVATE_KEY))

	if err != nil {
		panic("Error generating JWT token")
	}
	return jwtString
}

func DecodeJwt(token string) (*CurrentUser, error) {
	decodedJwt, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unable to decode JWT token")
		}

		return []byte(PRIVATE_KEY), nil
	})

	jwtClaims, ok := decodedJwt.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid JWT token")
	}

	return &CurrentUser{
		UserId: int64(jwtClaims["x-userId"].(float64)),
	}, nil
}

type CurrentUser struct {
	UserId int64
}
