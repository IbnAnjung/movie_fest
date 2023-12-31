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

func (h userWatchHandler) Playback(c *gin.Context) {
	req := presenters.PlaybackRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	// claims, err := getUserJwt(c)
	// if err != nil {
	// 	utils.GeneralErrorResponse(c, err)
	// 	return
	// }

	err := h.userWatchUC.Playback(c, enUserWatch.PlaybackInput{
		PlaybackID: req.PlaybackID,
		EndTime:    time.Duration(req.Endtime) * time.Second,
	})
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "success", nil)
}

func (h userWatchHandler) Histories(c *gin.Context) {
	claims, err := getUserJwt(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	histories, err := h.userWatchUC.History(c, enUserWatch.HistoryInput{
		UserID: claims.UserID,
	})

	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	response := []presenters.UserWathHistoriesRespones{}
	for _, v := range histories {
		response = append(response, presenters.UserWathHistoriesRespones{
			MovieID:     v.MovieID,
			Title:       v.Title,
			Description: v.Description,
			Artists:     v.Artists,
			Duration:    int64(v.Duration.Seconds()),
		})
	}

	utils.SuccessResponse(c, http.StatusOK, "success", response)
}
