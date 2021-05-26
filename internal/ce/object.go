package ce

import (
	"github.com/google/uuid"
	"time"
)

// Object represents the root entity from which all models should be derived from.
// Contains fields common for all
type Object struct {
	// Uniquely identifies this object
	ObjectId uuid.UUID

	IsDeleted bool

	Owner string

	// Audit Fields
	CreatedBy   string
	DateCreated time.Time
	UpdatedBy   string
	DateUpdated time.Time
}
