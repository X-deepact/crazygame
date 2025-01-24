package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"golang.org/x/net/context"
)

func UploadFileToMinio(client *minio.Client, file *multipart.FileHeader, bucketName string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))

	_, err = client.PutObject(context.Background(), bucketName, filename, src, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("http://%s/%s/%s", os.Getenv("MINIO_ENDPOINT"), bucketName, filename)
	return fileURL, nil
}
