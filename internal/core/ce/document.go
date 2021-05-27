package ce

type Attribute interface {
	CustomAttribute
	SetValue(value interface{})
	GetValue() interface{}
}

type Document interface {
	EnableVersioning()
	VersionEnabled() bool

	// SetAttributes defines the attributes of a document
	SetAttributes(attributes []Attribute) error
	GetAttributes() []Attribute

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

	// HasUnderlyingDocument returns true when this Document is backed
	// by an underlying document e.g. pdf, word, xls, xml, etc
	HasUnderlyingDocument() bool

	// SetUnderlyingDocument sets the underlying document for this document
	SetUnderlyingDocument(document []byte)
	GetUnderlyingDocument() []byte

	Object
	DocumentClass
	Version
}
