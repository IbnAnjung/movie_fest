package utils

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct{}

func NewBycrypt() Bcrypt {
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
