package helper

import "github.com/golang-jwt/jwt/v5"

var secretKey = "lelestacia"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	return signedToken
}
