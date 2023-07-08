package handler

import (
	"net/http"
	"strconv"

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

func (h movieHandler) GetDetailMovie(c *gin.Context) {
	strId := c.Param("id")
	movieId, err := strconv.Atoi(strId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnprocessableEntity, "invalid params")
		return
	}

	enMoveDetail, err := h.movieUC.GetMovieDetail(c, int64(movieId))
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	if err = h.movieUC.ViewMovie(c, enMoveDetail); err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	response := presenters.DetailMovieResponse{
		ID:           enMoveDetail.ID,
		Title:        enMoveDetail.Title,
		Duration:     enMoveDetail.Duration,
		Artists:      enMoveDetail.Artists,
		Genres:       []presenters.MovieDetailGenreResponse{},
		Description:  enMoveDetail.Description,
		PublicUrl:    enMoveDetail.PublicUrl,
		ViewsCounter: enMoveDetail.ViewsCounter,
		VotesCounter: enMoveDetail.VotesCounter,
	}

	for _, genre := range enMoveDetail.Genres {
		response.Genres = append(response.Genres, presenters.MovieDetailGenreResponse{
			ID:   genre.ID,
			Name: genre.Name,
		})
	}

	utils.SuccessResponse(c, http.StatusOK, "success", response)
}

func (h movieHandler) GetMostView(c *gin.Context) {
	mov, err := h.movieUC.GetMostView(c)
	if err != nil {
		utils.GeneralErrorResponse(c, err)
		return
	}

	response := presenters.MovieMostViewResponse{
		ID:    mov.ID,
		Title: mov.Title,
		Views: mov.ViewsCounter,
	}

	utils.SuccessResponse(c, http.StatusOK, "success", response)
}
