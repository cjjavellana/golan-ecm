package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentClass struct {
	// WorkspaceId refers to the object id of the workspace that this DocumentClass belongs to
	WorkspaceId primitive.ObjectID `bson:"WorkspaceId,omitempty"`

	// PropertyFields
	PropertyFields []*ce.PropertyField `bson:"PropertyFields,omitempty"`

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

func (d *DocumentClass) SetPropertyFields(propertyFields []ce.PropertyField) {
	s := make([]*ce.PropertyField, len(propertyFields))
	for i, v := range propertyFields {
		s[i] = &v
	}

	d.PropertyFields = s
}

func (d *DocumentClass) GetPropertyFields() []ce.PropertyField {
	s := make([]ce.PropertyField, len(d.PropertyFields))
	for i, v := range d.PropertyFields {
		s[i] = *v
	}

	return s
}
