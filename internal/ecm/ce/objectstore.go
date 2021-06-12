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
	NewWorkspace(descriptor ObjectDescriptor) Workspace

	// NewDocumentClass creates an instance of a DocumentClass without persisting it
	NewDocumentClass(descriptor ObjectDescriptor) DocumentClass

	NewPropertyField(descriptor ObjectDescriptor, fieldType FieldType) PropertyField

	// NewDocument returns a new non-persisted instance of a Document of DocumentClass identified by
	// the documentClassId parameter
	//
	// Can return an error when the DocumentClass identified by documentClassId does not exist
	NewDocument(descriptor ObjectDescriptor, documentClassId string) (Document, error)

	// SaveWorkspace persists the given Workspace
	// returns an error when there is an error persisting the workspace
	SaveWorkspace(workspace Workspace) (Workspace, error)

	// SaveDocumentClass persists the given document class
	SaveDocumentClass(documentClass DocumentClass) (DocumentClass, error)

	// GetWorkspaceByObjectId returns a workspace identified by the workspace's unique id
	GetWorkspaceByObjectId(objectId string) (Workspace, error)

	GetWorkspaceByName(name string) (Workspace, error)

	// CheckOut checks a ce.Modifiable Object out
	//
	// Can return an error if the Object identified by objectId is already checked out
	// or the objectId does not correspond to a ce.Modifiable object
	CheckOut(objectId string, owner string) (interface{}, error)

	// CheckIn commits a ce.Modifiable Object back to the repository making it available
	// for modification to other parties.
	//
	// Can return an error if the modifiableObject does not implement the ce.Modifiable interface or
	// the owner in the second parameter does not correspond to the owner who checked out the document
	CheckIn(modifiableObject interface{}, owner string) error

	// CreateFolder persists an Object of type Folder to the underlying store.
	//
	// An error can be thrown when parentId does not correspond to a Folder or a Workspace
	CreateFolder(parentId string, folder Folder) error

	// CreateDocument persists an Object of type Document to the underlying store.
	//
	// An error can be thrown when parentId does not correspond to a Folder or a Workspace
	CreateDocument(parentId string, folder Folder) error

	// GetFolders returns the Folders that are the immediate children of the current Container
	GetFolders() []Folder

	// GetDocuments returns the Document types that the the immediate children
	// of the current Container
	GetDocuments() []Document

	// List lists of Objects that are the immediate children
	// of the current Container. An Object may end up being a Folder or a Document
	List() []Object
}
