package common

import (
	"github.com/dgrijalva/jwt-go"
)

// Consumer defined consumer iss and secret use jwt token
type Consumer struct {
	Iss    string `json:"iss"`
	Secret string `json:"secret"`
}

// CreateToken with iss , secret
func (c *Consumer) CreateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = c.Iss

	tokenString, err := token.SignedString([]byte(c.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
