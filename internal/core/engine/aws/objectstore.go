package aws

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/core/ce"
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
	MongoDBURI      string
	MongoDBUser     string
	MongoDBPassword string

	ElasticSearchURI      string
	ElasticSearchUser     string
	ElasticSearchPassword string
}

type ObjectStore struct {
}

// GetObjectStore returns an instance of AWS ObjectStore
func GetObjectStore(config *cfg.AppConfig) *ObjectStore {
	var objectStoreConfig ObjectStoreConfig

	err := mapstructure.Decode(config.StoreConfig, &objectStoreConfig)
	if err != nil {
		// no point in continuing when we cannot connect to our object store
		log.Fatalf("unable to decode store config: %v", err)
	}

	return &ObjectStore{}
}

func (o *ObjectStore) GetObjectStoreId() uuid.UUID {
	return uuid.New()
}

func (o *ObjectStore) CreateWorkspace(workspace *ce.Workspace) {
	panic("implement me")
}

func (o *ObjectStore) GetWorkspaceByObjectId(objectId uuid.UUID) *ce.Workspace {
	panic("implement me")
}

func (o *ObjectStore) GetWorkspaceByName(name string) *ce.Workspace {
	panic("implement me")
}
