package app

import (
	"net/http"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
	"github.com/IbnAnjung/movie_fest/handler"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

func LoadGinRouter(
	config Config,
	movieUC enMovie.MovieUseCase,
	movieGenresUC enMovieGenres.MovieGenresUseCase,
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
	movieHandler := handler.NewMovieHandler(movieUC)
	movieGenresHandler := handler.NewMovieGenresHandler(movieGenresUC)

	adminMiddleware := BasicAuth(config.Admin.Username, config.Admin.Password)
	adminRoute := router.Group("/admin", adminMiddleware)
	adminRoute.POST("/movie/upload", movieHandler.UplodeNewFile)
	adminRoute.PUT("/movie/meta", movieHandler.UpdateMetaData)
	adminRoute.GET("/movie/most-views", movieHandler.GetMostView)

	adminRoute.GET("/movie-genres/most-views", movieGenresHandler.GetMostView)

	// public
	router.GET("/movie/:id", movieHandler.GetDetailMovie)

	return router
}
