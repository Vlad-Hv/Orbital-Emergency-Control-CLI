package ui

import (
	"errors"
)

func IntInput(err error) error {
	if err != nil {
		return errors.New("input must be int")
	}

	return nil
}
