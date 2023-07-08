package movie

import "mime/multipart"

type UploadMovieInput struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type UploadMovieOutput struct {
	Movie
	PublicUrl string
}
