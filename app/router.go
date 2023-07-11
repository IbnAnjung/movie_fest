package app

import (
	"net/http"

	enAuthentication "github.com/IbnAnjung/movie_fest/entity/authentication"
	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
	enUserWatch "github.com/IbnAnjung/movie_fest/entity/user_watch"
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
	userVoteUC enUserVote.UserVoteUseCase,
	userWatchUC enUserWatch.UserWatchUseCase,
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
	userVoteHandler := handler.NewUserVoteHandler(userVoteUC)
	userWatchHandler := handler.NewUserWatchHandler(userWatchUC)

	adminRoute := router.Group("/admin", adminMiddleware)
	adminMovieRoute := adminRoute.Group("/movie")
	adminMovieRoute.POST("/upload", movieHandler.UplodeNewFile)
	adminMovieRoute.PUT("/meta", movieHandler.UpdateMetaData)

	adminMovieVotesRoute := adminMovieRoute.Group("/votes")
	adminMovieVotesRoute.GET("/", movieHandler.GetVotes)
	adminMovieVotesRoute.GET("/most", movieHandler.GetMostVote)

	adminMovieViewsRoute := adminMovieRoute.Group("/views")
	adminMovieViewsRoute.GET("/", movieHandler.GetViews)
	adminMovieViewsRoute.GET("/most", movieHandler.GetMostView)

	adminGenreRoute := adminRoute.Group("/genres")
	adminGenreViewsRoute := adminGenreRoute.Group("/views")
	adminGenreViewsRoute.GET("/most", movieGenresHandler.GetMostView)

	adminGenreVotesRoute := adminGenreRoute.Group("/votes")
	adminGenreVotesRoute.GET("/most", movieGenresHandler.GetMostVote)

	// movie
	movieRoute := router.Group("/movie")
	movieRoute.GET("/", movieHandler.GetListMoviewPagination)

	// movie - authenticated
	movieAuth := movieRoute.Group("/").Use(userMiddleware)
	movieAuth.POST("/vote", userVoteHandler.Vote)
	movieAuth.POST("/unvote", userVoteHandler.Unvote)
	movieAuth.GET("/voted", userVoteHandler.GetUserVotedMovies)
	movieAuth.POST("/start", userWatchHandler.StartWatch)
	movieAuth.POST("/playback", userWatchHandler.Playback)
	movieAuth.GET("/history", userWatchHandler.Histories)

	movieRoute.GET("/:id", movieHandler.GetDetailMovie)

	//auth
	authRoute := router.Group("/auth")
	authRoute.POST("/register", authHandler.RegisterUser)
	authRoute.POST("/login", authHandler.Login)
	authRoute.Use(userMiddleware).POST("/logout", authHandler.Logout)

	return router
}
