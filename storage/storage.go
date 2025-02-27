package storage

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	apperrors "the-wedding-game-api/errors"
)

type StorageInterface interface {
	UploadFile(reader bytes.Reader, fileName string) (string, error)
}

var GetStorage = getStorage

func getStorage() (StorageInterface, error) {
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, apperrors.NewStorageError(err.Error())
	}

	bucketName := os.Getenv("AWS_BUCKET_NAME")
	folderName := os.Getenv("AWS_FOLDER_NAME")

	return NewS3Storage(sess, region, bucketName, folderName), nil
}
