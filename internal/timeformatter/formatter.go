package timeformatter

import (
	"fmt"
	"time"
)

type TimeFormatter interface {
	Format(t time.Time) string
}

type defaultFormatter struct {
	layout string
}

func NewDefaultFormatter(t time.Time) TimeFormatter {
	return &defaultFormatter{
		layout: time.RFC3339,
	}
}

func NewCustomFormatter(layout string) TimeFormatter {
	return &defaultFormatter{
		layout: layout,
	}
}

func (f *defaultFormatter) Format(t time.Time) string {
	return fmt.Sprintf("Current time: %s", t.Format(f.layout))
}
