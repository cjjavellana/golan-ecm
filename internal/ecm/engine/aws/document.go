package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
)

// Workspace defines a boundary of a particular (business) process.
// e.g. Sales Department, Finance Department, etc
type Workspace struct {
	*Document `bson:",inline"`
}

type Folder struct {
	*Document `bson:",inline"`
}

type Document struct {
	IsVersioningEnabled   bool            `bson:"IsVersioningEnabled"`
	Type                  ce.ObjectType   `bson:"Type"`
	ContentType           string          `bson:"ContentType"`
	Attributes            []*ce.Attribute `bson:"Attribute"`
	HasUnderlyingDocument bool            `bson:"HasUnderlyingDocument"`
	SizeInBytes           uint64          `bson:"SizeInBytes"`
	Content               []byte          `bson:"Content"`
	Filename              string          `bson:"Filename"`
	PreviousVersions      []Document      `bson:"PreviousVersions"`

	ce.LockStatus `bson:"LockStatus"`

	// Version is a human-friendly versioning scheme e.g. 1.1, 1.2, 1.3, etc
	ce.Version `bson:"Version"`

	// Revision is used as an optimistic lock
	Revision uint32 `bson:"Revision"`

	// DocumentClass describes the category that this Folder belongs to
	*DocumentClass `bson:"DocumentClass"`
	Object         `bson:",inline"`

	// flag that makes the Promote function idempotent in-between persists
	promoteDirty bool
}

func (d *Document) EnableVersioning() {
	d.IsVersioningEnabled = true
}

func (d *Document) VersionEnabled() bool {
	return d.IsVersioningEnabled
}

func (d *Document) SetAttributes(attrs []*ce.Attribute) {
	d.Attributes = attrs
}

func (d *Document) GetAttributes() []*ce.Attribute {
	return d.Attributes
}

func (d *Document) SetFilename(filename string) {
	d.Filename = filename
}

func (d *Document) GetFilename() string {
	return d.Filename
}

func (d *Document) GetSize() uint64 {
	return d.SizeInBytes
}

func (d *Document) SetContentType(contentType string) {
	d.ContentType = contentType
}

func (d *Document) GetContentType() string {
	return d.ContentType
}

func (d *Document) GetHasUnderlyingDocument() bool {
	return d.HasUnderlyingDocument
}

func (d *Document) SetUnderlyingDocument(document []byte) {
	d.Content = document
	d.HasUnderlyingDocument = len(document) > 0
}

func (d *Document) GetUnderlyingDocument() []byte {
	return d.Content
}

func (d *Document) PromoteVersion() ce.Version {
	if !d.promoteDirty {
		v := d.Version
		v.MajorVersion = v.MajorVersion + 1
		v.MinorVersion = 0
	}

	return d.Version
}

func (d *Document) GetVersion() ce.Version {
	return d.Version
}

func (d *Document) CheckOut(owner string) {
	panic("implement me")
}

func (d *Document) CheckIn(owner string) {
	panic("implement me")
}

func (d *Document) GetRevision() uint32 {
	return d.Revision
}
