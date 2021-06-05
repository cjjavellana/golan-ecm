package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Object struct {
	ID        primitive.ObjectID `bson:"_id"`
	isDeleted bool

	owner  string
	parent string

	createdBy   string    `bson:"CreatedBy"`
	dateCreated time.Time `bson:"DateCreated"`
	updatedBy   string    `bson:"UpdatedBy"`
	dateUpdated time.Time `bson:"DateUpdated"`
}

func (o *Object) ObjectId() interface{} {
	return o.ID.Hex()
}

func (o *Object) SetObjectId(objectId interface{}) {
	o.ID = objectId.(primitive.ObjectID)
}

func (o *Object) IsDeleted() bool {
	return o.isDeleted
}

func (o *Object) Owner() string {
	return o.owner
}

func (o *Object) SetCreatedBy(user string) {
	o.createdBy = user
}

func (o *Object) CreatedBy() string {
	return o.createdBy
}

func (o *Object) SetDateCreated(dateCreated time.Time) {
	o.dateCreated = dateCreated
}

func (o *Object) DateCreated() time.Time {
	return o.dateCreated
}

func (o *Object) SetUpdatedBy(user string) {
	o.updatedBy = user
}

func (o *Object) UpdatedBy() string {
	return o.updatedBy
}

func (o *Object) SetDateUpdated(dateUpdated time.Time) {
	o.dateUpdated = dateUpdated
}

func (o *Object) DateUpdated() time.Time {
	return o.dateUpdated
}

func (o *Object) GetObjectType() ce.ObjectType {
	panic("this must be overridden by structures embedding this object")
}

func (o *Object) GetParent() string {
	return o.parent
}

func (o *Object) SetParent(objectId string) {
	o.parent = objectId
}
