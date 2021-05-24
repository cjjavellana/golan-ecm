package ce

import (
	"github.com/google/uuid"
	"time"
)

// Object represents as the base entity from which all models are derived from.
// Contains fields common for all
type Object struct {
	// Uniquely identifies this object
	ObjectId uuid.UUID

	CreatedBy   string
	DateCreated time.Time
	UpdatedBy   string
	DateUpdated time.Time
}
