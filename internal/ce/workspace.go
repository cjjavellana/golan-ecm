package ce

// Workspace represents the top-most entity in the content engine. All documents and folders has to belong
// to a workspace
//
//							Workspace
//				_______________________________
//				|               |             |
//             Folder         Folder        Folder
//			___________
//          |          |
//       Document	Document
type Workspace struct {
	// The workspace's unique name.
	// Alphanumeric, dashes and dots. No spaces allowed
	Name        string

	// The human friendly name to refer to this workspace
	Label       string


	Description string

	Object
}
