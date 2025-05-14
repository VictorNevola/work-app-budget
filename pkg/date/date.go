package date

import (
	"time"
)

func FromString(dateStr string) (*time.Time, error) {
	time, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, ErrCannotConvertDate
	}

	return &time, nil
}
