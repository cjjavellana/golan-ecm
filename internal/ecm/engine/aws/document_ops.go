package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentOperation struct {
	db *CollectionStore
}

func (d *DocumentOperation) NewGenericDocument(
	descriptor ce.ObjectDescriptor,
	objectType ce.ObjectType,
) (ce.Document, error) {
	return d.NewDocument(descriptor, "", objectType)
}

func (d *DocumentOperation) NewDocument(
	descriptor ce.ObjectDescriptor,
	documentClassId string,
	objectType ce.ObjectType,
) (ce.Document, error) {

	dc, err := findDocClass(d.db.documentClass, documentClassId)
	if err != nil {
		return nil, err
	}

	return &Document{
		Type:          objectType,
		DocumentClass: dc,
		Object: Object{
			Name:        descriptor.Name,
			Label:       descriptor.Label,
			Description: descriptor.Description,
		},
		Version: ce.Version{
			// all documents start at v0.1
			MajorVersion: 0,
			MinorVersion: 1,
		},
		promoteDirty: false,
		Revision:     0,
	}, nil
}

func (d *DocumentOperation) CreateDocument(doc ce.Document) (ce.Document, error) {
	mongoDoc := doc.(*Document)

	if mongoDoc.Type != ce.ObjectTypeWorkspace {
		err := validateParent(d.db.document, mongoDoc)
		if err != nil {
			return mongoDoc, err
		}
	}

	res, err := d.db.document.InsertOne(context.TODO(), mongoDoc)
	if err != nil {
		return mongoDoc, err
	}

	// now that we have persisted the document, reset promoteDirty flag
	mongoDoc.promoteDirty = false

	// set newly inserted id into the document class
	mongoDoc.ID = res.InsertedID.(primitive.ObjectID)
	return mongoDoc, nil
}

func (d *DocumentOperation) UpdateDocument(doc ce.Document) (ce.Document, error) {
	mongoDoc := doc.(*Document)

	err := validateParent(d.db.document, mongoDoc)
	if err != nil {
		return mongoDoc, err
	}

	err = d.keepPreviousVersion(mongoDoc)
	if err != nil {
		return mongoDoc, err
	}

	// if document's version is not explicitly promoted we increment
	// minor version by 1 before we update
	//
	// if promoteDirty == true, means that it's been explicitly incremented
	// we dont touch the document's minor version
	if !mongoDoc.promoteDirty {
		mongoDoc.MinorVersion = mongoDoc.MinorVersion + 1
	}

	// Get the revision on which this modification was based on
	// we'll use this as a filter criteria when we run our update statement.
	// This is because we are going to increment the Revision in our mongoDoc
	currentRevision := mongoDoc.GetRevision()
	mongoDoc.Revision = currentRevision + 1

	_, err = d.db.document.UpdateOne(
		context.TODO(),
		bson.M{
			"_id":      mongoDoc.ID,
			"Revision": currentRevision,
		},
		mongoDoc,
	)
	if err != nil {
		log.Errorf(fmt.Sprintf(
			"update of document %v revision %d failed: %v",
			mongoDoc.ID,
			currentRevision,
			err,
		))

		return mongoDoc, errors.New(fmt.Sprintf(
			`document %v with revision %d does not exist or 
it has been modified since it was last read`,
			mongoDoc.ID,
			currentRevision,
		))
	}

	// now that we have persisted the document, reset promoteDirty flag
	mongoDoc.promoteDirty = false

	return mongoDoc, nil
}

func (d *DocumentOperation) keepPreviousVersion(mongoDoc *Document) error {
	if !mongoDoc.IsVersioningEnabled {
		return nil
	}

	currentRevision := mongoDoc.GetRevision()
	var currentDoc Document
	currentDocRes := d.db.document.FindOne(context.TODO(), bson.M{
		"_id":      mongoDoc.ID,
		"Revision": currentRevision,
	})

	if currentDocRes.Err() != nil {
		return errors.New(fmt.Sprintf(
			"unable to find document %v revision %d: %v",
			mongoDoc.ID,
			currentRevision,
			currentDocRes.Err(),
		))
	}

	err := currentDocRes.Decode(&currentDoc)
	if err != nil {
		return errors.New(fmt.Sprintf(
			"unable to decode document %v revision %d: %v",
			mongoDoc.ID,
			currentRevision,
			err,
		))
	}

	mongoDoc.PreviousVersions = append(
		mongoDoc.PreviousVersions,
		currentDoc,
	)

	return nil
}

// findDocClass returns nothing (nil, nil) when documentClassId is empty
// otherwise it will try to find from the underlying store, returning an
// error when it cannot the DocumentClass identified by the parameter
// documentClassId
func findDocClass(
	docClassCollection *mongo.Collection,
	documentClassId string,
) (*DocumentClass, error) {

	if documentClassId == "" {
		// nothing to do
		return nil, nil
	}

	docClassId, err := primitive.ObjectIDFromHex(documentClassId)
	if err != nil {
		return nil, err
	}

	// ensure document class exists
	findDocClassRes := docClassCollection.FindOne(context.TODO(), bson.M{
		"_id": docClassId,
	})
	if findDocClassRes.Err() != nil {
		return nil, errors.New(fmt.Sprintf("document class %s does not exist", documentClassId))
	}

	var dc DocumentClass
	err = findDocClassRes.Decode(&dc)
	if err != nil {
		return nil, err
	}

	return &dc, nil
}

// validateParent ensures the following:
// 1. That the parent Object exists
// 2. That the parent is of type Folder or Workspace
func validateParent(coll *mongo.Collection, mongoDoc *Document) error {
	parentRes := coll.FindOne(context.TODO(), bson.M{
		"_id": mongoDoc.Parent,
	})
	if parentRes.Err() != nil {
		return errors.New(fmt.Sprintf("unable to find parent identified by id %s", mongoDoc.GetParent()))
	}

	var m map[string]interface{}
	err := parentRes.Decode(&m)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to decode parent: %v", err))
	}

	err = isParentFolderOrWorkspace(m)
	if err != nil {
		return err
	}

	return nil
}

func isParentFolderOrWorkspace(m map[string]interface{}) error {
	if parentType := m["Type"]; parentType == nil ||
		parentType != ce.ObjectTypeWorkspace && parentType != ce.ObjectTypeFolder {
		return errors.New(fmt.Sprintf("invalid parent type: %s", parentType))
	}

	return nil
}
