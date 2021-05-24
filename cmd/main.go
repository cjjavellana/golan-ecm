package main

import (
	log "github.com/sirupsen/logrus"
)

func initLogging() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func main() {
	initLogging()
	log.Info("Hello World")
}
