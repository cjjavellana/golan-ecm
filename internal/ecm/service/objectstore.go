package service

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"cjavellana.me/ecm/golan/internal/ecm/pb"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type ObjectStoreService struct {
	ObjectStore ce.ObjectStore

	// Required for future compatibility see https://github.com/grpc/grpc-go/blob/master/cmd/protoc-gen-go-grpc/README.md
	pb.UnimplementedContentEngineServer
}

func (s *ObjectStoreService) CreateWorkspace(_ context.Context, in *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceResponse, error) {

	log.Infof("received create workspace request: %s", in.WorkspaceName)

	objectId, _ := uuid.NewUUID()

	return &pb.CreateWorkspaceResponse{
		ObjectId: objectId.String(),
	}, nil
}
