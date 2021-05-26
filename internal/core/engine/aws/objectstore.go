package aws

import (
	"cjavellana.me/ecm/golan/internal/core/ce"
	"github.com/google/uuid"
)

// This package contains the object store implementation for AWS.
//
// Object metadata are stored in AWS' DynamoDB with search-indexing capability provided by
// AWS Elastic Search.
//
// Metadata changes are streamed to AWS Elastic search via DynamoDB CDC. For more details
// see https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Streams.html
//

type ObjectStore struct {
}

// GetObjectStore returns an instance of AWS ObjectStore
func GetObjectStore() *ObjectStore {
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
