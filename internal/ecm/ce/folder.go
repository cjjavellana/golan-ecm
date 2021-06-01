package ce

// Folder is a logical container used for holding Documents
//
// A Folder can have a DocumentClass and have custom attributes
type Folder interface {
	// DocumentClass defines the document category of a document.
	DocumentClass
	Object

	// AddDocument creates a document in this Folder
	AddDocument(document Document) error
}
