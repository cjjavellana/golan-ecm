package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentClassOperation struct {
	db *CollectionStore
}

func (dcOp *DocumentClassOperation) NewDocumentClass(
	descriptor ce.ObjectDescriptor,
) ce.DocumentClass {
	return &DocumentClass{
		Object: Object{
			Name:        descriptor.Name,
			Label:       descriptor.Label,
			Description: descriptor.Description,
		},
	}
}

func (dcOp *DocumentClassOperation) SaveDocumentClass(dc ce.DocumentClass) (ce.DocumentClass, error) {
	workspaceId, err := primitive.ObjectIDFromHex(dc.GetWorkspaceId())
	if err != nil {
		return dc, err
	}

	// ensure Workspace exists
	findWorkspaceRes := dcOp.db.document.FindOne(context.TODO(), bson.M{
		"_id": workspaceId,
	})
	if findWorkspaceRes.Err() != nil {
		return dc, errors.New("Workspace " + dc.GetWorkspaceId() + " does not exist")
	}

	// ensure no doc class of the same name exists
	docClassExistRes := dcOp.db.documentClass.FindOne(context.TODO(), bson.M{
		"WorkspaceId": workspaceId,
		"Name":        dc.GetName(),
	})
	if docClassExistRes.Err() == nil {
		return dc, errors.New("document class " + dc.GetName() + " already exist")
	}

	res, err := dcOp.db.documentClass.InsertOne(context.TODO(), dc)
	if err != nil {
		return dc, err
	}

	// set newly inserted id into the document class
	dc.(*DocumentClass).ID = res.InsertedID.(primitive.ObjectID)

	return dc, nil
}
