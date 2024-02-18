package operator

import "encoding/json"

type OperatorValue struct {
	Op    string
	Value interface{}
}

const (
	EQUALS                     = "equals"
	CONTAINS                   = "contains"
	STARTS_WITH                = "startsWith"
	ENDS_WITH                  = "endsWith"
	IS                         = "is"         // boolean is and date is
	IS_EMPTY                   = "isEmpty"    // works for string, date, number
	IS_NOT_EMPTY               = "isNotEmpty" // works for string, date, number
	IS_ANY_OF                  = "isAnyOf"    // works for string, date, number
	NUMBER_EQUALS              = "="
	NUMBER_NOT_EQUALS          = "!="
	NUMBER_GREATER_THAN        = ">"
	NUMBER_GREATER_THAN_EQUALS = ">="
	NUMBER_LESS_THAN           = "<"
	NUMBER_LESS_THAN_EQUALS    = "<="
	DATE_IS_NOT                = "not"
	DATE_IS_AFTER              = "after"
	DATE_IS_ON_OR_AFTER        = "onOrAfter"
	DATE_IS_BEFORE             = "before"
	DATE_IS_ON_OR_BEFORE       = "onOrBefore"
)

var MapSqlOperators = map[string]string{
	EQUALS:                     "=",
	CONTAINS:                   "LIKE",
	STARTS_WITH:                "LIKE",
	ENDS_WITH:                  "LIKE",
	IS_EMPTY:                   "IS NULL",
	IS_NOT_EMPTY:               "IS NOT NULL",
	IS_ANY_OF:                  "IN",
	NUMBER_EQUALS:              "=",
	NUMBER_NOT_EQUALS:          "!=",
	NUMBER_GREATER_THAN:        ">",
	NUMBER_GREATER_THAN_EQUALS: ">=",
	NUMBER_LESS_THAN:           "<",
	NUMBER_LESS_THAN_EQUALS:    "<=",
	IS:                         "=",
	DATE_IS_NOT:                "!=",
	DATE_IS_AFTER:              ">",
	DATE_IS_BEFORE:             "<",
	DATE_IS_ON_OR_AFTER:        ">=",
	DATE_IS_ON_OR_BEFORE:       "<=",
}

func getJsonSlice(str string) []string {
	var strSlice []string
	isValid := json.Unmarshal([]byte(str), &strSlice) == nil
	if isValid {
		return strSlice
	}
	return []string{}
}

func GetSqlOperatorValue(govalidityOperator string, value string) *OperatorValue {
	op, ok := MapSqlOperators[govalidityOperator]
	if !ok {
		op = "="
	}
	var val interface{} = value
	switch govalidityOperator {
	case CONTAINS, STARTS_WITH, ENDS_WITH:
		val = "%" + value + "%"
	case IS_ANY_OF:
		val = getJsonSlice(value)
	}
	return &OperatorValue{
		Op:    op,
		Value: val,
	}
}
