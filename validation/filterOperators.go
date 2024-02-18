package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tahersoft-go/gocriteria/convertor"
	"github.com/tahersoft-go/gocriteria/message"
	"github.com/tahersoft-go/gocriteria/operator"
)

func FilterOperators(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	constrains := params[1].([]string)

	if len(constrains) == 0 {
		constrains = append(constrains,
			operator.EQUALS,
			operator.CONTAINS,
			operator.STARTS_WITH,
			operator.ENDS_WITH,
			operator.IS,
			operator.IS_EMPTY,
			operator.IS_NOT_EMPTY,
			operator.IS_ANY_OF,
			operator.NUMBER_EQUALS,
			operator.NUMBER_GREATER_THAN,
			operator.NUMBER_GREATER_THAN_EQUALS,
			operator.NUMBER_LESS_THAN,
			operator.NUMBER_LESS_THAN_EQUALS,
			operator.NUMBER_NOT_EQUALS,
			operator.DATE_IS_NOT,
			operator.DATE_IS_AFTER,
			operator.DATE_IS_ON_OR_AFTER,
			operator.DATE_IS_BEFORE,
			operator.DATE_IS_ON_OR_BEFORE,
		)
	}

	label, ok := (*message.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	value := ""
	number, errConv := convertor.ToNumber(dataValue)
	if errConv == nil && number != nil {
		value = fmt.Sprintf("%v", *number)
	} else {
		value = dataValue.(string)
	}

	errMessage := strings.ReplaceAll(message.Default.IsFilterOperators, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{in}", fmt.Sprintf("%v", constrains))
	err := errors.New(errMessage)

	found := false
	for _, inVal := range constrains {
		if value == inVal {
			found = true
			break
		}
	}
	if !found {
		return false, err
	}

	return true, nil
}
