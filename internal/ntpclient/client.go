// Package ntpclient предоставляет функциональность для получения точного времени с NTP серверов.
package ntpclient

import (
	"errors"
	"time"

	"github.com/beevik/ntp"
)

// Переменные ошибок для операций NTP клиента.
var (
	// ErrConnection возникает при неудачном подключении к NTP серверу.
	ErrConnection = errors.New("failed to connect to NTP server\n")

	// ErrInvalidResponse возникает при получении некорректного ответа от сервера.
	ErrInvalidResponse = errors.New("invalid response from NTP server\n")
)

// NTPClient интерфейс для получения времени с NTP сервера.
type NTPClient interface {
	GetTime() (time.Time, error)
}

// ntpClient реализация интерфейса NTPClient.
type ntpClient struct {
	server string // адрес NTP сервера
}

// NewNTPClient создает новый экземпляр NTP клиента.
// Параметр server - адрес NTP сервера для подключения.
func NewNTPClient(server string) NTPClient {
	return &ntpClient{
		server: server,
	}
}

// GetTime получает точное время с настроенного NTP сервера.
// Возвращает скорректированное время или ошибку.
func (c *ntpClient) GetTime() (time.Time, error) {
	response, err := ntp.Query(c.server)
	if err != nil {
		return time.Time{}, ErrConnection
	}

	err = response.Validate()
	if err != nil {
		return time.Time{}, ErrInvalidResponse
	}

	return time.Now().Add(response.ClockOffset), nil
}
