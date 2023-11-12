package helper

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "seed"

func GenerateToken(id uint, email string) (string, Error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", InternalServerError("Failed to generate token")
	}

	return signedToken, nil
}

func VerifyToken(context *gin.Context) (interface{}, error) {
	errResponse := errors.New("Sign In to Proceed")
	headerToken := context.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	verifiedToken, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, errResponse
	}

	return verifiedToken, nil
}
