package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Folder struct {
	Type ce.ObjectType `bson:"Type"`

	// WorkspaceId refers to the object id of the workspace that this DocumentClass belongs to
	WorkspaceId primitive.ObjectID `bson:"WorkspaceId,omitempty"`

	// ParentId refers to the entity that this Folder belongs to.
	// A folder's parent can be a workspace (for top-level folders) or another folder.
	ParentId primitive.ObjectID `bson:"ParentId,omitempty"`

	// DocumentClass describes the category that this Folder belongs to
	DocumentClass `bson:"DocumentClass"`
	Object        `bson:",inline"`

	// Non-persisted, transient fields
	objectStore *ObjectStore
}

func (f *Folder) GetWorkspaceId() string {
	panic("implement me")
}

func (f *Folder) SetWorkspaceId(objectId string) error {
	panic("implement me")
}

func (f *Folder) AddFolder(folder ce.Folder) error {
	panic("implement me")
}

func (f *Folder) AddFolders(folders ...ce.Folder) error {
	panic("implement me")
}

func (f *Folder) AddDocument(document ce.Document) error {
	panic("implement me")
}

func (f *Folder) AddDocuments(documents ...ce.Document) error {
	panic("implement me")
}

func (f *Folder) GetFolders() []ce.Folder {
	panic("implement me")
}

func (f *Folder) GetDocuments() []ce.Document {
	panic("implement me")
}

func (f *Folder) GetChildren() []ce.Object {
	panic("implement me")
}

