package common

import (
	_ "fmt"
	"github.com/dgrijalva/jwt-go"
)

type Consumer struct {
	Iss    string `json:"iss"`
	Secret string `json:"secret"`
}

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

//func main() {
//	consumer := Consumer{
//		Iss: "a36c3049b36249a3c9f8891cb127243c",
//		Secret : "e71829c351aa4242c2719cbfbe671c09",
//	}
//	t, err := consumer.CreateToken()
//	fmt.Println(t, err)
//}
