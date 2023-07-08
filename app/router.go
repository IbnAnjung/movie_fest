package app

import (
	"net/http"

	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

func LoadGinRouter(
	config Config,
) *gin.Engine {

	if config.App.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	//healt
	router.GET("/", func(c *gin.Context) {
		utils.SuccessResponse(c, http.StatusOK, "i'm OK", nil)
	})

	return router
}
