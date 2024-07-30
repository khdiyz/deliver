package config

import (
	"deliver/pkg/logger"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Host        string
	Port        int
	Environment string
	Debug       bool

	DBPostgreDriver string
	DBPostgreDsn    string
	DBPostgreURL    string

	JWTSecret                string
	JWTAccessExpirationHours int
	JWTRefreshExpirationDays int

	HashKey string

	MinioEndpoint    string
	MinioAccessKeyId string
	MinioSecretKey   string
	MinioBucketName  string
	MinioUseSSL      bool
}

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Host:        cast.ToString(getOrReturnDefault("HOST", "localhost")),
			Port:        cast.ToInt(getOrReturnDefault("PORT", "4040")),
			Environment: cast.ToString(getOrReturnDefault("ENVIRONMENT", "")),
			Debug:       cast.ToBool(getOrReturnDefault("DEBUG", "")),

			DBPostgreDriver: cast.ToString(getOrReturnDefault("DB_DRIVER", "postgres")),
			DBPostgreDsn:    cast.ToString(getOrReturnDefault("DB_DSN", "")),
			DBPostgreURL:    cast.ToString(getOrReturnDefault("DB_URL", "")),

			JWTSecret:                cast.ToString(getOrReturnDefault("JWT_SECRET", "")),
			JWTAccessExpirationHours: cast.ToInt(getOrReturnDefault("JWT_ACCESS_EXPIRATION_HOURS", "")),
			JWTRefreshExpirationDays: cast.ToInt(getOrReturnDefault("JWT_REFRESH_EXPIRATION_DAYS", "")),

			HashKey: cast.ToString(getOrReturnDefault("HASH_KEY", "")),

			MinioEndpoint:    cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "")),
			MinioAccessKeyId: cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "")),
			MinioSecretKey:   cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY", "")),
			MinioBucketName:  cast.ToString(getOrReturnDefault("MINIO_BUCKET_NAME", "")),
			MinioUseSSL:      cast.ToBool(getOrReturnDefault("MINIO_USE_SLL", false)),
		}
	})

	return instance
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	err := godotenv.Load(".env")
	if err != nil {
		logger.GetLogger().Error(err)
	}
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
