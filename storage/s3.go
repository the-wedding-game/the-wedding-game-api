package storage

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	apperrors "the-wedding-game-api/errors"
)

type S3Storage struct {
	svc        *s3.S3
	region     string
	bucketName string
	folderName string
}

func NewS3Storage(session *session.Session, region string, bucketName string, folderName string) *S3Storage {
	svc := s3.New(session)

	return &S3Storage{
		svc:        svc,
		region:     region,
		bucketName: bucketName,
		folderName: folderName,
	}
}

func (s *S3Storage) UploadFile(reader bytes.Reader, fileName string) (string, error) {
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(s.folderName + "/" + fileName),
		Body:   &reader,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", apperrors.NewStorageError(err.Error())
	}

	return "https://" + s.bucketName + ".s3." + s.region + ".amazonaws.com/" + s.folderName + "/" + fileName, nil
}
