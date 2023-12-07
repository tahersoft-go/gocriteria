package validation

import (
	"net"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsIp(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsIp, label, value)
	str := value.(string)
	isValid := net.ParseIP(str) != nil
	if isValid {
		return true, nil
	}
	return false, err
}
