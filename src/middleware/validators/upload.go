package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"the-wedding-game-api/config"
	apperrors "the-wedding-game-api/errors"
)

func ValidateUploadImageRequest(c *gin.Context) (*multipart.FileHeader, error) {
	file, err := c.FormFile("image")
	if err != nil {
		if err.Error() == "missing form body" {
			return &multipart.FileHeader{}, apperrors.NewValidationError("image is required")
		}

		if err.Error() == "http: no such file" {
			return &multipart.FileHeader{}, apperrors.NewValidationError("image is required")
		}

		return &multipart.FileHeader{}, fmt.Errorf("error getting file from form: %v", err)
	}

	if !isAllowedExtension(file) {
		return &multipart.FileHeader{}, apperrors.NewValidationError("file must be an image")
	}

	if file.Size == 0 {
		return &multipart.FileHeader{}, apperrors.NewValidationError("file is empty")
	}

	if file.Size > int64(config.MAX_UPLOAD_SIZE) {
		return &multipart.FileHeader{}, apperrors.NewValidationError("maximum file size is " + fmt.Sprint(config.MAX_UPLOAD_SIZE) + " bytes")
	}

	return file, nil
}

func isAllowedExtension(file *multipart.FileHeader) bool {
	AllowedExtensions := []string{".jpg", ".jpeg", ".png"}

	fileExtension := filepath.Ext(file.Filename)
	for _, extension := range AllowedExtensions {
		if extension == fileExtension {
			return true
		}
	}

	return false
}
