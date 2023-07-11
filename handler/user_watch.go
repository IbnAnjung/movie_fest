package handler

import (
	"net/http"
	"time"

	enUserWatch "github.com/IbnAnjung/movie_fest/entity/user_watch"
	"github.com/IbnAnjung/movie_fest/handler/presenters"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

type userWatchHandler struct {
	userWatchUC enUserWatch.UserWatchUseCase
}

func NewUserWatchHandler(userWatchUC enUserWatch.UserWatchUseCase) userWatchHandler {
	return userWatchHandler{
		userWatchUC: userWatchUC,
	}
}

func (h userWatchHandler) StartWatch(c *gin.Context) {
	req := presenters.StartWatchRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	claims, err := getUserJwt(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	_, playbackID, err := h.userWatchUC.StartPlay(c, enUserWatch.StartPlayInput{
		UserID:    claims.UserID,
		MovieID:   req.MovieID,
		StartTime: time.Duration(req.StartTime) * time.Second,
	})
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "success", presenters.StartWatchResponse{
		PlaybackID: playbackID,
	})
}
