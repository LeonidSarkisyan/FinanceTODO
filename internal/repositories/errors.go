package repositories

import (
	"fmt"
)

type TooLongValueError struct {
	Field    string
	MaxValue int
}

func (e TooLongValueError) Error() string {
	return fmt.Sprintf("%s превышает максимально допустимый размер (%d)", e.Field, e.MaxValue)
}
