package builder

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"cjavellana.me/ecm/golan/internal/ecm/pb"
	"reflect"
	"testing"
)

type fakeObjStore struct {
}

func (f fakeObjStore) NewWorkspace(descriptor ce.ObjectDescriptor) ce.Workspace {
	panic("implement me")
}

func (f fakeObjStore) NewGenericFolder(descriptor ce.ObjectDescriptor) ce.Folder {
	panic("implement me")
}

func (f fakeObjStore) NewFolderWithDocumentClass(descriptor ce.ObjectDescriptor, documentClassId string) ce.Folder {
	panic("implement me")
}

func (f fakeObjStore) NewDocumentClass(descriptor ce.ObjectDescriptor) ce.DocumentClass {
	panic("implement me")
}

func (f fakeObjStore) NewPropertyField(descriptor ce.ObjectDescriptor, fieldType ce.FieldType) ce.PropertyField {
	return ce.PropertyField{
		Name:        descriptor.Name,
		Label:       descriptor.Label,
		Description: descriptor.Description,
		FieldType:   fieldType,
	}
}

func (f fakeObjStore) NewDocument(descriptor ce.ObjectDescriptor, documentClassId string) (ce.Document, error) {
	panic("implement me")
}

func (f fakeObjStore) SaveWorkspace(workspace ce.Workspace) (ce.Workspace, error) {
	panic("implement me")
}

func (f fakeObjStore) SaveDocumentClass(documentClass ce.DocumentClass) (ce.DocumentClass, error) {
	panic("implement me")
}

func (f fakeObjStore) GetWorkspaceByObjectId(objectId string) (ce.Workspace, error) {
	panic("implement me")
}

func (f fakeObjStore) GetWorkspaceByName(name string) (ce.Workspace, error) {
	panic("implement me")
}

func (f fakeObjStore) CheckOut(objectId string, owner string) (interface{}, error) {
	panic("implement me")
}

func (f fakeObjStore) CheckIn(modifiableObject interface{}, owner string) error {
	panic("implement me")
}

func (f fakeObjStore) CreateFolder(folder ce.Folder) (ce.Folder, error) {
	panic("implement me")
}

func (f fakeObjStore) CreateDocument(document ce.Document) (ce.Document, error) {
	panic("implement me")
}

func (f fakeObjStore) GetFolders() []ce.Folder {
	panic("implement me")
}

func (f fakeObjStore) GetDocuments() []ce.Document {
	panic("implement me")
}

func (f fakeObjStore) List() []ce.Object {
	panic("implement me")
}

func TestPropertyFieldHierarchy_Build(t *testing.T) {
	type fields struct {
		ObjectStore ce.ObjectStore
	}
	type args struct {
		gRpcPropertyFields []*pb.PropertyField
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*ce.PropertyField
	}{
		{
			name:   "itShouldBuildPropertyFieldTree",
			fields: fields{ObjectStore: &fakeObjStore{}},
			args: args{
				gRpcPropertyFields: []*pb.PropertyField{
					{
						Name:        "Amount",
						Label:       "Total Amount",
						Description: "Total Amount Due",
						FieldType:   pb.FieldType_DOUBLE,
					},
					{
						Name:        "LineItems",
						Label:       "Line Items",
						Description: "Items Purchased",
						FieldType:   pb.FieldType_COMPLEX,
						Repeatable:  true,
						SubProperties: []*pb.PropertyField{
							{
								Name:        "Qty",
								Label:       "Quantity",
								Description: "Quantity",
								FieldType:   pb.FieldType_DOUBLE,
							},
							{
								Name:        "ItemDescription",
								Label:       "Item Description",
								Description: "Item Description",
								FieldType:   pb.FieldType_STRING,
							},
							{
								Name:        "UnitCost",
								Label:       "Unit Cost",
								Description: "Price Per Item",
								FieldType:   pb.FieldType_DOUBLE,
							},
						},
					},
				},
			},
			want: []*ce.PropertyField{
				{
					Name:        "Amount",
					Label:       "Total Amount",
					Description: "Total Amount Due",
					FieldType:   ce.FieldTypeDouble,
				},
				{
					Name:        "LineItems",
					Label:       "Line Items",
					Description: "Items Purchased",
					FieldType:   ce.FieldTypeComplex,
					Repeatable:  true,
					PropertyField: []*ce.PropertyField{
						{
							Name:        "Qty",
							Label:       "Quantity",
							Description: "Quantity",
							FieldType:   ce.FieldTypeDouble,
						},
						{
							Name:        "ItemDescription",
							Label:       "Item Description",
							Description: "Item Description",
							FieldType:   ce.FieldTypeString,
						},
						{
							Name:        "UnitCost",
							Label:       "Unit Cost",
							Description: "Price Per Item",
							FieldType:   ce.FieldTypeDouble,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PropertyFieldHierarchy{
				ObjectStore: tt.fields.ObjectStore,
			}
			if got := p.Build(tt.args.gRpcPropertyFields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
