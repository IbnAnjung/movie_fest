package app

import (
	"net/http"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	"github.com/IbnAnjung/movie_fest/handler"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

func LoadGinRouter(
	config Config,
	movieUC enMovie.MovieUseCase,
) *gin.Engine {

	if config.App.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	//healt
	router.GET("/", func(c *gin.Context) {
		utils.SuccessResponse(c, http.StatusOK, "i'm OK", nil)
	})

	// static route
	router.Static("/videos", "./public/files")

	//handlers
	moviewHandler := handler.NewMovieHandler(movieUC)

	adminMiddleware := BasicAuth(config.Admin.Username, config.Admin.Password)
	adminRoute := router.Group("/admin", adminMiddleware)
	adminRoute.POST("/upload", moviewHandler.UplodeNewFile)

	return router
}
