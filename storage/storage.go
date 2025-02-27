package storage

import (
	"bytes"
)

type StorageInterface interface {
	UploadFile(reader bytes.Reader, fileName string) (string, error)
}

var GetStorage = getS3Storage
