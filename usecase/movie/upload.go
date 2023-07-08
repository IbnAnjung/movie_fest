package movie

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"time"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	"github.com/IbnAnjung/movie_fest/utils"
)

func (uc MovieUC) UploadMovie(ctx context.Context, input enMovie.UploadMovieInput) (output enMovie.UploadMovieOutput, err error) {
	if input.FileHeader.Size > 1024*1024*10 {
		e := utils.UnprocessableEntityError
		e.Message = "maximum file size limit"
		return output, e
	}

	ext := filepath.Ext(input.FileHeader.Filename)

	if ok, err := regexp.MatchString("(?i).(MP4|MPEG-2|WMV|MOV|WEBM)", ext); !ok || err != nil {
		e := utils.UnprocessableEntityError
		e.Message = "invalid file extension"
		return output, e
	}

	newFilename := fmt.Sprintf("mv_%s_%d%s", uc.stringGenerator.UUID(), time.Now().Unix(), ext)
	newMovie := enMovie.Movie{
		Filename: newFilename,
	}

	transCtx := uc.unitOfWork.Begin(ctx)

	if err := uc.movieRepository.AddMovie(&transCtx, &newMovie); err != nil {
		return output, err
	}

	if _, err := uc.storageService.UploadFiles(input.File, newFilename); err != nil {
		uc.unitOfWork.Rollback(transCtx)
		return output, err
	}

	uc.unitOfWork.Commit(transCtx)

	output.Movie = newMovie
	output.PublicUrl = uc.storageService.GetPublicFullPath(newMovie.Filename)

	return
}
