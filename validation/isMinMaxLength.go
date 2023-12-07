package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tahersoft-go/gocriteria/convertor"
	"github.com/tahersoft-go/gocriteria/message"
)

func IsMinMaxLength(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	min := params[1].(int)
	max := params[2].(int)

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

	errMessage := strings.ReplaceAll(message.Default.IsMinMaxLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{min}", strconv.Itoa(min))
	errMessage = strings.ReplaceAll(errMessage, "{max}", strconv.Itoa(max))
	err := errors.New(errMessage)

	if len(value) < min || len(value) > max {
		return false, err
	}

	return true, nil
}
