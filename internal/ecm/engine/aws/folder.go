package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Folder struct {
	Type ce.ObjectType `bson:"Type"`

	// ParentId refers to the entity that this Folder belongs to.
	// A folder's parent can be a workspace (for top-level folders) or another folder.
	ParentId primitive.ObjectID `bson:"ParentId,omitempty"`

	Attributes []ce.Attribute `bson:"Attributes,omitempty"`

	// DocumentClass describes the category that this Folder belongs to
	DocumentClass *DocumentClass `bson:"DocumentClass"`
	Object        `bson:",inline"`
}

func (f Folder) SetAttributes(attrs []ce.Attribute) {
	f.Attributes = attrs
}

func (f Folder) GetAttributes() []ce.Attribute {
	return f.Attributes
}

func (f Folder) SetDocumentClass(documentClass ce.DocumentClass) {
	f.DocumentClass = documentClass.(*DocumentClass)
}

func (f Folder) GetDocumentClass() ce.DocumentClass {
	return f.DocumentClass
}
