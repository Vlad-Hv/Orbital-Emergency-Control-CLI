package main

import (
	"OStation/internal/authorization"
	"OStation/internal/orbital"
	"OStation/internal/ui"

	"OStation/internal/orbital/storage"
	"fmt"
)

func main() {
	ui.FirstWarning()
	err := authorization.Auth()

	if err != nil {
		fmt.Println(err)
		return
	}
	ui.Access()
	stationStorage := storage.CreateStorage()
	inventory := storage.CreateInventory()
	orbitalStation := orbital.CreateStation()

	for {
		option, err := ui.GetMenuOption()
		err = ui.ValidateInput(err)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 0 {
			fmt.Println("developer exit")
			break
		}

		switch option {
		case 1:
			fmt.Println(orbitalStation.CurrentZone)

		case 2:
			id, err := ui.GetRoomID()

			err = ui.ValidateInput(err)
			if err != nil {
				fmt.Println(err)
				continue
			}

			switch id {
			case 361:
				//validation is it possiblle to get into this room and after for loop and bussiness logik of this zone

			case 362:
				//validation and after for loop and bussiness logik of this zone

			case 363:
				//validation and after for loop and bussiness logik of this zone

			case 364:
				//validation and after for loop and bussiness logik of this zone

			case 365:
				//validation and after for loop and bussiness logik of this zone

			default:
				ui.InvalidRoom()
			}
		case 3:

		case 4:
			orbitalStation.CurrentZone = *orbitalStation.Zones[364]
			for {
				option, err := ui.GetStoregOption()
				err = ui.ValidateInput(err)

				if err != nil {
					fmt.Println(err)
					continue
				}

				if option == 3 {
					fmt.Println("You come back into controll room")
					orbitalStation.CurrentZone = *orbitalStation.Zones[361]
					break
				}

				switch option {
				case 1:
					ui.PrintStorage(stationStorage)
				case 2:
					//сделать как раз таки обработчик инвентаря, сделать изменине стореджа и перекладывать в инвентарь + валидация
					resourse, amount, err := ui.GetTakingData()

					err = ui.ValidateInput(err)
					if err != nil {
						fmt.Println(err)
						continue
					}

					err = storage.CheckStoragesData(resourse, amount, stationStorage)

					if err != nil {
						fmt.Println(err)
						continue
					}

					storage.TakeResourse(&stationStorage, &inventory, resourse, amount)
					//теперь сделать как раз забирание элемента и сторедж и добавление в инвентарь

				case 3:

				}
			}
		case 5:

		case 6:
			ui.PrintInv(inventory)
		default:
			fmt.Println("\nMESSAGE: invalid option")
		}

	}

}
