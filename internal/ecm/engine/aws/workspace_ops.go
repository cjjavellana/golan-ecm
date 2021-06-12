package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkspaceOperation struct {
	db *CollectionStore
}

func (w *WorkspaceOperation) GetWorkspaceByObjectId(objectId string) (ce.Workspace, error) {
	id, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return nil, err
	}

	res := w.db.document.FindOne(context.TODO(), bson.M{
		"_id": id,
	})

	if res.Err() != nil {
		return nil, res.Err()
	}

	var workspace Workspace
	err = res.Decode(&workspace)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (w *WorkspaceOperation) NewWorkspace(
	descriptor ce.ObjectDescriptor,
) ce.Workspace {

	// TODO: Check if name already exists

	return &Workspace{
		Type: ce.ObjectTypeWorkspace,
		Object: Object{
			Name:        descriptor.Name,
			Label:       descriptor.Label,
			Description: descriptor.Description,
		},
	}
}

func (w *WorkspaceOperation) SaveWorkspace(workspace ce.Workspace) (ce.Workspace, error) {
	m, ok := bson.Marshal(workspace)
	if ok != nil {
		return workspace, ok
	}

	res, err := w.db.document.InsertOne(context.TODO(), m)
	if err != nil {
		return workspace, err
	}

	// cast the interface to a struct so that we can assign the generated id
	v := workspace.(*Workspace)
	v.ID = res.InsertedID.(primitive.ObjectID)

	log.Infof("workspace %s created: %s", workspace.GetName(), workspace.ObjectId())

	return workspace, nil
}
