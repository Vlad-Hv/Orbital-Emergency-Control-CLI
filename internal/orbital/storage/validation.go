package storage

import (
	"errors"
)

func CheckStoragesData(resourse string, amount int, storage map[string]int) error {
	resAmount, ok := storage[resourse]

	if !ok {
		return errors.New("ERROR: incorrect resourse name")
	}

	if resAmount < amount {
		return errors.New("ERROR: in storage not enough resourse")
	}

	if amount < 1 {
		return errors.New("ERROR: invalid resourse amount")
	}

	return nil
}
