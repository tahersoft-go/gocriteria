package validation

import "github.com/tahersoft-go/gocriteria/message"

func IsHost(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsHost, label, value)
	str := value.(string)
	isIp, _ := IsIp(field, str)
	isDNSName, _ := IsDNSName(field, str)
	isValid := isIp || isDNSName
	if isValid {
		return true, nil
	}
	return false, err
}
