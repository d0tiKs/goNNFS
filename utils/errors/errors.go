package errorsutils

import (
	"errors"
	"fmt"
)

func BuildError(err error, format string, vars ...interface{}) error {
	errorMessage := fmt.Sprintf(format, vars...)

	if err == nil {
		return errors.New(errorMessage)
	}

	embeddedError := fmt.Sprintf("\nSee error bellow:\n%s", err.Error())
	return errors.New(errorMessage + embeddedError)
}
