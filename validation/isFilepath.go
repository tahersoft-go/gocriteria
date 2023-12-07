package validation

import (
	"regexp"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsFilepath(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsFilepath, label, value)
	str := value.(string)
	rxFilepath := regexp.MustCompile("^(\\/.*?)[\\.][\\w]{1,}$")
	isValid := rxFilepath.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
