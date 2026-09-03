package orbital

import (
	"errors"
	"fmt"
)

func ZoneValidate(option int, station OrbitalStation) error {
	if station.CurrentZone.Condition == "unstable" {
		if option < 1 || option > 3 {
			return errors.New("invalid option")
		}
	} else {
		if option < 1 || option == 2 || option > 3 {
			return errors.New("invalid option")
		}
	}

	return nil
}

func FixValidate(room map[string]int, inventory map[string]int) error {
	var result int
	for resourse, amount := range room {
		inv_resourse, ok := inventory[resourse]
		if !ok {
			return fmt.Errorf("there isnot this resourse in inventory: %s", resourse)
		}
		result = inv_resourse - amount
		if result < 0 {
			return errors.New("you donot have enough material in the inventory")
		}
	}
	return nil
}

func CommunicationValidate(station OrbitalStation) error {
	if station.Zones[362].Condition == "unstable" {
		return errors.New("cannot fix the communication, when reactor isnot working")
	}

	return nil
}

func Signal(station OrbitalStation) error {
	if station.Zones[363].Condition == "unstable" {
		return errors.New("cannot send an emergy signal, when communication isnot working")
	}

	return nil
}
