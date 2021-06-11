package ce

type FieldType string

const (
	FieldTypeString   FieldType = "string"
	FieldTypeInteger            = "integer"
	FieldTypeDouble             = "double"
	FieldTypeDateTime           = "date"
)

// RuleType identifies how the rule shall be evaluated
type RuleType string

const (
	RuleTypeRegex          RuleType = "regex"
	RuleTypeJavascriptRule          = "javascript"
)

type FieldRule interface {
	SetRuleType(ruleType RuleType)
	GetRuleType() RuleType

	SetRuleExpression(ruleExpression string)
	GetRuleExpression() string

	SetErrorMessage(errorMessage string)
	GetErrorMessage() string
}

// PropertyField represents an attribute that is defined by the end-user at run time
//
// PropertyField is a per-document class construct. This means that a `property` e.g. IssueDate can appear
// in more than one DocumentClass yet can have very different validation rules
type PropertyField interface {
	SetFieldType(fieldType FieldType)
	GetFieldType() FieldType

	SetFieldRule(propertyRule FieldRule)
	GetFieldRule() FieldRule

	Object
}
