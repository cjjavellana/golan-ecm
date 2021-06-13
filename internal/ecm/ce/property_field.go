package ce

type FieldType string

const (
	FieldTypeString   FieldType = "string"
	FieldTypeInteger            = "integer"
	FieldTypeDouble             = "double"
	FieldTypeDateTime           = "date"
	FieldTypeComplex            = "complex"
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
// PropertyField(s) can be simple (a field to store any value) or can be derived (a field whose value depends on the
// result of the calculation expression). A PropertyField is considered derived when CalculationExpression is not empty.
//
// CalculationExpression must be a valid javascript expression since Golan ECM uses an embedded version of the v8 engine.
//
// PropertyField is used to described the fields of a particular DocumentClass
// e.g.
// Book (DocumentClass)
//   - ISBN
//  	- FieldType: string
// 		- RuleType: regex
//		- ValidationExpression: ^([A-Za-z0-9]){10,20}$
//		- ErrorMessage: ISBN must be alphanumeric and must be between 10 and 20 characters)
//		- Repeatable: false
//   - Author (string)
//   - DatePublished (date)
type PropertyField struct {
	// FieldType describes the datatype of the particular property field
	FieldType FieldType

	// RuleType determines the runtime engine to be used to evaluate the expression
	RuleType RuleType

	// ValidationExpression is an expression to be evaluated to determine the validity of the
	// of the value referred to by this property field
	ValidationExpression string

	// CalculationExpression is an expression to be evaluated to derive the value referred to by this
	// property field
	// LineItem (Property Field)
	//  - Qty (Property Field)
	//  - Unit Cost (Property Field)
	//  - Total (Derived Field, {{ qty * unitCode }}
	CalculationExpression string

	// ErrorMessage is the error message to be returned should the result of the
	// evaluation of ValidationExpression and/or an error is returned
	ErrorMessage string

	// Repeatable indicates that a property field can be repeated e.g. LineItems
	Repeatable bool

	Name        string
	Label       string
	Description string

	// PropertyField itself can contain other properties. An Example of which is an invoice line item property.
	// A line item can
	PropertyField []*PropertyField
}
