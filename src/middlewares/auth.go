package middlewares

import (
	"net/http"
	"project/pharmacy_api/src/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuth(c *gin.Context) {
	token, err := utils.ExtractToken(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	jwt_secret := utils.GetEnv("JWT_SECRET", "secret")
	check_token, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})

	if err != nil || !check_token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
}
