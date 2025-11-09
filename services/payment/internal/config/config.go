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
	OrderGRPCAddr string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:        getEnv("PAYMENT_DB_HOST", "localhost"),
		DBPort:        getEnv("PAYMENT_DB_PORT", "5432"),
		DBUser:        getEnv("PAYMENT_DB_USER", "paymentservice"),
		DBPassword:    getEnv("PAYMENT_DB_PASSWORD", "paymentpass123"),
		DBName:        getEnv("PAYMENT_DB_NAME", "payments_db"),
		HTTPPort:      getEnv("PAYMENT_SERVICE_HTTP_PORT", "8083"),
		GRPCPort:      getEnv("PAYMENT_SERVICE_GRPC_PORT", "9093"),
		OrderGRPCAddr: getEnv("ORDER_GRPC_ADDR", "localhost:9092"),
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
