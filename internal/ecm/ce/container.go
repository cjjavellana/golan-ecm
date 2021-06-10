package ce

// Container describes that Object(s) having this property can `contain` other Object(s)
// e.g. A Folder can `contain` other Folder(s) and Document(s)
type Container interface {
	AddFolder(folder Folder) error
	AddFolders(folders ...Folder) error
	AddDocument(document Document) error
	AddDocuments(documents ...Document) error

	// GetFolders returns the Folders that are the immediate children of the current Container
	GetFolders() []Folder

	// GetDocuments returns the Document types that the the immediate children
	// of the current Container
	GetDocuments() []Document

	// GetChildren returns the list of Objects that are the immediate children
	// of the current Container. An Object may end up being a Folder or a Document
	GetChildren() []Object
}
