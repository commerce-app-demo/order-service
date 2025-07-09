package config

import "os"

type Config struct {
}

func (cfg *Config) LoadDB() DBConfig {
	return DBConfig{
		Driver:   getEnvOrDefault("DB_DRIVER", "mysql"),
		Host:     getEnvOrDefault("DB_PORT", "localhost"),
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

func getEnvOrDefault(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
