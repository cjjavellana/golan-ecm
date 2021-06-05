package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentClass struct {
	WorkspaceId primitive.ObjectID

	Name        string
	Label       string
	Description string

	Object
}

func (d *DocumentClass) GetWorkspaceId() string {
	return d.WorkspaceId.Hex()
}

func (d *DocumentClass) SetWorkspaceId(objectId string) error {
	id, ok := primitive.ObjectIDFromHex(objectId)
	if ok != nil {
		return ok
	}

	d.WorkspaceId = id
	return nil
}

func (d *DocumentClass) GetName() string {
	return d.Name
}

func (d *DocumentClass) GetLabel() string {
	return d.Label
}

func (d *DocumentClass) GetDescription() string {
	return d.Description
}

func (d *DocumentClass) SetPropertyFields(attrs []ce.PropertyField) {
	panic("implement me")
}

func (d *DocumentClass) GetPropertyFields() []ce.PropertyField {
	panic("implement me")
}
