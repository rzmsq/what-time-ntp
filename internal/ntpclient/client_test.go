package ntpclient

import (
	"errors"
	"testing"
	"time"
)

func TestNewNTPClient(t *testing.T) {
	server := "test.ntp.org"
	client := NewNTPClient(server)

	if client == nil {
		t.Error("Expected non-nil client")
	}
}

// Интеграционный тест (требует сетевого подключения)
func TestNTPClient_GetTime_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := NewNTPClient("0.beevik-ntp.pool.ntp.org")
	ntpTime, err := client.GetTime()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	now := time.Now()
	diff := ntpTime.Sub(now)

	// Проверяем, что разница не больше 1 минуты
	if diff < -time.Minute || diff > time.Minute {
		t.Errorf("Time difference too large: %v", diff)
	}
}

func TestNTPClient_GetTime_InvalidServer(t *testing.T) {
	client := NewNTPClient("invalid.server.com")
	_, err := client.GetTime()

	if err == nil {
		t.Error("Expected error for invalid server")
	}

	if !errors.Is(err, ErrConnection) {
		t.Errorf("Expected ErrConnection, got %v", err)
	}
}
