package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tahersoft-go/gocriteria/convertor"
	"github.com/tahersoft-go/gocriteria/message"
)

func IsIn(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	in := params[1].([]string)

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

	errMessage := strings.ReplaceAll(message.Default.IsIn, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{in}", fmt.Sprintf("%v", in))
	err := errors.New(errMessage)

	found := false
	for _, inVal := range in {
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
