package aws

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"errors"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// This package contains the object store implementation for AWS.
//
// Object metadata are stored in AWS' DynamoDB with search-indexing capability provided by
// AWS Elastic Search.
//
// Metadata changes are streamed to AWS Elastic search via DynamoDB CDC. For more details
// see https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Streams.html
//

type ObjectStoreConfig struct {
	DynamoDBURI      string
	DynamoDBUser     string
	DynamoDBPassword string

	ElasticSearchURI      string
	ElasticSearchUser     string
	ElasticSearchPassword string
}

type ObjectStore struct {
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

	err = verifyRequiredParameters(&objStoreConfig)
	if err != nil {
		log.Fatalf("unable to initialize aws object store: %v", err)
	}

	log.Debugf("aws object store config: %v", objStoreConfig)

	// TODO: Initialize connection to storage mediums here

	return &ObjectStore{}
}

func verifyRequiredParameters(objStoreConfig *ObjectStoreConfig) error {
	if objStoreConfig.DynamoDBUser == "" {
		return errors.New("dynamodb user is required")
	}

	return nil
}
