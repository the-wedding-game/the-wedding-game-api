package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"the-wedding-game-api/config"
	"the-wedding-game-api/constants"
	apperrors "the-wedding-game-api/errors"
)

func ValidateUploadImageRequest(c *gin.Context) (*multipart.FileHeader, error) {
	file, err := c.FormFile("image")
	if err != nil {
		if err.Error() == "missing form body" {
			return &multipart.FileHeader{}, apperrors.NewValidationError(constants.ImageIsRequiredError)
		}

		if err.Error() == "http: no such file" {
			return &multipart.FileHeader{}, apperrors.NewValidationError(constants.ImageIsRequiredError)
		}

		return &multipart.FileHeader{}, fmt.Errorf("error getting file from form: %v", err)
	}

	if !isAllowedExtension(file) {
		return &multipart.FileHeader{}, apperrors.NewValidationError(constants.FileMustBeAnImageError)
	}

	if file.Size == 0 {
		return &multipart.FileHeader{}, apperrors.NewValidationError(constants.FileIsEmptyError)
	}

	if file.Size > int64(config.MAX_UPLOAD_SIZE) {
		return &multipart.FileHeader{}, apperrors.NewValidationError(constants.MaxFileSizeError)
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
