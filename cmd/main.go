package main

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/cli"
	"cjavellana.me/ecm/golan/internal/grpc"
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

func getAppConfig() cfg.AppConfig {
	args := cli.ParseCli()
	return cfg.ParseConfigFromYamlFile(args.ConfigFile)
}

func main() {
	initLogging()

	grpc.StartServer(getAppConfig())
}


