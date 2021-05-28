package golan

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func StartServer(appCfg cfg.AppConfig) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
