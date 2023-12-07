package validation

import (
	"time"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsMinTime(field string, value interface{}, min time.Time) (bool, error) {
	if value.(time.Time).Before(min) {
		return false, GetErrorMessageByFieldValue(message.Default.IsMinTime, field, value)
	}
	return true, nil
}
