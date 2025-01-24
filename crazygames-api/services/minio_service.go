package services

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"crazygames.io/config"
	"github.com/minio/minio-go/v7"
)

type MinIOService struct {
	client     *minio.Client
	bucketName string
}

type MinIOServiceInterface interface {
	UploadFile(filePath string, destinationPath string, contentType string) (string, error)
	UploadFromURL(url string, destinationPath string) (string, error)
	GetObjectURL(objectPath string) string
	DownloadFile(objectPath string, destinationPath string) error
	DeleteFile(objectPath string) error
	GeneratePresignedURL(objectPath string, expiry time.Duration) (string, error)
}

func NewMinIOService(client *minio.Client) *MinIOService {
	return &MinIOService{
		client:     client,
		bucketName: config.AppConfig.MinIOBucketName,
	}
}

func (m *MinIOService) UploadFile(filePath string, destinationPath string, contentType string) (string, error) {
	if contentType == "" {
		contentType = mime.TypeByExtension(filepath.Ext(filePath))
	}

	_, err := m.client.FPutObject(context.Background(), m.bucketName, destinationPath, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("error uploading file: %v", err)
	}

	return m.GetObjectURL(destinationPath), nil
}

func (m *MinIOService) UploadFromURL(url string, destinationPath string) (string, error) {
	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error downloading file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create a temp file
	tempFile, err := os.CreateTemp("", "minio-upload-*"+filepath.Ext(url))
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Save downloaded content to temp file
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("error saving downloaded file: %v", err)
	}
	tempFile.Close()

	// Upload to MinIO
	return m.UploadFile(tempFile.Name(), destinationPath, resp.Header.Get("Content-Type"))
}

func (m *MinIOService) GetObjectURL(objectPath string) string {
	return fmt.Sprintf("%s/%s/%s", strings.TrimSuffix(config.AppConfig.MinIOEndpoint, "/"), m.bucketName, strings.TrimPrefix(objectPath, "/"))
}

func (m *MinIOService) DownloadFile(objectPath string, destinationPath string) error {
	err := m.client.FGetObject(context.Background(), m.bucketName, objectPath, destinationPath, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("error downloading file: %v", err)
	}
	return nil
}

func (m *MinIOService) DeleteFile(objectPath string) error {
	err := m.client.RemoveObject(context.Background(), m.bucketName, objectPath, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("error deleting file: %v", err)
	}
	return nil
}

func (m *MinIOService) GeneratePresignedURL(objectPath string, expiry time.Duration) (string, error) {
	url, err := m.client.PresignedGetObject(context.Background(), m.bucketName, objectPath, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("error generating presigned URL: %v", err)
	}
	return url.String(), nil
}
