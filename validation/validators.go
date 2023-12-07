package validation

import (
	"errors"
	"strings"

	"github.com/tahersoft-go/gocriteria/message"
)

func GetFieldLabelAndValue(field string, params []interface{}) (string, interface{}) {
	fieldLabel := field
	value := params[0]
	label, ok := (*message.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}
	return fieldLabel, value
}

func GetErrorMessageByFieldValue(baseErrorMessage string, field string, value interface{}) error {
	errMessage := strings.ReplaceAll(baseErrorMessage, "{field}", field)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value.(string))
	return errors.New(errMessage)
}
