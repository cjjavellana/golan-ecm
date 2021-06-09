package aws

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
)

type PropertyField struct {
	// Name represents an internal name for the property.
	// e.g. IssueDate, MaturityDate, IssueAmount
	Name string

	// Label represents a human-friendly name used for for display purposes
	// e.g. Issue Date, Maturity Date, Issue Amount
	Label string

	// FieldType identifies the data type that the field supports
	// e.g. string, number, date
	FieldType ce.FieldType

	// Description Describes the field
	Description string

	Object `bson:",inline"`
}

func (p *PropertyField) GetName() string {
	return p.Name
}

func (p *PropertyField) GetLabel() string {
	return p.Label
}

func (p *PropertyField) SetFieldType(fieldType ce.FieldType) {
	p.FieldType = fieldType
}

func (p *PropertyField) GetFieldType() ce.FieldType {
	return p.GetFieldType()
}

func (p *PropertyField) SetDescription(description string) {
	p.Description = description
}

func (p *PropertyField) GetDescription() string {
	return p.Description
}

func (p *PropertyField) SetFieldRule(propertyRule ce.FieldRule) {
	panic("implement me")
}

func (p *PropertyField) GetFieldRule() ce.FieldRule {
	panic("implement me")
}
