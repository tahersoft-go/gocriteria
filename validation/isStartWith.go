package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tahersoft-go/gocriteria/convertor"
	"github.com/tahersoft-go/gocriteria/message"
)

func IsStartWith(field string, dataValue interface{}, str string) (bool, error) {
	fieldLabel := field

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

	if strings.Index(value, str) == 0 {
		return true, nil
	}

	errMessage := strings.ReplaceAll(message.Default.IsMaxLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{str}", str)
	err := errors.New(errMessage)

	return false, err
}
