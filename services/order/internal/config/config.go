package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	HTTPPort      string
	GRPCPort      string
	UserGRPCAddr  string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:       getEnv("ORDER_DB_HOST", "localhost"),
		DBPort:       getEnv("ORDER_DB_PORT", "5432"),
		DBUser:       getEnv("ORDER_DB_USER", "orderservice"),
		DBPassword:   getEnv("ORDER_DB_PASSWORD", "orderpass123"),
		DBName:       getEnv("ORDER_DB_NAME", "orders_db"),
		HTTPPort:     getEnv("ORDER_SERVICE_HTTP_PORT", "8082"),
		GRPCPort:     getEnv("ORDER_SERVICE_GRPC_PORT", "9092"),
		UserGRPCAddr: getEnv("USER_GRPC_ADDR", "localhost:9091"),
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
