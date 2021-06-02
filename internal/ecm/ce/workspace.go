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

	// GetName Returns the name of the Workspace.
	// e.g. Legal & Sons (Law Firm)
	GetName() string

	// GetDescription returns the description for this Workspace
	GetDescription() string
	SetDescription(desc string)

	AddFolder(folder Folder) error
	AddFolders(folders ...Folder) error
	AddDocument(document Document) error
	AddDocuments(documents ...Document) error

	// GetFolders returns the Folders that are the immediate children of the
	// Workspace
	GetFolders() []Folder

	// GetDocuments returns the Document types that the the immediate children
	// of the Workspace
	GetDocuments() []Document

	// GetChildren returns the list of Objects that are the immediate children
	// of the Workspace. An Object may end up being a Folder or a Document
	GetChildren() []Object

	Object
}
