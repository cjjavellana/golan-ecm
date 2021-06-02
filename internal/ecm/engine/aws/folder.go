package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
)

type Folder struct {
	name        string
	label       string
	description string

	Object
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) GetLabel() string {
	return f.label
}

func (f *Folder) GetDescription() string {
	panic("implement me")
}

func (f *Folder) SetPropertyFields(attrs []ce.PropertyField) {
	panic("implement me")
}

func (f *Folder) GetPropertyFields() []ce.PropertyField {
	panic("implement me")
}

func (Folder) GetObjectType() ce.ObjectType {
	return ce.ObjectTypeFolder
}

func (Folder) AddDocument(document ce.Document) error {
	panic("implement me")
}
