package ce

type DocStatusType string

type Version struct {
	MajorVersion uint32
	MinorVersion uint32
}

type Document interface {
	EnableVersioning()
	VersionEnabled() bool

	SetAttributes(attrs []*Attribute)
	GetAttributes() []*Attribute

	SetFilename(filename string)
	GetFilename() string

	// PromoteVersion increments the Document's major version
	// and resetting the minor version to 0
	PromoteVersion() Version
	GetVersion() Version

	GetRevision() uint32

	CheckOut(owner string)
	CheckIn(owner string)

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
}
