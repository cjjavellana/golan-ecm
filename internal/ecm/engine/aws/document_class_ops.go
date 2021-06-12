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

func (dcOp *DocumentClassOperation) SaveDocumentClass(documentClass ce.DocumentClass) (ce.DocumentClass, error) {
	workspaceId, err := primitive.ObjectIDFromHex(documentClass.GetWorkspaceId())
	if err != nil {
		return documentClass, err
	}

	// ensure workspace exists
	findWorkspaceRes := dcOp.db.document.FindOne(context.TODO(), bson.M{
		"_id": workspaceId,
	})
	if findWorkspaceRes.Err() != nil {
		return documentClass, errors.New("workspace " + documentClass.GetWorkspaceId() + " does not exist")
	}

	// ensure no doc class of the same name exists
	docClassExistRes := dcOp.db.documentClass.FindOne(context.TODO(), bson.M{
		"WorkspaceId": workspaceId,
		"Name":        documentClass.GetName(),
	})
	if docClassExistRes.Err() == nil {
		return documentClass, errors.New("document class " + documentClass.GetName() + " already exist")
	}

	dc, ok := bson.Marshal(documentClass)
	if ok != nil {
		return documentClass, ok
	}

	res, err := dcOp.db.documentClass.InsertOne(context.TODO(), dc)
	if err != nil {
		return documentClass, err
	}

	// set newly inserted id into the document class
	documentClass.(*DocumentClass).ID = res.InsertedID.(primitive.ObjectID)

	return documentClass, nil
}
