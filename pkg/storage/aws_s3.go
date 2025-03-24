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

// NewAwsS3Repo initializes a new S3 repository
func NewAwsS3Repo(bucketName, awsRegion string) (*AwsS3Repo, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		log.Println("Error initializing AWS session:", err)
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	s3Client := s3.New(sess)
	return &AwsS3Repo{s3Client: s3Client, bucket: bucketName}, nil
}

// UploadCompanyLogo uploads a file to the S3 bucket
func (repo *AwsS3Repo) UploadCompanyLogo(fileName string, fileBuffer *bytes.Buffer) error {
	if repo.s3Client == nil {
		return fmt.Errorf("S3 client is not initialized")
	}

	// Determine content type
	ext := filepath.Ext(fileName)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Upload file to S3
	_, err := repo.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(repo.bucket),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(fileBuffer.Bytes()),
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		log.Println("Error uploading file to S3:", err)
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}

	log.Printf("Successfully uploaded %s to S3", fileName)
	return nil
}

// DeleteCompanyLogo deletes a file from the S3 bucket
func (repo *AwsS3Repo) DeleteCompanyLogo(fileName string) error {
	if repo.s3Client == nil {
		return fmt.Errorf("S3 client is not initialized")
	}

	// Perform delete operation
	_, err := repo.s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(repo.bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		log.Println("Error deleting file from S3:", err)
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}

	// Confirm deletion by calling WaitUntilObjectNotExists
	err = repo.s3Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(repo.bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		log.Println("Error confirming S3 object deletion:", err)
		return fmt.Errorf("error confirming file deletion: %w", err)
	}

	log.Printf("Successfully deleted %s from S3", fileName)
	return nil
}
