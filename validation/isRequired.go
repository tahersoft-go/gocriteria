package validation

import (
	"strings"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsRequired(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsRequired, label, value)
	str := value.(string)
	if strings.Trim(str, " ") == "" {
		return false, err
	}
	return true, nil
}
