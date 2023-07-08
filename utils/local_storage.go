package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
)

type localStorage struct {
	basePath   string
	host       string
	publicPath string
}

func NewLocalStorage(host, publicPath string) enUtil.Storage {
	return &localStorage{
		basePath:   "public/files",
		host:       host,
		publicPath: publicPath,
	}
}

func (s *localStorage) UploadFiles(file multipart.File, filePath string) (path string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(dir, s.basePath, filePath)
	targetFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		return "", err
	}

	return "", nil
}

func (s *localStorage) GetPublicFullPath(filename string) string {
	return fmt.Sprintf("%s/%s/%s", s.host, s.publicPath, filename)
}
