package ecm

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type ServiceHandler struct {
}

func (s *ServiceHandler) mustEmbedUnimplementedContentEngineServer() {
}

func (s *ServiceHandler) CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error) {

	log.Infof("received create workspace request: %s", in.WorkspaceName)

	objectId, _ := uuid.NewUUID()

	return &CreateWorkspaceResponse{
		ObjectId: objectId.String(),
	}, nil
}
