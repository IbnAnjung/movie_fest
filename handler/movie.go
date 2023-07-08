package handler

import (
	"net/http"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	"github.com/IbnAnjung/movie_fest/handler/presenters"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/gin-gonic/gin"
)

type movieHandler struct {
	movieUC enMovie.MovieUseCase
}

func NewMovieHandler(movieUC enMovie.MovieUseCase) movieHandler {
	return movieHandler{
		movieUC: movieUC,
	}
}

func (h movieHandler) UplodeNewFile(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(1024 * 1024 * 10); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}
	defer file.Close()

	output, err := h.movieUC.UploadMovie(c, enMovie.UploadMovieInput{
		File:       file,
		FileHeader: fileHeader,
	})

	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "berhasil", presenters.UploadMovie{
		ID:  output.ID,
		Url: output.PublicUrl,
	})
}

func (h movieHandler) UpdateMetaData(c *gin.Context) {
	req := presenters.UpdateMovieMetaDataRequest{}
	c.ShouldBindJSON(&req)

	if err := h.movieUC.UpdateMetaData(c, enMovie.UpdateMetaDataInput{
		Title:       req.Title,
		Description: req.Description,
		Duration:    req.Duration,
		Artists:     req.Artists,
		Genres:      req.Genres,
		MovieID:     req.MovieID,
	}); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "success", nil)
}
