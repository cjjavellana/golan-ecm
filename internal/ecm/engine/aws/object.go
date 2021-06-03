package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"github.com/google/uuid"
	"time"
)

type Object struct {
	objectId  uuid.UUID
	isDeleted bool

	owner  string
	parent uuid.UUID

	createdBy   string
	dateCreated time.Time
	updatedBy   string
	dateUpdated time.Time
}

func (o *Object) ObjectId() uuid.UUID {
	return o.objectId
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

func (o *Object) GetParent() uuid.UUID {
	return o.parent
}

func (o *Object) SetParent(objectId uuid.UUID) {
	o.parent = objectId
}