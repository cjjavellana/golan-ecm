package aws

import "cjavellana.me/ecm/golan/internal/ecm/ce"

type Document struct {
	IsVersioningEnabled bool
	Type ce.ObjectType `bson:"Type"`

	Version
	ce.Modifiable

	// DocumentClass describes the category that this Folder belongs to
	DocumentClass `bson:"DocumentClass"`
	Object        `bson:",inline"`

	// Non-persisted, transient fields
	objectStore *ObjectStore
}

func (d *Document) EnableVersioning() {
	d.IsVersioningEnabled = true
}

func (d *Document) VersionEnabled() bool {
	return d.IsVersioningEnabled
}

func (d *Document) SetAttributes(attrs []ce.Attribute) {
	panic("implement me")
}

func (d *Document) GetAttributes() []ce.Attribute {
	panic("implement me")
}

func (d *Document) SetFilename(filename string) {
	panic("implement me")
}

func (d *Document) GetFilename() string {
	panic("implement me")
}

func (d *Document) GetSize() uint64 {
	panic("implement me")
}

func (d *Document) SetContentType(contentType string) {
	panic("implement me")
}

func (d *Document) GetContentType() string {
	panic("implement me")
}

func (d *Document) HasUnderlyingDocument() bool {
	panic("implement me")
}

func (d *Document) SetUnderlyingDocument(document []byte) {
	panic("implement me")
}

func (d *Document) GetUnderlyingDocument() []byte {
	panic("implement me")
}
