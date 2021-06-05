package service

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"cjavellana.me/ecm/golan/internal/ecm/pb"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type ObjectStoreService struct {
	ObjectStore ce.ObjectStore

	// Required for future compatibility see https://github.com/grpc/grpc-go/blob/master/cmd/protoc-gen-go-grpc/README.md
	pb.UnimplementedContentEngineServer
}

func (s *ObjectStoreService) CreateWorkspace(_ context.Context, in *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceResponse, error) {

	log.Infof("received create workspace request: %s", in.WorkspaceName)

	now := time.Now()

	w := s.ObjectStore.NewWorkspace(in.WorkspaceName, in.Description)
	w.SetCreatedBy("UserFromAuthToken")
	w.SetDateCreated(&now)

	w, _ = s.ObjectStore.SaveWorkspace(w)

	return &pb.CreateWorkspaceResponse{
		ObjectId: w.ObjectId(),
	}, nil
}

func (s *ObjectStoreService) GetWorkspace(_ context.Context, in *pb.GetWorkspaceRequest) (*pb.GetWorkspaceResponse, error) {

	log.Infof("received get workspace query: %s", in.Query)

	w, err := s.ObjectStore.GetWorkspaceByObjectId(in.GetQuery())
	if err != nil {
		return nil, err
	}

	return &pb.GetWorkspaceResponse{
		ObjectId: w.ObjectId(),
	}, nil
}
