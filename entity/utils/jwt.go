package utils

import "time"

type JwtTokenDetail struct {
	ID        string
	Token     string
	ExpiresAt time.Time
}

//go:generate mockery --name Jwt
type Jwt interface {
	GenerateToken(claims JwtClaim) (tokenDetail JwtTokenDetail, err error)
	ParseToken(tokenString string) (claims interface{}, err error)
}

type JwtClaim interface {
	GenerateID(randomString string)
	GetID() string
}
