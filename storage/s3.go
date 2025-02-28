package storage

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"strings"
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
		return "", fmt.Errorf("error uploading file to s3: %w", err)
	}

	if os.Getenv("ENV") == "development" {
		return fmt.Sprintf("http://localhost:9445/%s/%s/%s", RemoveLeadingSlash(s.bucketName), s.folderName, fileName), nil
	}

	return "https://" + RemoveLeadingSlash(s.bucketName) + ".s3." + s.region + ".amazonaws.com/" + s.folderName + "/" + fileName, nil
}

func getS3Storage() (StorageInterface, error) {
	region := os.Getenv("AWS_REGION")

	sess, err := getAwsSession()
	if err != nil {
		return nil, apperrors.NewStorageError(err.Error())
	}

	bucketName := os.Getenv("AWS_BUCKET_NAME")
	folderName := os.Getenv("AWS_FOLDER_NAME")

	return NewS3Storage(sess, region, bucketName, folderName), nil
}

func getAwsSession() (*session.Session, error) {
	region := os.Getenv("AWS_REGION")
	awsBucketEndpoint := os.Getenv("AWS_BUCKET_ENDPOINT")

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(awsBucketEndpoint),
	})
	if err != nil {
		return nil, fmt.Errorf("error creating AWS session: %w", err)
	}
	return sess, nil

}

func RemoveLeadingSlash(s string) string {
	if strings.HasPrefix(s, "/") {
		return s[1:]
	}
	return s
}
