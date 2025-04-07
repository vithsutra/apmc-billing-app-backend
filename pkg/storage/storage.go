package storage

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

type LocalFileStorage struct {
	basePath string
}

// NewAwsS3Repo initializes a new S3 repository
func NewLocalFileStorageRepo(basePath string) *LocalFileStorage {
	return &LocalFileStorage{
		basePath: basePath,
	}
}

// UploadCompanyLogo uploads a file to the local file storage
func (repo *LocalFileStorage) UploadCompanyLogo(fileName string, fileBuffer *bytes.Buffer) error {

	if err := os.MkdirAll(repo.basePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	fullPath := filepath.Join(repo.basePath, fileName)

	file, err := os.Create(fullPath)

	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	if _, err := fileBuffer.WriteTo(file); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// DeleteCompanyLogo deletes a file from the local file storage
func (repo *LocalFileStorage) DeleteCompanyLogo(fileName string) error {

	fullPath := filepath.Join(repo.basePath, fileName)

	if err := os.Remove(fullPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", fileName)
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}
