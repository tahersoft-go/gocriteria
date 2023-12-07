package validation

import (
	"github.com/tahersoft-go/gocriteria/convertor"
	"github.com/tahersoft-go/gocriteria/message"
)

func IsNumber(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsNumber, label, value)
	number, numberError := convertor.ToNumber(value)
	if numberError != nil || number == nil {
		return false, err
	}
	return true, nil
}
