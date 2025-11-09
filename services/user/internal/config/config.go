package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	HTTPPort   string
	GRPCPort   string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("USER_DB_HOST", "localhost"),
		DBPort:     getEnv("USER_DB_PORT", "5432"),
		DBUser:     getEnv("USER_DB_USER", "userservice"),
		DBPassword: getEnv("USER_DB_PASSWORD", "userpass123"),
		DBName:     getEnv("USER_DB_NAME", "users_db"),
		HTTPPort:   getEnv("USER_SERVICE_HTTP_PORT", "8081"),
		GRPCPort:   getEnv("USER_SERVICE_GRPC_PORT", "9091"),
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
