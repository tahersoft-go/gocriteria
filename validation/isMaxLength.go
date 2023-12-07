package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tahersoft-go/gocriteria/convertor"
	"github.com/tahersoft-go/gocriteria/message"
)

func IsMaxLength(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	max := params[1].(int)

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

	errMessage := strings.ReplaceAll(message.Default.IsMaxLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{max}", strconv.Itoa(max))
	err := errors.New(errMessage)

	if len(value) > max {
		return false, err
	}

	return true, nil
}
