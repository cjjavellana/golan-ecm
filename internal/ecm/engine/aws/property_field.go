package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
)

type PropertyField struct {

	// FieldType identifies the data type that the field supports
	// e.g. string, number, date
	FieldType ce.FieldType `bson:"FieldType,omitempty"`

	Object `bson:",inline"`
}

func (p *PropertyField) SetFieldType(fieldType ce.FieldType) {
	p.FieldType = fieldType
}

func (p *PropertyField) GetFieldType() ce.FieldType {
	return p.GetFieldType()
}

func (p *PropertyField) SetFieldRule(propertyRule ce.FieldRule) {
	panic("implement me")
}

func (p *PropertyField) GetFieldRule() ce.FieldRule {
	panic("implement me")
}
