package handler

import (
	"errors"
	"fmt"

	enAuth "github.com/IbnAnjung/movie_fest/entity/authentication"
	"github.com/gin-gonic/gin"
)

func getUserJwt(c *gin.Context) (enAuth.UserJwtClaims, error) {
	fmt.Println("get =>", string(enAuth.JwtKey_User))
	claims, ok := c.Get(string(enAuth.JwtKey_User))
	if !ok {
		fmt.Println("claims =>", claims)
		return enAuth.UserJwtClaims{}, errors.New("claims not found")
	}

	return claims.(enAuth.UserJwtClaims), nil
}
