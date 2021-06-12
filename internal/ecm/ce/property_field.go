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

// PropertyField represents an attribute that is defined by the end-user at run time
//
// PropertyField is a per-document class construct. This means that a `property` e.g. IssueDate can appear
// in more than one DocumentClass yet can have very different validation rules
//
// PropertyField is used to described the fields of a particular DocumentClass
// e.g.
// Book (DocumentClass)
//   - ISBN (string)
//   - Author (string)
//   - DatePublished (date)
type PropertyField struct {
	// FieldType describes the datatype of the particular property field
	FieldType FieldType

	// RuleType determines the runtime engine to be used to evaluate the expression
	RuleType RuleType

	// RuleExpression is the value to be evaluated
	RuleExpression string

	// ErrorMessage is the error message to be returned should the result of the
	// evaluation of RuleExpression and/or an error is returned
	ErrorMessage string

	Name        string
	Label       string
	Description string
}
