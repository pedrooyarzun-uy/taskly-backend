package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId int) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userId,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
			"iat": time.Now().Unix(),
		})
	return t.SignedString([]byte(os.Getenv("JWT_SIGNATURE")))

}

func VerifyJWT(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNATURE")), nil
	})

	if err != nil {
		if err != nil {
			return false
		}
	}

	if !token.Valid {
		return false
	}

	return true
}

func GetSubject(tokenString string) int {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNATURE")), nil
	})

	if err != nil {
		return 0
	}

	iss, err := token.Claims.GetSubject()

	if err != nil {
		return 0
	}

	res, _ := strconv.Atoi(iss)

	return res
}
