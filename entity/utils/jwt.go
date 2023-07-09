package utils

//go:generate mockery --name Jwt
type Jwt interface {
	GenerateToken(claims interface{}) (signToken string, err error)
	ParseToken(tokenString string) (claims interface{}, err error)
}
