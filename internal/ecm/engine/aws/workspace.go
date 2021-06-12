package aws

// Workspace defines a boundary of a particular (business) process.
// e.g. Sales Department, Finance Department, etc
type Workspace struct {
	*Document `bson:",inline"`
}
