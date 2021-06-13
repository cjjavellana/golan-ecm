package builder

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"cjavellana.me/ecm/golan/internal/ecm/pb"
	"strings"
)

type PropertyFieldHierarchy struct {
	ObjectStore ce.ObjectStore
}

func (p *PropertyFieldHierarchy) Build(gRpcPropertyFields []*pb.PropertyField) []*ce.PropertyField {
	propFields := make([]*ce.PropertyField, len(gRpcPropertyFields))
	p.buildInternal(propFields, gRpcPropertyFields)

	return propFields
}

func (p *PropertyFieldHierarchy) buildInternal(acc []*ce.PropertyField, gRpcPropertyFields []*pb.PropertyField) {
	for i, v := range gRpcPropertyFields {
		pf := p.ObjectStore.NewPropertyField(
			ce.ObjectDescriptor{
				Name:        v.Name,
				Label:       v.Label,
				Description: v.Description,
			},
			ce.FieldType(strings.ToLower(v.FieldType.String())),
		)
		pf.ValidationExpression = v.ValidationExpr
		pf.CalculationExpression = v.CalculationExpr
		pf.Repeatable = v.Repeatable

		acc[i] = &pf

		if len(v.SubProperties) > 0 {
			pf.PropertyField = make([]*ce.PropertyField, len(v.SubProperties))
			p.buildInternal(pf.PropertyField, v.SubProperties)
		}
	}

}
