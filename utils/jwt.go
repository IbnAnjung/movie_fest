package utils

import (
	"fmt"
	"time"

	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	issuer          string
	signKey         string
	expireDuration  int64
	stringGenerator enUtil.StringGenerator
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Claims interface{}
}

func NewJwt(
	issuer string,
	signKey string,
	expireDuration int64,
	stringGenerator enUtil.StringGenerator,
) enUtil.Jwt {
	return Jwt{
		issuer:          issuer,
		signKey:         signKey,
		expireDuration:  expireDuration,
		stringGenerator: stringGenerator,
	}
}

func (j Jwt) GenerateToken(claims enUtil.JwtClaim) (tokenDetail enUtil.JwtTokenDetail, err error) {
	claims.GenerateID(j.stringGenerator.UUID())
	jwtclaims := jwtClaims{}
	jwtclaims.Issuer = j.issuer
	jwtclaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.expireDuration) * time.Minute))
	jwtclaims.IssuedAt = jwt.NewNumericDate(time.Now())
	jwtclaims.NotBefore = jwt.NewNumericDate(time.Now())
	jwtclaims.Claims = claims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtclaims)
	signToken, err := token.SignedString([]byte(j.signKey))

	tokenDetail.ID = claims.GetID()
	tokenDetail.Token = signToken
	tokenDetail.ExpiresAt = jwtclaims.ExpiresAt.Time

	return
}

func (j Jwt) ParseToken(tokenString string) (claims interface{}, err error) {
	jwtClaims := jwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.signKey), nil
	})

	if err != nil {
		return jwtClaims, err
	}

	if token.Valid {
		return jwtClaims.Claims, nil
	}

	return nil, err
}
