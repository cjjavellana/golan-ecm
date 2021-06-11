package ce

type LockStatusType string

const (
	LockStatusLocked   LockStatusType = "locked"
	LockStatusUnlocked                = "unlocked"
)

type LockStatus struct {
	Owner  string
	Status LockStatusType
}

// Modifiable interface indicates that an Object(s) property can be modified.
//
// Modifiable provides functions for locking / unlocking Object(s)
type Modifiable interface {
	// LockStatus returns the status of an Object
	LockStatus() LockStatus

	// Lock marks an Object for exclusive modification of the `owner`
	//
	// Owner can be any string. The same value must be used to unlock the document
	Lock(owner string) error

	// Unlock releases a Lock to the Object.
	// owner must be the same value as the one used to Lock the Object
	Unlock(owner string) error
}
