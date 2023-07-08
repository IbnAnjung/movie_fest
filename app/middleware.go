package app

import (
	"net/http"

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
