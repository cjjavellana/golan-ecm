package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestDocumentClass_SetPropertyFields(t *testing.T) {
	type fields struct {
		WorkspaceId    primitive.ObjectID
		PropertyFields []*PropertyField
		Object         Object
	}
	type args struct {
		propertyFields []ce.PropertyField
	}

	propFields := make([]ce.PropertyField, 1)
	propFields[0] = &PropertyField{
		Object: Object{
			Name: "IssueDate",
		},
	}

	expected := make([]*PropertyField, 1)
	expected[0] = &PropertyField{
		Object: Object{
			Name: "IssueDate",
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "itStoresAnArrayOfPropertyFields",
			args: args{
				propFields,
			},
			fields: fields{
				PropertyFields: expected,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DocumentClass{
				WorkspaceId:    tt.fields.WorkspaceId,
				PropertyFields: tt.fields.PropertyFields,
				Object:         tt.fields.Object,
			}

			props := d.GetPropertyFields()
			if props[0].GetName() != "IssueDate" {
				t.Errorf("Incorrect Property Field() = %v, want %v", props, tt.fields)
			}
		})
	}
}
