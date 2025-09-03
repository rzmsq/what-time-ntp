package main

import (
	"os"
	"what-time-ntp/internal/config"
	"what-time-ntp/internal/ntpclient"
	"what-time-ntp/internal/timeformatter"
)

const (
	errorCon = 1
)

func main() {
	runApp()
}

func runApp() {
	cfg := config.NewConfig()

	ntpClient := ntpclient.NewNTPClient(cfg.NTPServer)
	currentTime, err := ntpClient.GetTime()
	if err != nil {
		_, err = os.Stderr.WriteString(err.Error())
		if err != nil {
			panic(err)
		}
		os.Exit(errorCon)
	}

	formatter := timeformatter.NewFormatter(cfg.TimeFormat)
	println(formatter.Format(currentTime))
}
