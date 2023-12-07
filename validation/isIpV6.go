package validation

import (
	"net"
	"strings"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsIpV6(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsIpV6, label, value)
	str := value.(string)
	ip := net.ParseIP(str)
	isValid := ip != nil && strings.Contains(str, ":")
	if isValid {
		return true, nil
	}
	return false, err
}
