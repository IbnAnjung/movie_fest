package utils

import "mime/multipart"

type Storage interface {
	UploadFiles(file multipart.File, filePath string) (path string, err error)
	GetPublicFullPath(filename string) string
}
