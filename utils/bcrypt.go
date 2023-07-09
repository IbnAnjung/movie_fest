package utils

import (
	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}

func NewBycrypt() enUtil.Crypt {
	return Bcrypt{}
}

func (b Bcrypt) HashString(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)

	return string(bytes), err
}

func (b Bcrypt) VerifyHash(hashText, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashText), []byte(plainText))
	return err == nil
}
