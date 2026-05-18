package local

import (
	"errors"
	"mime/multipart"
)

var allowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
}

func validateMimeType(fileHeader *multipart.FileHeader) error {

	contentType := fileHeader.Header.Get("Content-Type")

	if !allowedMimeTypes[contentType] {
		return errors.New("invalid file type")
	}

	return nil
}
