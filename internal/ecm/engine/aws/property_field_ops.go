package aws

import "cjavellana.me/ecm/golan/internal/ecm/ce"

type PropertyFieldOperation struct {
	db *CollectionStore
}

func (p *PropertyFieldOperation) NewPropertyField(
	descriptor ce.ObjectDescriptor,
	fieldType ce.FieldType,
) ce.PropertyField {
	return ce.PropertyField{
		Name:        descriptor.Name,
		Label:       descriptor.Label,
		Description: descriptor.Description,
		FieldType:   fieldType,
	}
}