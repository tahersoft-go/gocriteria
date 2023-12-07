package validation

import (
	"encoding/json"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsJson(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsJson, label, value)
	str := value.(string)
	var js json.RawMessage
	isValid := json.Unmarshal([]byte(str), &js) == nil
	if isValid {
		return true, nil
	}
	return false, err
}
