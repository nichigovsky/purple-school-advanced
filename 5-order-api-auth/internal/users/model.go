package users

import (
	"math/rand"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone string `json:"phone" gorm:"uniqueIndex"`
	SessionId string `json:"sessionId" gorm:"uniqueIndex"`
	Code string `json:"code"`
}

type UserRegistry struct {
	SessionId string `json:"sessionId"`
}

// То, что будем создавать, чтобы потом положить это в БД
func NewUser(phone string) *User {
	user := &User{
		Phone: phone,
	}

	user.GenerateSessionId()
	user.GenerateCode()

	return user
}

func (user *User) GenerateSessionId() {
	user.SessionId = RandRunes(letterRunes, 10)
}

func (user *User) GenerateCode() {
	user.Code = RandRunes(digitRunes, 4)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digitRunes = []rune("0123456789")

func RandRunes(runes []rune, n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	return string(b)
}