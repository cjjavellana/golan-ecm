package ce

type DocStatusType string

type Document interface {
	EnableVersioning()
	VersionEnabled() bool

	SetAttributes(attrs []*Attribute)
	GetAttributes() []*Attribute

	SetFilename(filename string)
	GetFilename() string

	// GetSize Returns the size of the document in bytes.
	// Derived when SetUnderlyingDocument is called.
	//
	// Can be 0 when this Document is a virtual document - A document with
	// attributes but no underlying file
	GetSize() uint64

	SetContentType(contentType string)
	GetContentType() string

	// GetHasUnderlyingDocument returns true when this Document is backed
	// by an underlying document e.g. pdf, word, xls, xml, etc
	GetHasUnderlyingDocument() bool

	// SetUnderlyingDocument sets the underlying document for this document
	SetUnderlyingDocument(document []byte)
	GetUnderlyingDocument() []byte

	Object
	DocumentClass
	Modifiable
	Version
}
