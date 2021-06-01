package ce

import "github.com/google/uuid"

// ObjectStore is responsible for providing storage and retrieval functionalities.
//
// An instance of this service can only connect to a single ObjectStore. ObjectStore supports the following
// storage mediums:
//  1. On AWS, S3 + MySQL
//  2. On-premise MongoDB
//  3. On-premise MariaDB
type ObjectStore interface {
	GetObjectStoreId() uuid.UUID

	// CreateWorkspace Creates a new workspace identified by the name parameter
	CreateWorkspace(workspace *Workspace)

	// GetWorkspaceByObjectId returns a workspace identified by the workspace's unique UUIDa
	GetWorkspaceByObjectId(objectId uuid.UUID) *Workspace

	GetWorkspaceByName(name string) *Workspace
}
