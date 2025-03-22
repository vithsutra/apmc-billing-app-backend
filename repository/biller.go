package repository

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
	"github.com/vsynclabs/billsoft/pkg/storage"
)

type BillerRepo struct {
	db       *sql.DB
	s3Client *storage.AwsS3Repo
}

func NewBillerRepo(db *sql.DB, s3Client *storage.AwsS3Repo) *BillerRepo {
	return &BillerRepo{
		db:       db,
		s3Client: s3Client,
	}
}

func (b *BillerRepo) CreateBiller(r *http.Request) error {
	var biller models.Biller

	if err := json.NewDecoder(r.Body).Decode(&biller); err != nil {
		return fmt.Errorf("invalid request payload: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(biller); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	biller.BillerId = uuid.NewString()

	query := database.NewQuery(b.db)
	if err := query.CreateBiller(&biller); err != nil {
		log.Println("Database error:", err)
		return fmt.Errorf("failed to create biller: %w", err)
	}
	return nil
}

func (b *BillerRepo) DeleteBiller(r *http.Request) error {
	vars := mux.Vars(r)
	billerId := vars["biller_id"]

	query := database.NewQuery(b.db)
	if err := query.DeleteBiller(billerId); err != nil {
		log.Println("Database error:", err)
		return fmt.Errorf("failed to delete biller: %w", err)
	}
	return nil
}

func (repo *BillerRepo) GetBiller(r *http.Request) ([]*models.Biller, error) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	if userId == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	query := database.NewQuery(repo.db)
	billers, err := query.GetBiller(userId)
	if err != nil {
		log.Println("Database error:", err)
		return nil, fmt.Errorf("failed to get billers: %w", err)
	}
	return billers, nil
}

func (repo *BillerRepo) UploadCompanyLogo(r *http.Request) error {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return errors.New("missing userId in URL parameters")
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return fmt.Errorf("error parsing form: %w", err)
	}

	file, header, err := r.FormFile("logo")
	if err != nil {
		return fmt.Errorf("error retrieving file: %w", err)
	}
	defer file.Close()

	if header.Size > 5<<20 {
		return errors.New("file size exceeds 5MB limit")
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return fmt.Errorf("error copying file data: %w", err)
	}

	fileName := fmt.Sprintf("%s-%s", userId, header.Filename)

	if err := repo.s3Client.UploadCompanyLogo(fileName, buf); err != nil {
		return fmt.Errorf("error uploading to S3: %w", err)
	}

	log.Printf("Successfully uploaded %s to S3", fileName)
	return nil
}

func (repo *BillerRepo) DeleteCompanyLogo(r *http.Request) error {
	vars := mux.Vars(r)
	fileName, ok := vars["fileName"]
	if !ok {
		return errors.New("missing fileName in URL parameters")
	}

	if err := repo.s3Client.DeleteCompanyLogo(fileName); err != nil {
		return fmt.Errorf("error deleting file from S3: %w", err)
	}

	log.Printf("Successfully deleted %s from S3", fileName)
	return nil
}
