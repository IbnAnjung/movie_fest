package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	enAuth "github.com/IbnAnjung/movie_fest/entity/authentication"
	enUser "github.com/IbnAnjung/movie_fest/entity/users"
	enUtils "github.com/IbnAnjung/movie_fest/entity/utils"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

func BasicAuth(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authUsername, authPass, ok := c.Request.BasicAuth()
		if !ok || authUsername != username || authPass != password {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
		}

		c.Next()

	}
}

func AuthMiddleware(conf Config, stringGenerator enUtils.StringGenerator, userRepository enUser.UserTokenRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := utils.NewJwt(conf.App.Name, conf.Jwt.SecretKey, conf.Jwt.ExpireDuration, stringGenerator)

		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "missing authorization")
			c.Abort()
			return
		}

		authPart := strings.Split(auth, " ")
		bearer := strings.ToLower(authPart[0])

		if !strings.EqualFold(bearer, "bearer") || len(authPart) != 2 {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid authorization format")
			c.Abort()
			return
		}

		token := authPart[1]
		claim, err := jwt.ParseToken(token)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		claimMap, ok := claim.(map[string]interface{})
		if !ok {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}
		var userJwtClaims enAuth.UserJwtClaims
		jsonClaims, _ := json.Marshal(claimMap)

		json.Unmarshal(jsonClaims, &userJwtClaims)

		ctx := context.Background()
		tokenDetail := enUser.UserTokenDetail{ID: userJwtClaims.ID}

		if err := userRepository.GetToken(&ctx, &enUser.UserToken{Token: tokenDetail}); err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "expired token")
			c.Abort()
			return
		}

		fmt.Println("set =>", string(enAuth.JwtKey_User))
		c.Set(string(enAuth.JwtKey_User), userJwtClaims)

		c.Next()
	}
}
