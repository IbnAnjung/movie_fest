package handler

import (
	"errors"

	enAuth "github.com/IbnAnjung/movie_fest/entity/authentication"
	"github.com/gin-gonic/gin"
)

func getUserJwt(c *gin.Context) (enAuth.UserJwtClaims, error) {
	claims, ok := c.Get(string(enAuth.JwtKey_User))
	if !ok {
		return enAuth.UserJwtClaims{}, errors.New("claims not found")
	}

	return claims.(enAuth.UserJwtClaims), nil
}
