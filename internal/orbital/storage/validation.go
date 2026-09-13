package storage

import (
	"errors"
)

var (
	errNameIncorrect     = errors.New("ERROR: incorrect resource name")
	errResourceNotEnough = errors.New("ERROR: in storage not enough resource")
	errAmountInvalid     = errors.New("ERROR: invalid resource amount")
)

func CheckData(resource string, amount int, storage map[string]int) error {
	resAmount, ok := storage[resource]

	if !ok {
		return errNameIncorrect
	}

	if resAmount < amount {
		return errResourceNotEnough
	}

	if amount < 1 {
		return errAmountInvalid
	}

	return nil
}
