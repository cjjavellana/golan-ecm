package ce

import "github.com/google/uuid"

type FieldType string

const (
	FieldTypeString   FieldType = "string"
	FieldTypeInteger            = "integer"
	FieldTypeDouble             = "double"
	FieldTypeDateTime           = "date"
)

type RuleType string

const (
	RuleTypeRegex          RuleType = "regex"
	RuleTypeJavascriptRule          = "javascript"
)

type PropertyRule interface {
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

	SetPropertyRule(propertyRule PropertyRule)
	GetPropertyRule() PropertyRule

	// GetWorkspaceObjectId returns the ObjectId of the workspace that this attribute belongs to
	GetWorkspaceObjectId() uuid.UUID

	Object
}
