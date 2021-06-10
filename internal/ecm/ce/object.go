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

// ObjectDescriptor describes an ecm Object.
//
// All Object(s) in ecm is described minimally by 3 properties, namely:
// Name - An alphanumeric field used to internally described an Object. e.g. IssueDate, MaturityDate, IssueAmount
// Label - A human-readable property used for display purposes e.g. Issue Date, Maturity Date
// Description - A long text that can be used to describe an Object in detail.
type ObjectDescriptor struct {
	Name        string
	Label       string
	Description string
}

// Object represents the root entity from which all models should be derived from.
// Contains fields common for all
type Object interface {
	GetName() string
	GetLabel() string
	GetDescription() string

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
