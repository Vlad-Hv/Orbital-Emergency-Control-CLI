package orbital

import (
	"errors"
	"fmt"
)

func ReactorValidate(option int, station OrbitalStation) error {
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
		result = amount - inv_resourse
		if result < 0 {
			return errors.New("you donot have enough material in the inventory")
		}
	}
	return nil
}
