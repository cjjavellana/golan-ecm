package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
)

type Workspace struct {
	name string
	*ObjectStore

	Object
}

func (w *Workspace) GetName() string {
	return w.name
}

func (w *Workspace) GetDescription() string {
	panic("implement me")
}

func (Workspace) SetDescription(desc string) {
	panic("implement me")
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
	return ce.ObjectTypeWorkspace
}
