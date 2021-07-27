package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (uint, error) {
	token, err := ExtractToken(c)
	if err != nil {
		return 0, err
	}

	jwt_secret := GetEnv("JWT_SECRET", "secret")
	check_token, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})

	if err != nil || !check_token.Valid {
		return 0, err
	}

	payload := check_token.Claims.(*jwt.StandardClaims)

	id, _ := strconv.Atoi(payload.Id)

	return uint(id),nil

}

func ExtractToken(c *gin.Context) (string, error) {
	type AuthorzationHeader struct {
		Authorzation string `header:"Authorization"`
	}

	var authHeader AuthorzationHeader

	err := c.ShouldBindHeader(&authHeader)

	if err != nil {
		return "", err
	}

	token, err := extract(authHeader.Authorzation)
	if err != nil {
		return "", err
	}
	return token, nil

}

func extract(bearerString string) (string, error) {
	if len(bearerString) == 0 {
		return "", fmt.Errorf("No token")
	}
	strArr := strings.Split(bearerString, " ")
	if len(strArr) == 2 {
		return strArr[1], nil
	}
	return "", fmt.Errorf("Invalid token")
}
