package handler

import (
	"net/http"

	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
	"github.com/IbnAnjung/movie_fest/handler/presenters"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

type userVoteHandler struct {
	userVoteUC enUserVote.UserVoteUseCase
}

func NewUserVoteHandler(userVoteUC enUserVote.UserVoteUseCase) userVoteHandler {
	return userVoteHandler{
		userVoteUC: userVoteUC,
	}
}

func (h userVoteHandler) Vote(c *gin.Context) {
	req := presenters.UserVoteUnvoteRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	claims, err := getUserJwt(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	if err = h.userVoteUC.Vote(c, &enUserVote.UserVote{
		UserID:  claims.UserID,
		MovieID: req.MovieID,
	}); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "success", nil)
}

func (h userVoteHandler) Unvote(c *gin.Context) {
	req := presenters.UserVoteUnvoteRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	claims, err := getUserJwt(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	if err = h.userVoteUC.Unvote(c, &enUserVote.UserVote{
		UserID:  claims.UserID,
		MovieID: req.MovieID,
	}); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "success", nil)
}

func (h userVoteHandler) GetUserVotedMovies(c *gin.Context) {
	claims, err := getUserJwt(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	movies, err := h.userVoteUC.GetVotedMovies(c, claims.UserID)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	response := []presenters.VotedMoviesResponse{}
	for _, m := range movies {
		response = append(response, presenters.VotedMoviesResponse{
			ID:       m.ID,
			Title:    m.Title,
			Duration: int64(m.Duration.Seconds()),
			Artist:   m.Artist,
		})
	}
	utils.SuccessResponse(c, http.StatusOK, "success", response)
}
