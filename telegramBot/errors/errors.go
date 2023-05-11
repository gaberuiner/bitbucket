package errors

import "fmt"

func Error(str string, err error) error {
	return fmt.Errorf("%s: %w", str, err)
}
