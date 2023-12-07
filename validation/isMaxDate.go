package validation

import (
	"time"

	"github.com/tahersoft-go/gocriteria/message"
)

func IsMaxDate(field string, value interface{}, max time.Time) (bool, error) {
	if value.(time.Time).After(max) {
		return false, GetErrorMessageByFieldValue(message.Default.IsMaxDate, field, value)
	}
	return true, nil
}
