package validation

import (
	"strings"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsUpperCase(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsUpperCase, label, value)
	str := value.(string)
	isValid := str == strings.ToUpper(str)
	if isValid {
		return true, nil
	}
	return false, err
}
