package ce

import "github.com/google/uuid"

type AttributeType string

const (
	AttributeTypeString   AttributeType = "string"
	AttributeTypeInteger                = "integer"
	AttributeTypeDouble                 = "double"
	AttributeTypeDateTime               = "date"
)

// CustomAttribute represents an attribute that is defined by the end-user at run time
type CustomAttribute interface {
	// GetName returns the name of the attribute. Alphanumeric characters only.
	// interestRate, issueDate
	GetName() string
	GetLabel() string
	GetAttributeType() AttributeType
	GetDescription() string

	// GetWorkspaceObjectId returns the ObjectId of the workspace that this attribute belongs to
	GetWorkspaceObjectId() uuid.UUID

	Object
}
