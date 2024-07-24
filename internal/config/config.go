package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort             string
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	JWTKey                 string
	SMTPUsername           string
	SMTPPassword           string
	SMTPHost               string
	SMTPAddress            string
	CloudinaryCloudName    string
	CloudinaryAPIKey       string
	CloudinaryAPISecret    string
	CloudinaryUploadPreset string
	ConnectionString       string
}

var ENV = initConfig()

func initConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Printf("unable to load .env")
	}

	return Config{
		ServerPort:             getEnv("PORT", "8080"),
		DBPort:                 getEnv("DB_PORT", "5432"),
		DBHost:                 getEnv("DB_HOST", "localhost"),
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("DB_PASSWORD", "mypassword"),
		DBName:                 getEnv("DB_NAME", "my_db_name"),
		JWTKey:                 getEnv("JWT_KEY", "someJWTKey"),
		SMTPUsername:           getEnv("SMTP_USERNAME", "someEmail"),
		SMTPPassword:           getEnv("SMTP_PASSWORD", "somePassword"),
		SMTPHost:               getEnv("SMTP_HOST", "smtp.emailprovider.com"),
		SMTPAddress:            getEnv("SMTP_ADDR", "smtp.gmail.com:587"),
		CloudinaryCloudName:    getEnv("CLOUDINARY_CLOUD_NAME", "myCloudName"),
		CloudinaryAPIKey:       getEnv("CLOUDINARY_API_KEY", "myApiKey"),
		CloudinaryAPISecret:    getEnv("CLOUDINARY_API_SECRET", "myApiSecret"),
		CloudinaryUploadPreset: getEnv("CLOUDINARY_UPLOAD_PRESET", "myUploadPreset"),
		ConnectionString:       getEnv("CONNECTION_STRING", "con_str"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
