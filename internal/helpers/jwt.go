package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId int, name string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":  strconv.Itoa(userId),
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
			"iat":  time.Now().Unix(),
			"name": name,
		})
	return t.SignedString([]byte(os.Getenv("JWT_SIGNATURE")))

}

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNATURE")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil

}
