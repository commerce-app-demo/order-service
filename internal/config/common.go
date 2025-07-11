package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EnvFile string
}

type Environment int

const (
	ProductionEnv Environment = iota
	DevelopmentEnv
	StagingEnv
)

var envType = map[Environment]string{
	ProductionEnv:  "production",
	DevelopmentEnv: "development",
	StagingEnv:     "staging",
}

func (cfg *Config) LoadDB() DBConfig {
	return DBConfig{
		Driver:   getEnvOrDefault("DB_DRIVER", "mysql"),
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		Port:     getEnvOrDefault("DB_PORT", "3308"),
		User:     getEnvOrDefault("DB_USER", "order_user"),
		Password: getEnvOrDefault("DB_PASSWORD", "order_pass"),
		DBName:   getEnvOrDefault("DB_NAME", "order_db"),
	}
}

func (cfg *Config) LoadServer() ServerConfig {
	return ServerConfig{
		Hostname:        getEnvOrDefault("SERVER_HOSTNAME", "localhost"),
		Port:            getEnvOrDefault("SERVER_PORT", "3000"),
		UserGrpcPort:    getEnvOrDefault("SERVER_USER_GRPC_PORT", "50052"),
		ProductGrpcPort: getEnvOrDefault("SERVER_PRODUCT_GRPC_PORT", "50051"),
	}
}

func (cfg *Config) LoadEnv(environment Environment) error {
	err := godotenv.Load(getEnvFile(environment))
	if err != nil {
		return err
	}
	return nil
}

func getEnvOrDefault(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvFile(environment Environment) string {
	switch environment {
	case ProductionEnv:
		return ".env.production"
	case DevelopmentEnv:
		return ".env.development"
	case StagingEnv:
		return ".env.staging"
	default:
		return ".env"
	}
}
