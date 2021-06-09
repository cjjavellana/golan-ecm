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
type PropertyField interface {
	// GetName returns the name of the attribute. Alphanumeric characters only.
	// interestRate, issueDate
	GetName() string
	GetLabel() string

	SetFieldType(fieldType FieldType)
	GetFieldType() FieldType

	SetDescription(description string)
	GetDescription() string

	SetFieldRule(propertyRule FieldRule)
	GetFieldRule() FieldRule

	Object
}
