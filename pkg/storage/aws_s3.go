package storage

import (
	"bytes"
	"fmt"
	"log"
	"mime"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsS3Repo struct {
	s3Client *s3.S3
	bucket   string
}

func NewAwsS3Repo(bucketName string, awsRegion string) (*AwsS3Repo, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return &AwsS3Repo{
		s3Client: s3.New(sess),
		bucket:   bucketName,
	}, nil
}

func (repo *AwsS3Repo) UploadCompanyLogo(fileName string, fileBuffer *bytes.Buffer) error {

	ext := filepath.Ext(fileName)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := repo.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(repo.bucket),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(fileBuffer.Bytes()),
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}

	log.Printf("Successfully uploaded %s to S3", fileName)
	return nil
}

func (repo *AwsS3Repo) DeleteCompanyLogo(fileName string) error {
	_, err := repo.s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(repo.bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}

	log.Printf("Successfully deleted %s from S3", fileName)
	return nil
}
