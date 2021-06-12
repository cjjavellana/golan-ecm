package ce

// Folder is a logical container used for holding Documents
//
// A Folder can have a DocumentClass and have custom attributes
type Folder interface {
	Object

	SetAttributes(attrs []Attribute)
	GetAttributes() []Attribute

	SetDocumentClass(documentClass DocumentClass)
	GetDocumentClass() DocumentClass
}
