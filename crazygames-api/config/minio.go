package config

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func ConnectMinIO() *minio.Client {
	useSSL := AppConfig.MinIOUseSSL

	client, err := minio.New(AppConfig.MinIOEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AppConfig.MinIOAccessKey, AppConfig.MinIOSecretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalf("Failed to connect to MinIO: %v", err)
	}

	log.Println("MinIO connected successfully!")
	return client
}
