package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tahersoft-go/gocriteria/convertor"
	"github.com/tahersoft-go/gocriteria/message"
)

func IsInRange(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	from := params[1].(int)
	to := params[2].(int)

	label, ok := (*message.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	value := ""
	number, errConv := convertor.ToNumber(dataValue)
	if errConv == nil && number != nil {
		value = fmt.Sprintf("%v", *number)
	}

	errMessage := strings.ReplaceAll(message.Default.IsInRange, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{from}", strconv.Itoa(from))
	errMessage = strings.ReplaceAll(errMessage, "{to}", strconv.Itoa(to))
	err := errors.New(errMessage)

	if number == nil {
		return false, err
	}

	if int(*number) < from || int(*number) > to {
		return false, err
	}

	return true, nil
}
