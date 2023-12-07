package validation

import (
	"regexp"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsInt(field string, params ...interface{}) (bool, error) {

	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsInt, label, value)
	str := value.(string)
	rxInt := regexp.MustCompile("^(?:[-+]?(?:0|[1-9][0-9]*))$")
	isValid := rxInt.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
