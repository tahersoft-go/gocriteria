package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsMin(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	min := params[1].(int)

	label, ok := (*message.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	errMessage := message.Default.IsMin
	errMessage = strings.ReplaceAll(errMessage, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", fmt.Sprintf("%v", dataValue))
	errMessage = strings.ReplaceAll(errMessage, "{min}", strconv.Itoa(min))
	err := errors.New(errMessage)

	if dataValue.(float64) < float64(min) {
		return false, err
	}

	return true, nil
}
