package app

import (
	"net/http"

	enAuthentication "github.com/IbnAnjung/movie_fest/entity/authentication"
	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
	enUser "github.com/IbnAnjung/movie_fest/entity/users"
	enUtils "github.com/IbnAnjung/movie_fest/entity/utils"
	"github.com/IbnAnjung/movie_fest/handler"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

func LoadGinRouter(
	config Config,
	stringGenerator enUtils.StringGenerator,
	userTokenRepository enUser.UserTokenRepository,
	movieUC enMovie.MovieUseCase,
	movieGenresUC enMovieGenres.MovieGenresUseCase,
	authenticationUC enAuthentication.AuthenticationUsecase,
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

	//middlewares
	adminMiddleware := BasicAuth(config.Admin.Username, config.Admin.Password)
	userMiddleware := AuthMiddleware(config, stringGenerator, userTokenRepository)

	//handlers
	movieHandler := handler.NewMovieHandler(movieUC)
	movieGenresHandler := handler.NewMovieGenresHandler(movieGenresUC)
	authHandler := handler.NewAuthenticationHandler(authenticationUC)

	adminRoute := router.Group("/admin", adminMiddleware)
	adminRoute.POST("/movie/upload", movieHandler.UplodeNewFile)
	adminRoute.PUT("/movie/meta", movieHandler.UpdateMetaData)
	adminRoute.GET("/movie/most-views", movieHandler.GetMostView)

	adminRoute.GET("/movie-genres/most-views", movieGenresHandler.GetMostView)

	// public
	router.GET("/movie/:id", movieHandler.GetDetailMovie)
	router.GET("/movie", movieHandler.GetListMoviewPagination)

	//auth
	authRoute := router.Group("/auth")
	authRoute.POST("/register", authHandler.RegisterUser)
	authRoute.POST("/login", authHandler.Login)
	authRoute.Use(userMiddleware).POST("/logout", authHandler.Logout)
	return router
}
