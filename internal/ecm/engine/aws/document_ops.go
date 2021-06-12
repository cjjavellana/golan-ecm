package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentOperation struct {
	db *CollectionStore
}

func (d *DocumentOperation) NewDocument(
	descriptor ce.ObjectDescriptor,
	documentClassId string,
) (ce.Document, error) {
	docClassId, err := primitive.ObjectIDFromHex(documentClassId)
	if err != nil {
		return nil, err
	}

	// ensure document class exists
	findDocClassRes := d.db.documentClass.FindOne(context.TODO(), bson.M{
		"_id": docClassId,
	})
	if findDocClassRes.Err() != nil {
		return nil, errors.New(fmt.Sprintf("Document class %s does not exist", documentClassId))
	}

	var dc DocumentClass
	err = findDocClassRes.Decode(&dc)
	if err != nil {
		return nil, err
	}

	return &Document{
		Type:          ce.ObjectTypeDocument,
		DocumentClass: dc,
		Object: Object{
			Name:        descriptor.Name,
			Label:       descriptor.Label,
			Description: descriptor.Description,
		},
	}, nil
}