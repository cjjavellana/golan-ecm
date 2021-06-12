package aws

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"context"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// This package contains the object store implementation for AWS.
//
// Object metadata are stored in AWS' DynamoDB with search-indexing capability provided by
// AWS Elastic Search.
//
// Metadata changes are streamed to AWS Elastic search via DynamoDB CDC. For more details
// see https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Streams.html
//

type DB struct {
	URI          string `validate:"required"`
	User         string `validate:"required"`
	Password     string `validate:"required"`
	DatabaseName string `validate:"required"`
}

type ElasticSearch struct {
	URI      string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
}

// ObjectStoreConfig in the format of:
//
// storeconfig:
//  db:
//    uri: mongodb://localhost:27017
//    user: root
//    password: example
//    databasename: golan
//  elasticsearch:
//    uri: elasticsearchuri
//    user: elasticsearchuser
//    password: elasticsearchpassword
type ObjectStoreConfig struct {
	DB            DB
	ElasticSearch ElasticSearch
}

type CollectionStore struct {
	document      *mongo.Collection
	documentClass *mongo.Collection
}

type ObjectStore struct {
	db *CollectionStore

	docOps       *DocumentOperation
	docClassOps  *DocumentClassOperation
	workspaceOps *WorkspaceOperation
	propFieldOps *PropertyFieldOperation
}

func (o *ObjectStore) GetObjectStoreId() uuid.UUID {
	// TODO: Implement support for multi-object store aka multi-tenant
	return uuid.New()
}

func (o *ObjectStore) NewPropertyField(
	descriptor ce.ObjectDescriptor,
	fieldType ce.FieldType,
) ce.PropertyField {
	return o.propFieldOps.NewPropertyField(descriptor, fieldType)
}

func (o *ObjectStore) NewWorkspace(
	descriptor ce.ObjectDescriptor,
) ce.Workspace {
	return o.workspaceOps.NewWorkspace(descriptor)
}

func (o *ObjectStore) NewDocumentClass(
	descriptor ce.ObjectDescriptor,
) ce.DocumentClass {
	return o.docClassOps.NewDocumentClass(descriptor)
}

func (o *ObjectStore) NewDocument(
	descriptor ce.ObjectDescriptor,
	documentClassId string,
) (ce.Document, error) {
	return o.docOps.NewDocument(descriptor, documentClassId)
}

func (o *ObjectStore) SaveWorkspace(workspace ce.Workspace) (ce.Workspace, error) {
	return o.workspaceOps.SaveWorkspace(workspace)
}

func (o *ObjectStore) SaveDocumentClass(documentClass ce.DocumentClass) (ce.DocumentClass, error) {
	return o.docClassOps.SaveDocumentClass(documentClass)
}

func (o *ObjectStore) GetWorkspaceByObjectId(objectId string) (ce.Workspace, error) {
	return o.workspaceOps.GetWorkspaceByObjectId(objectId)
}

func (o *ObjectStore) GetWorkspaceByName(name string) (ce.Workspace, error) {
	panic("implement me")
}

func (o *ObjectStore) CheckOut(objectId string, owner string) (interface{}, error) {
	panic("implement me")
}

func (o *ObjectStore) CheckIn(modifiableObject interface{}, owner string) error {
	panic("implement me")
}

func (o *ObjectStore) CreateFolder(parentId string, folder ce.Folder) error {
	panic("implement me")
}

func (o *ObjectStore) CreateDocument(parentId string, folder ce.Folder) error {
	panic("implement me")
}

func (o *ObjectStore) GetFolders() []ce.Folder {
	panic("implement me")
}

func (o *ObjectStore) GetDocuments() []ce.Document {
	panic("implement me")
}

func (o *ObjectStore) List() []ce.Object {
	panic("implement me")
}

// GetObjectStore returns an instance of AWS ObjectStore
func GetObjectStore(config *cfg.AppConfig) *ObjectStore {
	var objStoreConfig ObjectStoreConfig

	err := mapstructure.Decode(config.StoreConfig, &objStoreConfig)
	if err != nil {
		// no point in continuing when we cannot connect to our object store
		log.Fatalf("unable to decode store config: %v", err)
	}

	err = validateObjectStoreConfig(&objStoreConfig)
	if err != nil {
		log.Fatalf("unable to initialize aws object store: %v", err)
	}

	log.Debugf("aws object store config: %v", objStoreConfig)

	mongoClient := initDb(&objStoreConfig)
	database := mongoClient.Database(objStoreConfig.DB.DatabaseName)
	docCollection := getMongoCollection(database, "documents")
	docClassCollection := getMongoCollection(database, "document_class")

	createIndexOnName(
		docCollection,
		docClassCollection,
	)

	collStore := &CollectionStore{
		document:      docCollection,
		documentClass: docClassCollection,
	}

	return &ObjectStore{
		db: collStore,
		propFieldOps: &PropertyFieldOperation{
			db: collStore,
		},
		docClassOps: &DocumentClassOperation{
			db: collStore,
		},
		docOps: &DocumentOperation{
			db: collStore,
		},
		workspaceOps: &WorkspaceOperation{
			db: collStore,
		},
	}
}

func createIndexOnName(collections ...*mongo.Collection) {
	for _, v := range collections {
		if idxName, err := createIndex(v, bson.M{"Name": 1}); err != nil {
			log.Errorf("unable to create index for %s collection: %v", v.Name(), err)
		} else {
			log.Infof("index %s created successfully on %s", idxName, v.Name())
		}
	}
}

func getMongoCollection(database *mongo.Database, collection string) *mongo.Collection {
	c := database.Collection(collection)
	if c == nil {

		err := database.CreateCollection(context.TODO(), collection)
		if err != nil {
			log.Fatalf("unable to create %s collection: %v", collection, err)
		}

		return database.Collection(collection)
	}

	return c
}

func validateObjectStoreConfig(objStoreConfig *ObjectStoreConfig) error {
	validate := validator.New()
	err := validate.Struct(objStoreConfig)

	if err != nil {
		return err
	}

	return nil
}
