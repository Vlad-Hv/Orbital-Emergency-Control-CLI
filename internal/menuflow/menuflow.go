package menuflow

import (
	"OStation/internal/ui"
	"fmt"
)

func ChooseZoneMenu() error {
	id, err := ui.GetRoomID()

	err = ui.ValidateInput(err)
	if err != nil {
		return fmt.Errorf("cannot change zone: %w", err)
	}

	switch id {
	case 361:
		ui.ControlRoom()
		return nil
	case 362:
		//validation and after for loop and bussiness logik of this zone
		//меню реактора делать в файле зоне меню
		return nil
	case 363:
		//validation and after for loop and bussiness logik of this zone

		return nil
	case 364:
		//validation and after for loop and bussiness logik of this zone

		return nil
	case 365:
		//validation and after for loop and bussiness logik of this zone

		return nil
	default:
		ui.InvalidRoom()
		return nil
	}
}
