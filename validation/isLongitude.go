package validation

import (
	"regexp"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsLogitude(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsLogitude, label, value)
	str := value.(string)
	rxLogitude := regexp.MustCompile("^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$")
	isValid := rxLogitude.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
