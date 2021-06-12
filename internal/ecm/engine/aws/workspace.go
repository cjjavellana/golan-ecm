package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
)

// Workspace defines a boundary of a particular process.
// e.g. Sales Department, Finance Department, etc
//
// Workspace implements both the ce.Object and ce.Container interfaces
type Workspace struct {
	// These fields are persisted
	// Data fields are persisted to the underlying store

	Type   ce.ObjectType `bson:"Type"`
	Object `bson:",inline"`
}

func (w *Workspace) GetObjectType() ce.ObjectType {
	return w.Type
}
