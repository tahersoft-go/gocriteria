package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsMax(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	max := params[1].(int)

	label, ok := (*message.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	errMessage := message.Default.IsMax
	errMessage = strings.ReplaceAll(errMessage, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", fmt.Sprintf("%v", dataValue))
	errMessage = strings.ReplaceAll(errMessage, "{max}", strconv.Itoa(max))
	err := errors.New(errMessage)

	if dataValue.(float64) > float64(max) {
		return false, err
	}

	return true, nil
}
