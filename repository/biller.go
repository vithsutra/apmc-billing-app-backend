package repository

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
	"github.com/vsynclabs/billsoft/pkg/storage"
)

type BillerRepo struct {
	db       *sql.DB
	s3Client *storage.LocalFileStorage
}

func NewBillerRepo(db *sql.DB, s3Client *storage.LocalFileStorage) *BillerRepo {
	return &BillerRepo{
		db:       db,
		s3Client: s3Client,
	}
}
func (b *BillerRepo) CreateBillerWithLogo(r *http.Request) error {
	// Parse multipart form data (files and fields)
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		log.Println("Error parsing multipart form:", err)
		return fmt.Errorf("form parsing error: %w", err)
	}

	// Dump all form values for debugging
	for key, values := range r.MultipartForm.Value {
		log.Printf("Received form key: %s => Value: %v", key, values)
	}

	// Log specifically for 'biller_mobile'
	log.Println("Raw biller_mobile:", r.FormValue("biller_mobile"))

	// Trim all inputs to avoid issues with leading/trailing spaces
	biller := models.Biller{
		BillerName:    strings.TrimSpace(r.FormValue("biller_name")),
		BillerAddress: strings.TrimSpace(r.FormValue("biller_address")),
		BillerMobile:  strings.TrimSpace(r.FormValue("biller_mobile")),
		BillerGstin:   strings.TrimSpace(r.FormValue("biller_gstin")),
		BillerPan:     strings.TrimSpace(r.FormValue("biller_pan")),
		BillerMail:    strings.TrimSpace(r.FormValue("biller_mail")),
		UserId:        strings.TrimSpace(r.FormValue("user_id")),
	}

	// Log received biller object
	log.Printf("Received biller: %+v\n", biller)

	// Validate struct
	validate := validator.New()
	if err := validate.Struct(biller); err != nil {
		log.Println("Validation failed:", err)
		return fmt.Errorf("validation error: %w", err)
	}

	// Handle file (logo) upload
	file, header, err := r.FormFile("logo")
	if err != nil {
		log.Println("Logo not uploaded or error:", err)
		return fmt.Errorf("logo upload error: %w", err)
	}
	defer file.Close()

	// Copy the file to buffer
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return fmt.Errorf("file copy error: %w", err)
	}

	// Generate ID and save file
	biller.BillerId = uuid.NewString()
	fileName := fmt.Sprintf("%s-%s", biller.BillerId, header.Filename)

	// Upload logo to S3 (or your storage service)
	if err := b.s3Client.UploadCompanyLogo(fileName, buf); err != nil {
		return fmt.Errorf("error uploading logo: %w", err)
	}
	biller.BillerCompanyLogo = fileName

	// Store biller data in DB
	query := database.NewQuery(b.db)
	if err := query.CreateBiller(&biller); err != nil {
		log.Println("DB insert error:", err)
		return fmt.Errorf("failed to create biller: %w", err)
	}

	log.Println("✅ Biller with logo created successfully:", biller.BillerId)
	return nil
}

func (b *BillerRepo) DeleteBiller(r *http.Request) error {
	vars := mux.Vars(r)
	billerId := vars["biller_id"]

	if billerId == "" {
		return errors.New("biller_id is required")
	}

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
		return nil, errors.New("user_id cannot be empty")
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
	log.Println("Extracted Vars:", vars)

	userId, ok := vars["userId"]
	if !ok || userId == "" {
		log.Println("Error: userId not found in URL")
		return errors.New("missing userId in URL parameters")
	}

	log.Println("Extracted userId:", userId)

	log.Println("Received Content-Type:", r.Header.Get("Content-Type"))

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Error parsing form:", err)
		return fmt.Errorf("error parsing form: %w", err)
	}

	file, header, err := r.FormFile("logo")
	if err != nil {
		log.Println("Error retrieving file:", err)
		return fmt.Errorf("error retrieving file: %w", err)
	}
	defer file.Close()

	log.Println("Received file:", header.Filename, "Size:", header.Size)

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println("Error copying file:", err)
		return fmt.Errorf("error copying file data: %w", err)
	}

	fileName := fmt.Sprintf("%s-%s", userId, header.Filename)

	if err := repo.s3Client.UploadCompanyLogo(fileName, buf); err != nil {
		log.Println("Error uploading to S3:", err)
		return fmt.Errorf("error uploading to S3: %w", err)
	}

	log.Printf("Successfully uploaded %s to S3", fileName)
	return nil
}

func (repo *BillerRepo) DeleteCompanyLogo(r *http.Request) error {
	if repo.s3Client == nil {
		log.Println("S3 client is not initialized")
		return errors.New("S3 client is not initialized")
	}

	vars := mux.Vars(r)
	fileName, ok := vars["file_name"]
	if !ok || fileName == "" {
		return errors.New("missing file_name in URL parameters")
	}

	if err := repo.s3Client.DeleteCompanyLogo(fileName); err != nil {
		log.Println("Error deleting file from S3:", err)
		return fmt.Errorf("error deleting file from S3: %w", err)
	}

	log.Printf("Successfully deleted %s from S3", fileName)
	return nil
}
