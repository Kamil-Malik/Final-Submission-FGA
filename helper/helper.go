package helper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) string {
	salt := 8
	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func ComparePassword(loginPassword, localPassword string) error {
	login, local := []byte(loginPassword), []byte(localPassword)
	err := bcrypt.CompareHashAndPassword(local, login)
	return err
}

func GetContentType(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("please signin to continue")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errors.New("header is not present")
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
