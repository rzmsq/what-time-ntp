package timeformatter

import (
	"testing"
	"time"
)

func TestDefaultFormatter_Format(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		layout   string
		expected string
	}{
		{"RFC3339", time.RFC3339, "Current time: 2023-12-25T15:30:45Z"},
		{"Custom format", "2006-01-02 15:04:05", "Current time: 2023-12-25 15:30:45"},
		{"Date only", "2006-01-02", "Current time: 2023-12-25"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := NewFormatter(tt.layout)
			result := formatter.Format(testTime)

			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
