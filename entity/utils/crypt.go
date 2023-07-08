package utils

//go:generate mockery --name Crypt
type Crypt interface {
	HashString(str string) (string, error)
	VerifyHash(strHash, plainText string) bool
}
