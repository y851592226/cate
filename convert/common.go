package convert

import (
	"fmt"
	"time"
)

var (
	BaseTimeFormat    = "2006-01-02 15:04:05"
	DefaultTimeFormat = time.RFC3339
)

func unSupportTypeError(i interface{}) error {
	return fmt.Errorf("unsupport type of %T", i)
}

func unSupportTimeFormatError(s string) error {
	return fmt.Errorf("unsupport time format of %s", s)
}
