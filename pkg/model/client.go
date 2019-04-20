package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Client struct {
	Id         int    `gorm:"primary_key";"AUTO_INCREMENT"`
	ClientName string `json:"client_name"`
	Password   string `json: password`
}

func (client *Client) GenerateJWT(key []byte) map[string]string {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	clientId := client.Id

	claims["clientId"] = clientId
	claims["created"] = time.Now().Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(key)

	jwtMap := map[string]string{"token": tokenString}

	return jwtMap
}
