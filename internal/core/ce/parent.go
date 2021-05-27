package ce

import "github.com/google/uuid"

// RefType describes the type of reference an `Object` has with its Parent
type RefType string

const (
	// RefTypeSoft means that the relationship between the Object and its Parent
	// exists logically. You can imagine RefTypeSoft as some sort of a pointer to the Parent
	RefTypeSoft RefType = "soft"

	// RefTypeHard represents a belongs-to relationship - That is, this Object belongs to this Parent
	RefTypeHard = "hard"
)

type Parent interface {
	// GetObjectId Returns the ObjectId of the Parent Object
	GetObjectId() uuid.UUID

	// GetRefType returns the relationship this Object has with its Parent
	GetRefType() RefType
}
