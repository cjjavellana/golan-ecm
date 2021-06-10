package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
)

// Workspace defines a boundary of a particular process.
// e.g. Sales Department, Finance Department, etc
//
// Workspace implements both the ce.Object and ce.Container interfaces
type Workspace struct {
	// These fields are persisted
	// Data fields are persisted to the underlying store

	Type   ce.ObjectType `bson:"Type"`
	Object `bson:",inline"`

	// Non-persisted, transient fields
	objectStore *ObjectStore
}

func (w *Workspace) AddFolder(folder ce.Folder) error {
	panic("implement me")
}

func (w *Workspace) AddFolders(folders ...ce.Folder) error {
	panic("implement me")
}

func (w *Workspace) AddDocument(document ce.Document) error {
	panic("implement me")
}

func (w *Workspace) AddDocuments(documents ...ce.Document) error {
	panic("implement me")
}

func (w *Workspace) GetFolders() []ce.Folder {
	panic("implement me")
}

func (w *Workspace) GetDocuments() []ce.Document {
	panic("implement me")
}

func (w *Workspace) GetChildren() []ce.Object {
	panic("implement me")
}

func (w *Workspace) GetObjectType() ce.ObjectType {
	return w.Type
}
