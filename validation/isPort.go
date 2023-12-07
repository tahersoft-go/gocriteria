package validation

import (
	"strconv"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsPort(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsPort, label, value)
	str := value.(string)
	if i, err := strconv.Atoi(str); err == nil && i > 0 && i < 65536 {
		return true, nil
	}
	return false, err
}
