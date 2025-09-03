// Package config предоставляет конфигурацию приложения.
package config

import (
	"os"
	"time"
)

// Config структура для хранения конфигурации приложения.
type Config struct {
	NTPServer  string
	TimeFormat string
}

// NewConfig создает и возвращает конфигурацию приложения
func NewConfig() *Config {
	return &Config{
		NTPServer:  getEnvOrDefault("NTP_SERVER", "0.beevik-ntp.pool.ntp.org"),
		TimeFormat: getEnvOrDefault("TIME_FORMAT", time.RFC3339),
	}
}

// getEnvOrDefault возвращает значение переменной окружения или значение по умолчанию.
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
