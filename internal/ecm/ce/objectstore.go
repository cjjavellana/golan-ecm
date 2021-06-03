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

	// NewWorkspace creates an instance (in memory) of Workspace without persisting it.
	//
	// Clients, after obtaining a reference to the instance of the newly created Workspace
	// can set attribute attributes s
	NewWorkspace(name string) Workspace

	// SaveWorkspace persists the given Workspace
	// returns an error when there is an error persisting the workspace
	SaveWorkspace(workspace Workspace) error

	// GetWorkspaceByObjectId returns a workspace identified by the workspace's unique UUIDa
	GetWorkspaceByObjectId(objectId uuid.UUID) Workspace

	GetWorkspaceByName(name string) Workspace

	FindFolder() []Folder

	// FindDocuments returns the documents matching the search criteria
	// TODO: Search Criteria API
	FindDocuments() []Document
}
