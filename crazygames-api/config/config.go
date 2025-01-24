package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	// MySQL Config
	MySQLUser     string
	MySQLPassword string
	MySQLHost     string
	MySQLPort     string
	MySQLDB       string

	// Redis Config
	RedisHost     string
	RedisPort     string
	RedisPassword string

	// MinIO Config
	MinIOEndpoint   string
	MinIOAccessKey  string
	MinIOSecretKey  string
	MinIOUseSSL     bool
	MinIOBucketName string

	// JWT Secret
	JWTSecret string

	ALLOW_ORIGINS []string

	FRONTEND_URL string
}

type Oauth2Config struct {
	OauthClientID     string
	OauthClientSecret string
	RedirectURL       string
	Scopes            []string
	Endpoint          oauth2.Endpoint
}

type SMTPConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	RedirectUrl string
}

var AppConfig Config
var GoogleOauthConfig *oauth2.Config
var SMTP SMTPConfig

func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading environment variables")
	}

	// Parse values from environment
	AppConfig = Config{
		MySQLUser:     getEnv("MYSQL_USER", "root"),
		MySQLPassword: getEnv("MYSQL_PASSWORD", ""),
		MySQLHost:     getEnv("MYSQL_HOST", "127.0.0.1"),
		MySQLPort:     getEnv("MYSQL_PORT", "3306"),
		MySQLDB:       getEnv("MYSQL_DB", "testdb"),

		RedisHost:     getEnv("REDIS_HOST", "127.0.0.1"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		MinIOEndpoint:   getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinIOAccessKey:  getEnv("MINIO_ACCESS_KEY", ""),
		MinIOSecretKey:  getEnv("MINIO_SECRET_KEY", ""),
		MinIOUseSSL:     getEnvAsBool("MINIO_USE_SSL", false),
		MinIOBucketName: getEnv("MINIO_BUCKET_NAME", "crazygame"),

		JWTSecret:     getEnv("JWT_SECRET", ""),
		ALLOW_ORIGINS: strings.Split(getEnv("ALLOW_ORIGINS", "http://localhost:3000"), ","),
		FRONTEND_URL:  getEnv("FRONTEND_URL", "http://localhost:3000/home"),
	}

	GoogleOauthConfig = &oauth2.Config{
		ClientID:     getEnv("OAUTH2_ClientID", ""),
		ClientSecret: getEnv("OAUTH2_ClientSecret", ""),
		RedirectURL:  getEnv("OAUTH2_RedirectURL", "http://localhost:8080/api/auth/google/callback"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	SMTP = SMTPConfig{
		Host:        getEnv("SMTP_HOST", ""),
		Port:        465,
		Username:    getEnv("SMTP_USER", ""),
		Password:    getEnv("SMTP_PASS", ""),
		RedirectUrl: getEnv("REDIRECT_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return fallback
}
