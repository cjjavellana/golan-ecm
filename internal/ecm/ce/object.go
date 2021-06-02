package ce

import (
	"github.com/google/uuid"
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
	ObjectId() uuid.UUID
	IsDeleted() bool
	Owner() string

	GetObjectType() ObjectType

	SetCreatedBy(user string)
	CreatedBy() string

	SetDateCreated(dateCreated time.Time)
	DateCreated() time.Time

	SetUpdatedBy(user string)
	UpdatedBy() string

	SetDateUpdated(dateUpdated time.Time)
	DateUpdated() time.Time
}
