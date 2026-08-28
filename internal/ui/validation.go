package ui

import (
	"errors"
)

func ValidateInput(err error) error {
	if err != nil {
		return errors.New("input must have only number")
	}

	return nil
}
