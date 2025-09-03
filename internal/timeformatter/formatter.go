// Package timeformatter предоставляет функциональность для форматирования времени.
package timeformatter

import (
	"fmt"
	"time"
)

// TimeFormatter интерфейс для форматирования времени в строку.
type TimeFormatter interface {
	// Format форматирует время согласно заданному формату.
	Format(t time.Time) string
}

// defaultFormatter базовая реализация форматтера времени.
type defaultFormatter struct {
	layout string // шаблон форматирования времени
}

// NewFormatter создает форматтер с переданным форматом.
func NewFormatter(layout string) TimeFormatter {
	return &defaultFormatter{
		layout: layout,
	}
}

// Format возвращает отформатированное время с префиксом "Current time:".
func (f *defaultFormatter) Format(t time.Time) string {
	return fmt.Sprintf("Current time: %s", t.Format(f.layout))
}
