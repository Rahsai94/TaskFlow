package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtKey = []byte("your_secret_key")

func GenerateToken(userID uint) string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  userID,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
    })
    tokenString, _ := token.SignedString(jwtKey)
    return tokenString
}

func ValidateToken(tokenString string) bool {
    _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    return err == nil
}
