package ce

// CustomAttribute represents an attribute that is defined by the end-user at run time
type CustomAttribute struct {
	Name        string
	Label       string
	Description string

	// A custom attribute must belong to a workspace
	Workspace
	Object
}
