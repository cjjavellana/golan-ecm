package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Object struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	IsDeleted bool               `bson:"IsDeleted"`

	Name        string `bson:"Name,omitempty"`
	Label       string `bson:"Label,omitempty"`
	Description string `bson:"Description,omitempty"`

	Owner  string             `bson:"Owner,omitempty"`
	Parent primitive.ObjectID `bson:"Parent,omitempty"`

	CreatedBy   string     `bson:"CreatedBy,omitempty"`
	DateCreated *time.Time `bson:"DateCreated,omitempty"`
	UpdatedBy   string     `bson:"UpdatedBy,omitempty"`
	DateUpdated *time.Time `bson:"DateUpdated,omitempty"`
}

func (o *Object) ObjectId() string {
	return o.ID.Hex()
}

func (o *Object) GetName() string {
	return o.Name
}

func (o *Object) GetLabel() string {
	return o.Label
}

func (o *Object) GetDescription() string {
	return o.Description
}

func (o *Object) GetIsDeleted() bool {
	return o.IsDeleted
}

func (o *Object) GetOwner() string {
	return o.Owner
}

func (o *Object) SetCreatedBy(user string) {
	o.CreatedBy = user
}

func (o *Object) GetCreatedBy() string {
	return o.CreatedBy
}

func (o *Object) SetDateCreated(dateCreated *time.Time) {
	o.DateCreated = dateCreated
}

func (o *Object) GetDateCreated() *time.Time {
	return o.DateCreated
}

func (o *Object) SetUpdatedBy(user string) {
	o.UpdatedBy = user
}

func (o *Object) GetUpdatedBy() string {
	return o.UpdatedBy
}

func (o *Object) SetDateUpdated(dateUpdated *time.Time) {
	o.DateUpdated = dateUpdated
}

func (o *Object) GetDateUpdated() *time.Time {
	return o.DateUpdated
}

func (o *Object) GetObjectType() ce.ObjectType {
	panic("this must be overridden by structures embedding this Object")
}

func (o *Object) GetParent() string {
	return o.Parent.Hex()
}

func (o *Object) SetParent(objectId string) error {
	id, ok := primitive.ObjectIDFromHex(objectId)
	if ok != nil {
		return ok
	}

	o.Parent = id
	return nil
}
