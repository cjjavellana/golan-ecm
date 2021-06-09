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

	// NewWorkspace creates an instance (in memory) of a Workspace without persisting it.
	//
	// Clients can set additional attributes after obtaining a reference to the Workspace instance
	// before calling SaveWorkspace
	NewWorkspace(name string, description string) Workspace

	// NewDocumentClass creates an instance of a DocumentClass without persisting it
	NewDocumentClass(name string, label string, description string) DocumentClass

	NewPropertyField(name string, label string, fieldType FieldType, description string) PropertyField

	// SaveWorkspace persists the given Workspace
	// returns an error when there is an error persisting the workspace
	SaveWorkspace(workspace Workspace) (Workspace, error)

	// SaveDocumentClass persists the given document class
	SaveDocumentClass(documentClass DocumentClass) (DocumentClass, error)

	// GetWorkspaceByObjectId returns a workspace identified by the workspace's unique id
	GetWorkspaceByObjectId(objectId string) (Workspace, error)

	GetWorkspaceByName(name string) (Workspace, error)

	FindFolder() []Folder

	// FindDocuments returns the documents matching the search criteria
	// TODO: Search Criteria API
	FindDocuments() []Document
}
