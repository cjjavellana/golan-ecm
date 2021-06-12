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
type Workspace interface {
	Object
}
