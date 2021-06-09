package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentClass struct {
	WorkspaceId primitive.ObjectID `bson:"WorkspaceId,omitempty"`

	Name        string `bson:"Name,omitempty"`
	Label       string `bson:"Label,omitempty"`
	Description string `bson:"Description,omitempty"`

	// PropertyFields
	PropertyFields []*PropertyField `bson:"PropertyFields,omitempty"`

	Object `bson:",inline"`
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

func (d *DocumentClass) SetPropertyFields(propertyFields []ce.PropertyField) {
	s := make([]*PropertyField, len(propertyFields))
	for i, v := range propertyFields {
		s[i] = v.(*PropertyField)
	}

	d.PropertyFields = s
}

func (d *DocumentClass) GetPropertyFields() []ce.PropertyField {
	s := make([]ce.PropertyField, len(d.PropertyFields))
	for i, v := range d.PropertyFields {
		s[i] = v
	}

	return s
}
