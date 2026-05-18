package local

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type StorageService struct {
}

func NewStorageService() *StorageService {
	return &StorageService{}
}

func (s *StorageService) Save(
	c *gin.Context,
	fileHeader *multipart.FileHeader,
	directory string,
) (string, error) {

	// validate mime type
	if err := validateMimeType(fileHeader); err != nil {
		return "", err
	}

	// generate filename
	fileName := generateFileName(fileHeader.Filename)

	// physical directory
	physicalDirectory := filepath.Join(
		"storage",
		"uploads",
		directory,
	)

	// create directory if not exists
	if err := os.MkdirAll(physicalDirectory, os.ModePerm); err != nil {
		return "", err
	}

	// physical file path
	physicalPath := filepath.Join(
		physicalDirectory,
		fileName,
	)

	// save file
	if err := c.SaveUploadedFile(fileHeader, physicalPath); err != nil {
		return "", err
	}

	// public path
	publicPath := fmt.Sprintf(
		"/uploads/%s/%s",
		directory,
		fileName,
	)

	return publicPath, nil
}

func (s *StorageService) Delete(path string) error {

	if path == "" {
		return nil
	}

	physicalPath := filepath.Join(
		"storage",
		path,
	)

	// remove file
	if err := os.Remove(physicalPath); err != nil {

		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return nil
}
