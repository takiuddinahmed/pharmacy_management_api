package models

import (
	"project/pharmacy_api/src/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) JwtToken() (string, error) {
	jwt_secret := utils.GetEnv("JWT_SECRET", "secret")
	payload := jwt.StandardClaims{
		Id:        strconv.Itoa(int(user.Id)),
		Subject:   user.Email,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(jwt_secret))
}
