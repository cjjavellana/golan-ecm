package aws

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
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

type ObjectStoreConfig struct {
	DB            DB
	ElasticSearch ElasticSearch
}

type ObjectStore struct {
	mongoClient *mongo.Client
	database    *mongo.Database
}

func (o *ObjectStore) FindFolder() []ce.Folder {
	panic("implement me")
}

func (o *ObjectStore) FindDocuments() []ce.Document {
	panic("implement me")
}

func (o *ObjectStore) GetObjectStoreId() uuid.UUID {
	return uuid.New()
}

func (o *ObjectStore) NewWorkspace(name string) ce.Workspace {
	workspaceObjId := uuid.New()

	log.Debugf("creating disconnected workspace: %v", workspaceObjId.String())

	return &Workspace{
		Object: Object{
			objectId: workspaceObjId,
		},

		name: name,
	}
}

func (o *ObjectStore) SaveWorkspace(workspace ce.Workspace) {
	panic("implement me")
}

func (o *ObjectStore) GetWorkspaceByObjectId(objectId uuid.UUID) ce.Workspace {
	panic("implement me")
}

func (o *ObjectStore) GetWorkspaceByName(name string) ce.Workspace {
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

	return &ObjectStore{
		mongoClient: mongoClient,
		database:    database,
	}
}

func validateObjectStoreConfig(objStoreConfig *ObjectStoreConfig) error {
	validate := validator.New()
	err := validate.Struct(objStoreConfig)

	if err != nil {
		return err
	}

	return nil
}