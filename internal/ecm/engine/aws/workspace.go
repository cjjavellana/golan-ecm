package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
)

type Workspace struct {
	Name        string `bson:"Name"`
	Description string `bson:"Description"`

	*ObjectStore

	Object `bson:",inline"`
}

func (w *Workspace) GetName() string {
	return w.Name
}

func (w *Workspace) GetDescription() string {
	return w.Description
}

func (w *Workspace) SetDescription(desc string) {
	w.Description = desc
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
