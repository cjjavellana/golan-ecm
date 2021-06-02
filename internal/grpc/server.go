package grpc

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/ecm/objectstorefactory"
	"cjavellana.me/ecm/golan/internal/ecm/pb"
	"cjavellana.me/ecm/golan/internal/ecm/service"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func StartServer(appCfg cfg.AppConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", strconv.Itoa(appCfg.GrpcPort)))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc.EnableTracing = true

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	objectStore := objectstorefactory.GetObjectStore(appCfg)
	s := service.ObjectStoreService{
		ObjectStore: objectStore,
	}

	pb.RegisterContentEngineServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
