package config

import (
	"os"
	"testing"
	"time"
)

func TestNewConfig_DefaultValues(t *testing.T) {
	// Очищаем переменные окружения
	err := os.Unsetenv("NTP_SERVER")
	if err != nil {
		panic(err)
	}
	err = os.Unsetenv("TIME_FORMAT")
	if err != nil {
		panic(err)
	}

	cfg := NewConfig()

	if cfg.NTPServer != "0.beevik-ntp.pool.ntp.org" {
		t.Errorf("Expected default NTP server, got %s", cfg.NTPServer)
	}

	if cfg.TimeFormat != time.RFC3339 {
		t.Errorf("Expected RFC3339 format, got %s", cfg.TimeFormat)
	}
}

func TestNewConfig_EnvironmentVariables(t *testing.T) {
	err := os.Setenv("NTP_SERVER", "test.ntp.org")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("TIME_FORMAT", "2006-01-02")
	if err != nil {
		panic(err)
	}

	cfg := NewConfig()

	if cfg.NTPServer != "test.ntp.org" {
		t.Errorf("Expected test.ntp.org, got %s", cfg.NTPServer)
	}

	if cfg.TimeFormat != "2006-01-02" {
		t.Errorf("Expected 2006-01-02, got %s", cfg.TimeFormat)
	}

	// Очистка
	err = os.Unsetenv("NTP_SERVER")
	if err != nil {
		panic(err)
	}
	err = os.Unsetenv("TIME_FORMAT")
	if err != nil {
		panic(err)
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		envValue     string
		expected     string
	}{
		{"default value", "TEST_KEY", "default", "", "default"},
		{"env value", "TEST_KEY", "default", "custom", "custom"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				err := os.Setenv(tt.key, tt.envValue)
				if err != nil {
					panic(err)
				}
				defer func(key string) {
					err = os.Unsetenv(key)
					if err != nil {
						panic(err)
					}
				}(tt.key)
				if err != nil {
					panic(err)
				}
			}

			result := getEnvOrDefault(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
