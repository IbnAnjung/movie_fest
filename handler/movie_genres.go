package handler

import (
	"net/http"

	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
	"github.com/IbnAnjung/movie_fest/handler/presenters"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

type movieGenresHandler struct {
	movieGenresUC enMovieGenres.MovieGenresUseCase
}

func NewMovieGenresHandler(movieGenresUC enMovieGenres.MovieGenresUseCase) movieGenresHandler {
	return movieGenresHandler{
		movieGenresUC: movieGenresUC,
	}
}

func (h movieGenresHandler) GetMostView(c *gin.Context) {
	mg, err := h.movieGenresUC.GetMostView(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	response := presenters.MovieGenresMostViewResponse{
		ID:    mg.ID,
		Name:  mg.Name,
		Views: mg.ViewsCounter,
	}

	utils.SuccessResponse(c, http.StatusOK, "success", response)
}

func (h movieGenresHandler) GetMostVote(c *gin.Context) {
	mg, err := h.movieGenresUC.GetMostView(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	response := presenters.MovieGenresMostVoteResponse{
		ID:    mg.ID,
		Name:  mg.Name,
		Votes: mg.VotesCounter,
	}

	utils.SuccessResponse(c, http.StatusOK, "success", response)
}
