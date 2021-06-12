package aws

import "cjavellana.me/ecm/golan/internal/ecm/ce"

type Document struct {
	IsVersioningEnabled   bool            `bson:"IsVersioningEnabled"`
	Type                  ce.ObjectType   `bson:"Type"`
	ContentType           string          `bson:"ContentType"`
	Attributes            []*ce.Attribute `bson:"Attribute"`
	HasUnderlyingDocument bool            `bson:"HasUnderlyingDocument"`
	SizeInBytes           uint64          `bson:"SizeInBytes"`
	Content               []byte          `bson:"Content"`
	Filename              string          `bson:"Filename"`

	Version
	ce.Modifiable

	// DocumentClass describes the category that this Folder belongs to
	DocumentClass `bson:"DocumentClass"`
	Object        `bson:",inline"`
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
