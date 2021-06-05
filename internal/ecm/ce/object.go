package ce

import (
	"time"
)

type ObjectType string

const (
	ObjectTypeFolder    ObjectType = "folder"
	ObjectTypeWorkspace            = "workspace"
	ObjectTypeDocument             = "document"
)

// Object represents the root entity from which all models should be derived from.
// Contains fields common for all
type Object interface {
	ObjectId() string

	GetIsDeleted() bool
	GetOwner() string

	GetObjectType() ObjectType

	// GetParent returns the ObjectId of the parent of this Object
	//
	// Maybe nil if this object does not have a parent e.g. Workspace
	GetParent() string
	SetParent(objectId string) error

	SetCreatedBy(user string)
	GetCreatedBy() string

	SetDateCreated(dateCreated *time.Time)
	GetDateCreated() *time.Time

	SetUpdatedBy(user string)
	GetUpdatedBy() string

	SetDateUpdated(dateUpdated *time.Time)
	GetDateUpdated() *time.Time
}
