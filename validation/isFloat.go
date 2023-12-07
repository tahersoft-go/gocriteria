package validation

import (
	"regexp"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsFloat(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsFloat, label, value)
	str := value.(string)
	rxInt := regexp.MustCompile("^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$")
	isValid := rxInt.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
